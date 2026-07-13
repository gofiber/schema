package schema

import (
	"testing"

	utils "github.com/gofiber/utils/v2"
)

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

	d.cache.l.RLock()
	defer d.cache.l.RUnlock()
	if len(d.cache.pathCache) != 1 {
		t.Fatalf("expected 1 cached path, got %d", len(d.cache.pathCache))
	}
	for k := range d.cache.pathCache {
		if k.path != "entity" {
			t.Fatalf("pathCache key mutated to %q; key must be cloned before caching", k.path)
		}
	}
}
