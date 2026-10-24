package schema

import (
	"fmt"
	"reflect"
	"runtime"
	"slices"
	"strconv"
	"strings"
	"sync"
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
	if d.cache.get(reflect.TypeFor[onlyUnexported]()).needsDefaultsWalk {
		t.Error("unexported anonymous pointer must not force the defaults walk")
	}

	type Inner struct {
		Y string `schema:"y"`
	}
	type withExported struct {
		*Inner
		A string `schema:"a"`
	}
	if !d.cache.get(reflect.TypeFor[withExported]()).needsDefaultsWalk {
		t.Error("exported anonymous pointer must keep the defaults walk (it allocates)")
	}

	type withDefaults struct {
		A string `schema:"a,default:z"`
	}
	if !d.cache.get(reflect.TypeFor[withDefaults]()).needsDefaultsWalk {
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
	typ := reflect.TypeFor[Outer]()
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
// to agree with the utils folder it stands in for at every length, and leave
// non-ASCII bytes untouched.
func TestFoldASCIILower(t *testing.T) {
	t.Parallel()

	alphabet := make([]byte, 0, 256)
	for c := range 256 {
		alphabet = append(alphabet, byte(c))
	}

	for n := 0; n <= maxDirectKeyLen; n++ {
		for off := range 256 {
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

// A mixed-case path the direct map cannot answer walks the generic parser,
// which folds every segment before its field lookup; that must not allocate.
func BenchmarkParsePathCacheMissMixedCase(b *testing.B) {
	type Nested struct {
		Value string `schema:"value"`
	}
	type Outer struct {
		Items []Nested `schema:"items"`
	}
	d := NewDecoder()
	typ := reflect.TypeFor[Outer]()
	b.ReportAllocs()
	info := d.cache.get(typ)
	for b.Loop() {
		info.paths.clear()
		if _, err := d.cache.parsePath("Items.0.Value", typ); err != nil {
			b.Fatal(err)
		}
	}
}

// parsePathInfo rejects a dotless key that misses the direct map without
// parsing it, which is only sound if every field a bare alias can reach is in
// that map: a single-segment path must parse exactly when the map holds its
// folded form.
func TestDirectMapCoversEveryBareAlias(t *testing.T) {
	t.Parallel()

	type Leaf struct {
		Value string `schema:"value"`
	}
	type Embedded struct {
		Promoted string `schema:"promoted"`
	}
	type Dup struct {
		A string `schema:"same"`
		B string `schema:"same"`
	}
	// A named map type is still a map field, decoded entry by entry.
	type Attrs map[string]string
	type Shapes struct {
		Embedded
		Dup
		Plain      string     `schema:"plain"`
		MixedCase  string     `schema:"MixedCase"`
		Nested     Leaf       `schema:"nested"`
		NestedPtr  *Leaf      `schema:"nestedptr"`
		Items      []Leaf     `schema:"items"`
		Scalars    []int      `schema:"scalars"`
		Skipped    string     `schema:"-"`
		Unexported string     `schema:"unexported"`
		Dotted     string     `schema:"dot.ted"`
		Deep       [][]string `schema:"deep"`
		Attrs      Attrs      `schema:"attrs"`
	}

	types := []reflect.Type{
		reflect.TypeFor[Leaf](), reflect.TypeFor[Dup](),
		reflect.TypeFor[Embedded](), reflect.TypeFor[Shapes](),
	}
	// Every alias above, some case variants, and keys that only resemble one.
	keys := []string{
		"value", "same", "promoted", "plain", "mixedcase", "MixedCase", "MIXEDCASE",
		"nested", "nestedptr", "items", "scalars", "skipped", "-", "unexported",
		"dot.ted", "dotted", "deep", "embedded", "dup", "missing", "", "plainx", "Plain",
		"attrs", "ATTRS",
	}

	d := NewDecoder()
	for _, typ := range types {
		info := d.cache.get(typ)
		for _, key := range keys {
			if strings.IndexByte(key, '.') >= 0 {
				continue // the shortcut only claims anything about dotless keys
			}
			_, err := d.cache.parsePath(key, typ)
			_, inDirect := info.direct[utilstrings.ToLower(key)]
			if (err == nil) != inDirect {
				t.Fatalf("%s: parsePath(%q) err=%v but direct map entry=%v", typ, key, err, inDirect)
			}
		}
	}
}

// The path cache seals and spills under concurrent writers; every path must
// still resolve to its own index whichever side of the seal it landed on.
func TestPathCacheConcurrentSpill(t *testing.T) {
	type Item struct {
		Value string `schema:"value"`
	}
	type S struct {
		Items []Item `schema:"items"`
	}
	const goroutines, per = 8, maxFastPaths/2 + 64 // together well past the seal

	d := NewDecoder()
	var wg sync.WaitGroup
	errs := make(chan error, goroutines)
	for g := range goroutines {
		wg.Add(1)
		go func(g int) {
			defer wg.Done()
			for i := range per {
				idx := (g*per + i) % (maxFastPaths + 128) // overlap between goroutines
				key := "items." + strconv.Itoa(idx) + ".value"
				var s S
				if err := d.Decode(&s, map[string][]string{key: {key}}); err != nil {
					errs <- err
					return
				}
				if len(s.Items) != idx+1 || s.Items[idx].Value != key {
					errs <- fmt.Errorf("%s decoded into %d items, [%d]=%q", key, len(s.Items), idx, s.Items[idx].Value)
					return
				}
			}
		}(g)
	}
	wg.Wait()
	close(errs)
	for err := range errs {
		t.Fatal(err)
	}
	info := d.cache.get(reflect.TypeFor[S]())
	if !info.paths.sealed.Load() {
		t.Fatal("expected the cache to have sealed")
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

// Map keys are chosen by the client, so a path through a map field must not be
// cached: otherwise every distinct key ever sent would stay in the cache.
func TestMapKeysAreNotCached(t *testing.T) {
	t.Parallel()

	type Item struct {
		Name  string            `schema:"name"`
		Attrs map[string]string `schema:"attrs"`
	}
	type S struct {
		Attrs  map[string]string `schema:"attrs"`
		Nested struct {
			Attrs map[string]int `schema:"attrs"`
		} `schema:"nested"`
		Items []Item `schema:"items"`
	}

	d := NewDecoder()
	var s S
	// A path through a slice of structs is cached, as before.
	if err := d.Decode(&s, map[string][]string{"items.0.name": {"x"}}); err != nil {
		t.Fatal(err)
	}
	paths := &d.cache.get(reflect.TypeOf(s)).paths
	before := len(cachedPathKeys(paths))
	if before != 1 {
		t.Fatalf("expected the slice path to be cached, got %d entries", before)
	}

	const n = 2 * maxFastPaths
	for i := range n {
		key := strconv.Itoa(i)
		src := map[string][]string{
			"attrs." + key:         {"v"},
			"nested.attrs." + key:  {"1"},
			"items.0.attrs." + key: {"v"},
		}
		if err := d.Decode(&s, src); err != nil {
			t.Fatal(err)
		}
	}

	if after := len(cachedPathKeys(paths)); after != before {
		t.Fatalf("decoding %d distinct map keys grew the path cache from %d to %d entries", n, before, after)
	}
	if len(s.Attrs) != n || len(s.Nested.Attrs) != n || len(s.Items[0].Attrs) != n {
		t.Fatalf("got %d, %d and %d entries, want %d each", len(s.Attrs), len(s.Nested.Attrs), len(s.Items[0].Attrs), n)
	}
}

// cachedPathKeys lists everything the cache holds, published map and spill.
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

// Past maxFastPaths the cache spills instead of growing the map it copies on
// every write, but every path must still be cached and served.
func TestPathCacheSpill(t *testing.T) {
	t.Parallel()

	type Item struct {
		Value string `schema:"value"`
	}
	type S struct {
		Items []Item `schema:"items"`
	}
	d := NewDecoder()
	typ := reflect.TypeFor[S]()
	info := d.cache.get(typ)

	const n = maxFastPaths + 64
	want := make([]string, 0, n)
	for i := range n {
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
	// Every path, spilled or not, still resolves to the right index.
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

// The parser sizes its hop and part slices from a key's separator count, and
// that count arrives unvalidated: a key that is mostly separators is rejected
// at its first segment, so it must not reserve in proportion to its length
// first.
func TestParserReservationDoesNotScaleWithSeparators(t *testing.T) {
	c := newCache()
	typ := reflect.TypeOf(struct{ A string }{})

	const iters = 100
	allocated := func(dots int) uint64 {
		key := "a" + strings.Repeat(".", dots)
		if _, err := c.parsePath(key, typ); err == nil {
			t.Fatalf("parsePath(%d separators) succeeded, want rejection", dots)
		}
		var before, after runtime.MemStats
		runtime.GC()
		runtime.ReadMemStats(&before)
		for range iters {
			_, _ = c.parsePath(key, typ)
		}
		runtime.ReadMemStats(&after)
		return after.TotalAlloc - before.TotalAlloc
	}

	short := allocated(maxPathReserve)
	long := allocated(maxPathReserve * 512)
	if long > short*4 {
		t.Fatalf("%d separators allocate %d bytes over %d rejections against %d for %d separators: the reservation follows the key length",
			maxPathReserve*512, long, iters, short, maxPathReserve)
	}
}

// slotTypes returns n distinct struct types, more than the type cache has
// slots when n > typeCacheSlots, so that some must share one.
func slotTypes(n int) []reflect.Type {
	types := make([]reflect.Type, n)
	for i := range types {
		types[i] = reflect.StructOf([]reflect.StructField{{
			Name: "F" + strconv.Itoa(i),
			Type: reflect.TypeFor[string](),
		}})
	}
	return types
}

// The slots in front of the type cache hand every type its own metadata
// however many share a slot, and a slot stays with the first live type that
// claims it rather than changing hands on every call.
func TestTypeCacheSlotsShared(t *testing.T) {
	t.Parallel()

	c := newCache()
	types := slotTypes(3 * typeCacheSlots)
	infos := make([]*structInfo, len(types))
	for i, typ := range types {
		infos[i] = c.get(typ)
	}
	for range 2 {
		for i, typ := range types {
			if c.get(typ) != infos[i] {
				t.Fatalf("type %d: got different metadata on a repeat call", i)
			}
		}
	}
	for i, typ := range types {
		if f := infos[i].fields; len(f) != 1 || f[0].name != "F"+strconv.Itoa(i) {
			t.Fatalf("type %d: got the metadata of another type", i)
		}
		owner := slices.IndexFunc(types, func(u reflect.Type) bool { return typeSlot(u) == typeSlot(typ) })
		if e := c.slots[typeSlot(typ)].Load(); e == nil || e.typ != types[owner] {
			t.Fatalf("type %d: slot %d is not held by type %d, the first to claim it", i, typeSlot(typ), owner)
		}
	}
}

// A reconfiguration turns the slots' entries away, and an entry left from an
// older configuration gives its slot up to the current one.
func TestTypeCacheSlotsReset(t *testing.T) {
	t.Parallel()

	c := newCache()
	typ := slotTypes(1)[0]
	before := c.get(typ)

	c.l.Lock()
	c.reset()
	c.l.Unlock()
	for i := range c.slots {
		if c.slots[i].Load() != nil {
			t.Fatalf("slot %d survived the reset", i)
		}
	}
	after := c.get(typ)
	if after == before {
		t.Fatal("got the metadata built before the reset")
	}

	// An entry a build racing the reset stored after it.
	slot := &c.slots[typeSlot(typ)]
	slot.Store(&cacheEntry{info: before, typ: typ, gen: c.gen.Load() - 1})
	if c.get(typ) != after {
		t.Fatal("got the metadata of a stale slot entry")
	}
	if e := slot.Load(); e == nil || e.info != after || e.gen != c.gen.Load() {
		t.Fatal("the stale slot entry was not replaced")
	}
}

// Concurrent lookups of types sharing slots, interleaved with
// reconfigurations, always return metadata of the type asked for.
func TestTypeCacheSlotsConcurrent(t *testing.T) {
	t.Parallel()

	c := newCache()
	types := slotTypes(2 * typeCacheSlots)
	var wg sync.WaitGroup
	for g := range 4 {
		wg.Go(func() {
			for n := range 200 {
				i := (g*7 + n) % len(types)
				if f := c.get(types[i]).fields; len(f) != 1 || f[0].name != "F"+strconv.Itoa(i) {
					t.Errorf("type %d: got the metadata of another type", i)
					return
				}
			}
		})
	}
	wg.Go(func() {
		for range 20 {
			c.l.Lock()
			c.reset()
			c.l.Unlock()
		}
	})
	wg.Wait()
}

// caseMask upper-cases the letters of s whose position has its bit set in mask.
func caseMask(s string, mask int) string {
	b := []byte(s)
	for i := range b {
		if mask&(1<<i) != 0 && b[i] >= 'a' && b[i] <= 'z' {
			b[i] -= 'a' - 'A'
		}
	}
	return string(b)
}

// The cache is keyed by the path as sent, and "ITEMS.0001.value" is "items.1.value":
// the client can spell one path as many ways as it likes, so the cache has to
// stop taking them, while every spelling goes on decoding.
func TestPathCacheIsBoundedBySpellings(t *testing.T) {
	t.Parallel()

	type Item struct {
		Value string `schema:"value"`
	}
	type S struct {
		Items []Item `schema:"items"`
	}
	d := NewDecoder()
	info := d.cache.get(reflect.TypeFor[S]())

	// More spellings than the cache holds, none longer than it keeps.
	const pads, masks = 200, 128
	if pads*masks <= maxFastPaths+maxSpillPaths {
		t.Fatal("the test sends fewer spellings than the cache holds")
	}
	for pad := range pads {
		for mask := range masks {
			key := caseMask("items", mask) + "." + strings.Repeat("0", pad) + "1." + caseMask("value", mask>>5)
			var s S
			if err := d.Decode(&s, map[string][]string{key: {"x"}}); err != nil {
				t.Fatalf("%q: %v", key, err)
			}
			if len(s.Items) != 2 || s.Items[1].Value != "x" {
				t.Fatalf("%q decoded to %+v", key, s.Items)
			}
		}
	}
	if got := len(cachedPathKeys(&info.paths)); got < maxFastPaths || got > maxFastPaths+maxSpillPaths {
		t.Fatalf("cache holds %d paths, want between %d and %d", got, maxFastPaths, maxFastPaths+maxSpillPaths)
	}
	if !info.paths.full() {
		t.Fatal("the cache should have filled")
	}
}

// A key too long to be a path of the struct's is parsed but not kept.
func TestPathCacheSkipsLongKeys(t *testing.T) {
	t.Parallel()

	type Item struct {
		Value string `schema:"value"`
	}
	type S struct {
		Items []Item `schema:"items"`
	}
	d := NewDecoder()
	info := d.cache.get(reflect.TypeFor[S]())

	long := "items." + strings.Repeat("0", maxCachedPathLen) + "1.value"
	var s S
	if err := d.Decode(&s, map[string][]string{long: {"x"}}); err != nil {
		t.Fatal(err)
	}
	if len(s.Items) != 2 || s.Items[1].Value != "x" {
		t.Fatalf("decoded to %+v", s.Items)
	}
	if keys := cachedPathKeys(&info.paths); len(keys) != 0 {
		t.Fatalf("cache kept %d paths of a key of %d bytes", len(keys), len(long))
	}

	// A key at the limit is kept.
	exact := "items." + strings.Repeat("0", maxCachedPathLen-len("items.1.value")) + "1.value"
	if len(exact) != maxCachedPathLen {
		t.Fatalf("test key is %d bytes, want %d", len(exact), maxCachedPathLen)
	}
	if err := d.Decode(&s, map[string][]string{exact: {"y"}}); err != nil {
		t.Fatal(err)
	}
	if keys := cachedPathKeys(&info.paths); len(keys) != 1 {
		t.Fatalf("cache holds %d paths, want the one at the limit", len(keys))
	}
}

// Clearing the cache lets it take paths again, spill included.
func TestPathCacheClearReopensSpill(t *testing.T) {
	t.Parallel()

	var c pathCache
	for i := range maxFastPaths + 8 {
		c.loadOrStore(strconv.Itoa(i), nil)
	}
	if !c.sealed.Load() || c.spillLen.Load() != 8 {
		t.Fatalf("sealed=%v spillLen=%d, want sealed with 8 spilled", c.sealed.Load(), c.spillLen.Load())
	}
	c.clear()
	if c.sealed.Load() || c.spillLen.Load() != 0 || c.full() {
		t.Fatalf("clear left sealed=%v spillLen=%d full=%v", c.sealed.Load(), c.spillLen.Load(), c.full())
	}
}

// The spill stops at its cap whoever fills it, and every path still resolves.
func TestPathCacheSpillStopsAtCap(t *testing.T) {
	t.Parallel()

	var c pathCache
	for i := range maxFastPaths + maxSpillPaths {
		c.loadOrStore(strconv.Itoa(i), nil)
	}
	if !c.full() {
		t.Fatalf("cache of %d spilled paths is not full", c.spillLen.Load())
	}
	if got := len(cachedPathKeys(&c)); got != maxFastPaths+maxSpillPaths {
		t.Fatalf("cache holds %d paths, want %d", got, maxFastPaths+maxSpillPaths)
	}
}

// A path through a map field named by its bare alias is the same for every key
// the client puts after it, so it is shared rather than parsed for each.
func TestMapEntryPathIsShared(t *testing.T) { //nolint:paralleltest // testing.AllocsPerRun panics in a parallel test
	type S struct {
		Attrs map[string]string `schema:"attrs"`
		Other map[string]int    `schema:"other"`
		Name  string            `schema:"name"`
	}
	d := NewDecoder()
	info := d.cache.get(reflect.TypeFor[S]())

	parse := func(key string) []pathPart {
		t.Helper()
		parts, err := d.cache.parsePathInfo(key, info)
		if err != nil {
			t.Fatalf("parsePathInfo(%q): %v", key, err)
		}
		return parts
	}
	a, b, c, o := parse("attrs.a"), parse("ATTRS.b.c"), parse("attrs.zzzzzz"), parse("other.x")
	if &a[0] != &b[0] || &a[0] != &c[0] {
		t.Error("the keys of one map field do not share a path")
	}
	if &a[0] == &o[0] || a[0].field == o[0].field {
		t.Error("two map fields share a path")
	}
	if a[0].keyStart != 0 || a[0].field.mapValue == nil {
		t.Errorf("shared path: keyStart=%d, mapValue=%v; want 0 and a map field", a[0].keyStart, a[0].field.mapValue)
	}
	if allocs := testing.AllocsPerRun(100, func() { _, _ = d.cache.parsePathInfo("attrs.colour", info) }); allocs != 0 {
		t.Errorf("parsePathInfo of a map entry allocated %v times", allocs)
	}
	if got := len(cachedPathKeys(&info.paths)); got != 0 {
		t.Errorf("the cache kept %d map entry paths", got)
	}

	// Not a map entry: a bare alias, a missing key, a field that is not a map.
	for _, key := range []string{"attrs", "attrs.", ".a", "name.x"} {
		if parts, err := d.cache.parsePathInfo(key, info); err == nil {
			if parts[0].field.mapValue != nil {
				t.Errorf("parsePathInfo(%q) named a map entry", key)
			}
		}
	}
}

// A key that names no field is rejected before the parser reserves anything.
func TestUnknownPathDoesNotAllocate(t *testing.T) { //nolint:paralleltest // testing.AllocsPerRun panics in a parallel test
	type Item struct {
		Value string `schema:"value"`
	}
	type S struct {
		Name  string `schema:"name"`
		Items []Item `schema:"items"`
	}
	d := NewDecoder()
	info := d.cache.get(reflect.TypeFor[S]())

	for _, key := range []string{"nope.x", "utm.source.medium", "NOPE.0.value", "nope.", "a.b.c.d.e.f.g.h.i.j.k.l.m.n.o.p.q.r.s.t"} {
		if _, err := d.cache.parsePathInfo(key, info); err != errInvalidPath { //nolint:errorlint // the sentinel itself, as decodeKey compares it
			t.Errorf("parsePathInfo(%q) = %v, want errInvalidPath", key, err)
		}
		if allocs := testing.AllocsPerRun(100, func() { _, _ = d.cache.parsePathInfo(key, info) }); allocs != 0 {
			t.Errorf("parsePathInfo(%q) allocated %v times", key, allocs)
		}
	}
}
