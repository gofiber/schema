package schema

import (
	"reflect"
	"testing"

	utils "github.com/gofiber/utils/v2"
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
		info.paths.Clear()
		if _, err := d.cache.parsePath("items.0.value", typ); err != nil {
			b.Fatal(err)
		}
	}
}

// Decode callers (e.g. Fiber binders) pass map keys aliasing reused request
// buffers; pathCache must clone the key so later buffer reuse cannot poison it.
func TestParsePathDetachesCacheKey(t *testing.T) {
	type S struct {
		Entity string `schema:"entity"`
	}
	d := NewDecoder()
	buf := []byte("entity")
	var s S
	if err := d.Decode(&s, map[string][]string{utils.UnsafeString(buf): {"x"}}); err != nil {
		t.Fatal(err)
	}
	copy(buf, "policy") // simulate fasthttp buffer reuse mutating the key bytes

	count := 0
	d.cache.get(reflect.TypeOf(s)).paths.Range(func(key, _ any) bool {
		count++
		if k := key.(string); k != "entity" {
			t.Fatalf("path cache key mutated to %q; key must be cloned before caching", k)
		}
		return true
	})
	if count != 1 {
		t.Fatalf("expected 1 cached path, got %d", count)
	}
}
