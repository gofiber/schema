package schema

import (
	"reflect"
	"strconv"
	"testing"

	utils "github.com/gofiber/utils/v2"
	utilstrings "github.com/gofiber/utils/v2/strings"
)

func TestNextPathSegment(t *testing.T) {
	cases := []struct {
		path    string
		start   int
		end     int
		segment string
		wantErr bool
	}{
		{"a.b", 0, 1, "a", false},
		{"a.b", 2, 3, "b", false},
		{"abc.defgh", 0, 3, "abc", false},
		// dot exactly on the SWAR word boundary
		{"abcdefgh.x", 0, 8, "abcdefgh", false},
		// segments longer than one SWAR word
		{"averylongsegmentname.next", 0, 20, "averylongsegmentname", false},
		{"averylongsegmentname.next", 21, 25, "next", false},
		{"noseparatorhere", 0, 15, "noseparatorhere", false},
		// non-ASCII bytes must pass through the word scan untouched
		{"héllo.wörld", 0, 6, "héllo", false},
		{"", 0, 0, "", true},
		{".abc", 0, 0, "", true},
		{"a..b", 2, 0, "", true},
	}
	for _, c := range cases {
		end, segment, err := nextPathSegment(c.path, c.start)
		if c.wantErr {
			if err == nil {
				t.Errorf("nextPathSegment(%q, %d): expected error", c.path, c.start)
			}
			continue
		}
		if err != nil {
			t.Errorf("nextPathSegment(%q, %d): unexpected error %v", c.path, c.start, err)
			continue
		}
		if end != c.end || segment != c.segment {
			t.Errorf("nextPathSegment(%q, %d) = (%d, %q), want (%d, %q)",
				c.path, c.start, end, segment, c.end, c.segment)
		}
	}
}

func BenchmarkNextPathSegment(b *testing.B) {
	const path = "averylongsegmentname.middle.tail"
	for b.Loop() {
		for start := 0; ; {
			end, _, err := nextPathSegment(path, start)
			if err != nil || end == len(path) {
				break
			}
			start = end + 1
		}
	}
}

// The setDefaults fast path must not be defeated by anonymous pointers the
// walk can never allocate (unexported ones).
func TestNeedsDefaultsWalk(t *testing.T) {
	type hidden struct {
		X string `schema:"x"`
	}
	type onlyUnexported struct {
		*hidden
		A string `schema:"a"`
	}

	d := NewDecoder()
	if d.cache.get(reflect.TypeOf(onlyUnexported{})).needsDefaultsWalk {
		t.Error("unexported anonymous pointer must not force the defaults walk")
	}

	type Inner struct {
		Y string `schema:"y"`
	}
	type withExported struct {
		*Inner
		A string `schema:"a"`
	}
	if !d.cache.get(reflect.TypeOf(withExported{})).needsDefaultsWalk {
		t.Error("exported anonymous pointer must keep the defaults walk (it allocates)")
	}

	type withDefaults struct {
		A string `schema:"a,default:z"`
	}
	if !d.cache.get(reflect.TypeOf(withDefaults{})).needsDefaultsWalk {
		t.Error("default tags must keep the defaults walk")
	}
}

func BenchmarkParsePathCacheMiss(b *testing.B) {
	type Nested struct {
		Value string `schema:"value"`
	}
	type Outer struct {
		Items []Nested `schema:"items"`
	}
	d := NewDecoder()
	typ := reflect.TypeOf(Outer{})
	b.ReportAllocs()
	info := d.cache.get(typ)
	for b.Loop() {
		info.paths.clear()
		if _, err := d.cache.parsePath("items.0.value", typ); err != nil {
			b.Fatal(err)
		}
	}
}

// foldASCIILower backs the allocation-free case-insensitive probes, so it has
// to agree with the utils folder it stands in for at every length and leave
// non-ASCII bytes untouched.
func TestFoldASCIILower(t *testing.T) {
	t.Parallel()

	alphabet := make([]byte, 0, 256)
	for c := 0; c < 256; c++ {
		alphabet = append(alphabet, byte(c))
	}

	for n := 0; n <= maxDirectKeyLen; n++ {
		for off := 0; off < 256; off++ {
			in := make([]byte, n)
			for i := range in {
				in[i] = alphabet[(off+i*7)%256]
			}
			src := string(in)
			var buf [maxDirectKeyLen]byte
			changed := foldASCIILower(buf[:], src)
			got := string(buf[:n])
			want := utilstrings.ToLower(src)
			if got != want {
				t.Fatalf("foldASCIILower(%q) = %q, want %q", src, got, want)
			}
			if changed != (got != src) {
				t.Fatalf("foldASCIILower(%q) reported changed=%v, want %v", src, changed, got != src)
			}
		}
	}
}

// A mixed-case path that the precomputed direct map cannot answer walks the
// generic parser, where every segment is case-folded before its field lookup;
// that fold must not allocate.
func BenchmarkParsePathCacheMissMixedCase(b *testing.B) {
	type Nested struct {
		Value string `schema:"value"`
	}
	type Outer struct {
		Items []Nested `schema:"items"`
	}
	d := NewDecoder()
	typ := reflect.TypeOf(Outer{})
	b.ReportAllocs()
	info := d.cache.get(typ)
	for b.Loop() {
		info.paths.clear()
		if _, err := d.cache.parsePath("Items.0.Value", typ); err != nil {
			b.Fatal(err)
		}
	}
}

// Decode callers (e.g. Fiber binders) pass map keys aliasing reused request
// buffers; pathCache must clone the key so later buffer reuse cannot poison it.
func TestParsePathDetachesCacheKey(t *testing.T) {
	// Use a slice-index path: statically-resolvable keys are served from the
	// direct map and never stored in the sync.Map path cache.
	type Item struct {
		Value string `schema:"value"`
	}
	type S struct {
		Items []Item `schema:"items"`
	}
	d := NewDecoder()
	buf := []byte("items.0.value")
	var s S
	if err := d.Decode(&s, map[string][]string{utils.UnsafeString(buf): {"x"}}); err != nil {
		t.Fatal(err)
	}
	copy(buf, "xtems.9.qalue") // simulate fasthttp buffer reuse mutating the key bytes

	keys := cachedPathKeys(&d.cache.get(reflect.TypeOf(s)).paths)
	if len(keys) != 1 {
		t.Fatalf("expected 1 cached path, got %d (%v)", len(keys), keys)
	}
	if keys[0] != "items.0.value" {
		t.Fatalf("path cache key mutated to %q; key must be cloned before caching", keys[0])
	}
}

// cachedPathKeys lists everything the cache currently holds, across both its
// published map and its spill.
func cachedPathKeys(c *pathCache) []string {
	var keys []string
	if m := c.fast.Load(); m != nil {
		for k := range *m {
			keys = append(keys, k)
		}
	}
	c.spill.Range(func(k, _ any) bool {
		keys = append(keys, k.(string))
		return true
	})
	return keys
}

// Past maxFastPaths the cache stops growing the map it copies on every write
// and spills instead, but every path must still be cached and served.
func TestPathCacheSpill(t *testing.T) {
	t.Parallel()

	type Item struct {
		Value string `schema:"value"`
	}
	type S struct {
		Items []Item `schema:"items"`
	}
	d := NewDecoder()
	typ := reflect.TypeOf(S{})
	info := d.cache.get(typ)

	const n = maxFastPaths + 64
	want := make([]string, 0, n)
	for i := 0; i < n; i++ {
		want = append(want, "items."+strconv.Itoa(i)+".value")
	}
	for _, p := range want {
		if _, err := d.cache.parsePath(p, typ); err != nil {
			t.Fatalf("parsePath(%q): %v", p, err)
		}
	}
	if !info.paths.sealed.Load() {
		t.Fatal("cache should be sealed past maxFastPaths")
	}
	if got := len(*info.paths.fast.Load()); got != maxFastPaths {
		t.Fatalf("published map holds %d paths, want %d", got, maxFastPaths)
	}
	// Every path, spilled or not, is still served from the cache and still
	// resolves to the right index.
	for i, p := range want {
		parts, ok := info.paths.load(p)
		if !ok {
			t.Fatalf("path %q not cached", p)
		}
		if parts[0].index != i {
			t.Fatalf("path %q cached as index %d, want %d", p, parts[0].index, i)
		}
	}
	if got := len(cachedPathKeys(&info.paths)); got != n {
		t.Fatalf("cache holds %d paths, want %d", got, n)
	}

	info.paths.clear()
	if info.paths.sealed.Load() || info.paths.fast.Load() != nil {
		t.Fatal("clear must reset the cache")
	}
	if got := len(cachedPathKeys(&info.paths)); got != 0 {
		t.Fatalf("cache holds %d paths after clear, want 0", got)
	}
}
