package schema

import (
	"reflect"
	"strings"
	"testing"
	"time"
)

// Keys served by the direct-path map must behave exactly like the generic
// parser: case-insensitive, covering flat aliases and nested dotted chains.
func TestDirectPathLookup(t *testing.T) {
	type Inner struct {
		Value string `schema:"value"`
	}
	type Outer struct {
		Name   string `schema:"name"`
		Nested Inner  `schema:"nested"`
	}

	for _, keys := range []struct {
		name, nested string
	}{
		{"name", "nested.value"},
		{"NAME", "NESTED.VALUE"},
		{"NaMe", "NeStEd.VaLuE"},
	} {
		var s Outer
		data := map[string][]string{
			keys.name:   {"x"},
			keys.nested: {"y"},
		}
		if err := NewDecoder().Decode(&s, data); err != nil {
			t.Fatalf("Decode(%q, %q): %v", keys.name, keys.nested, err)
		}
		if s.Name != "x" || s.Nested.Value != "y" {
			t.Fatalf("Decode(%q, %q) = %+v, want Name=x Nested.Value=y", keys.name, keys.nested, s)
		}
	}
}

// Direct-path entries exist for statically-resolvable keys only; everything
// else must keep flowing through the generic parser unchanged.
func TestDirectPathLookupFallbacks(t *testing.T) {
	// Keys longer than the case-fold buffer take the generic path.
	longAlias := strings.Repeat("a", maxDirectKeyLen+8)
	type Long struct {
		Field string `schema:"aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"`
	}
	if len(longAlias) != maxDirectKeyLen+8 {
		t.Fatal("bad test setup")
	}
	var l Long
	if err := NewDecoder().Decode(&l, map[string][]string{longAlias: {"v"}}); err != nil {
		t.Fatalf("long key: %v", err)
	}
	if l.Field != "v" {
		t.Fatalf("long key decoded %+v", l)
	}

	// A bare slice-of-structs alias still needs a slice index; the direct map
	// must not make it valid.
	type Item struct {
		Value string `schema:"value"`
	}
	type WithItems struct {
		Items []Item `schema:"items"`
	}
	var w WithItems
	err := NewDecoder().Decode(&w, map[string][]string{"items": {"x"}})
	if err == nil {
		t.Fatal("bare slice-of-structs alias must remain an invalid path")
	}
	if _, ok := err.(MultiError)["items"].(UnknownKeyError); !ok {
		t.Fatalf("want UnknownKeyError, got %v", err)
	}

	// Pointer-to-struct chains are not precomputed but must keep decoding.
	type Inner struct {
		Value string `schema:"value"`
	}
	type WithPtr struct {
		Nested *Inner `schema:"nested"`
	}
	var p WithPtr
	if err := NewDecoder().Decode(&p, map[string][]string{"NESTED.value": {"y"}}); err != nil {
		t.Fatalf("pointer chain: %v", err)
	}
	if p.Nested == nil || p.Nested.Value != "y" {
		t.Fatalf("pointer chain decoded %+v", p)
	}
}

// The native slice decode path must keep the generic path's semantics: comma
// splitting, zeroEmpty, all-or-nothing assignment, ConversionError details.
func TestNativeSliceDecode(t *testing.T) {
	type S struct {
		Tags   []string  `schema:"tags"`
		IDs    []int     `schema:"ids"`
		Scores []float64 `schema:"scores"`
		Flags  []bool    `schema:"flags"`
		Big    []int64   `schema:"big"`
		U      []uint    `schema:"u"`
		U64    []uint64  `schema:"u64"`
	}

	var s S
	data := map[string][]string{
		"tags":   {"a", "b,c", ""}, // strings never split on commas
		"ids":    {"1,2", "3"},
		"scores": {"1.5", "2.5"},
		"flags":  {"true", "on"},
		"big":    {"9007199254740993"},
		"u":      {"7"},
		"u64":    {"18446744073709551615"},
	}
	if err := NewDecoder().Decode(&s, data); err != nil {
		t.Fatal(err)
	}
	want := S{
		Tags:   []string{"a", "b,c"},
		IDs:    []int{1, 2, 3},
		Scores: []float64{1.5, 2.5},
		Flags:  []bool{true, true},
		Big:    []int64{9007199254740993},
		U:      []uint{7},
		U64:    []uint64{18446744073709551615},
	}
	if !reflect.DeepEqual(s, want) {
		t.Fatalf("got %+v, want %+v", s, want)
	}

	// zeroEmpty appends zero values for empty items.
	var z S
	d := NewDecoder()
	d.ZeroEmpty(true)
	if err := d.Decode(&z, map[string][]string{"ids": {"1,,2", ""}}); err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(z.IDs, []int{1, 0, 2, 0}) {
		t.Fatalf("zeroEmpty got %v", z.IDs)
	}

	// A parse failure must leave the field untouched and carry the value's
	// index in the ConversionError.
	pre := S{IDs: []int{42}}
	err := NewDecoder().Decode(&pre, map[string][]string{"ids": {"1", "oops"}})
	if err == nil {
		t.Fatal("want conversion error")
	}
	convErr, ok := err.(MultiError)["ids"].(ConversionError)
	if !ok {
		t.Fatalf("want ConversionError, got %v", err)
	}
	if convErr.Index != 1 {
		t.Fatalf("want Index=1, got %d", convErr.Index)
	}
	if !reflect.DeepEqual(pre.IDs, []int{42}) {
		t.Fatalf("field must stay untouched on error, got %v", pre.IDs)
	}
}

// Named slice types take the generic reflect loop, not the native path; its
// semantics must match the native path's exactly.
func TestGenericSliceNamedTypes(t *testing.T) {
	type IDs []int
	type Tags []string
	type S struct {
		IDs  IDs  `schema:"ids"`
		Tags Tags `schema:"tags"`
	}

	var s S
	data := map[string][]string{
		"ids":  {"1,2", "3"},
		"tags": {"a", "b,c"},
	}
	if err := NewDecoder().Decode(&s, data); err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(s.IDs, IDs{1, 2, 3}) || !reflect.DeepEqual(s.Tags, Tags{"a", "b,c"}) {
		t.Fatalf("got %+v", s)
	}

	// zeroEmpty fills zero slots; empties are dropped without it.
	var z S
	d := NewDecoder()
	d.ZeroEmpty(true)
	if err := d.Decode(&z, map[string][]string{"ids": {"1,,2", ""}}); err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(z.IDs, IDs{1, 0, 2, 0}) {
		t.Fatalf("zeroEmpty got %v", z.IDs)
	}
	var trunc S
	if err := NewDecoder().Decode(&trunc, map[string][]string{"ids": {"1,,2", ""}}); err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(trunc.IDs, IDs{1, 2}) {
		t.Fatalf("truncation got %v", trunc.IDs)
	}

	// Parse failures, both inside a comma list and as a plain value.
	for _, bad := range []string{"1,x", "x"} {
		var b S
		err := NewDecoder().Decode(&b, map[string][]string{"ids": {bad}})
		if _, ok := err.(MultiError)["ids"].(ConversionError); !ok {
			t.Fatalf("value %q: want ConversionError, got %v", bad, err)
		}
	}
}

// Fields the fast scalar path skips (pointers) keep their generic behavior.
func TestGenericScalarFallbacks(t *testing.T) {
	type S struct {
		N *int `schema:"n"`
	}

	// Empty value + ZeroEmpty on a pointer field takes the generic zeroing.
	d := NewDecoder()
	d.ZeroEmpty(true)
	var s S
	if err := d.Decode(&s, map[string][]string{"n": {""}}); err != nil {
		t.Fatal(err)
	}
	if s.N == nil || *s.N != 0 {
		t.Fatalf("want zeroed *int, got %+v", s.N)
	}
}

// Encoding into a fresh dst must handle every appendValue branch: duplicate
// aliases appending to a scratch-backed entry, caller-derived strings and
// long formatter output bypassing the shared scratch, and short numeric
// values using it.
func TestEncodeDuplicateAliasFreshDst(t *testing.T) {
	type S struct {
		A int     `schema:"k"`
		B int     `schema:"k"`
		C string  `schema:"c"`
		L float64 `schema:"l"`
	}
	dst := map[string][]string{}
	// 1e300 in 'f' format is far longer than maxScratchValueLen.
	if err := NewEncoder().Encode(S{A: 1, B: 2, C: "3", L: 1e300}, dst); err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(dst["k"], []string{"1", "2"}) || !reflect.DeepEqual(dst["c"], []string{"3"}) {
		t.Fatalf("got %v", dst)
	}
	if len(dst["l"]) != 1 || len(dst["l"][0]) <= maxScratchValueLen {
		t.Fatalf("long value got %v", dst["l"])
	}
}

type sliceHeavyStruct struct {
	Tags   []string  `schema:"tags"`
	IDs    []int     `schema:"ids"`
	Scores []float64 `schema:"scores"`
	Flags  []bool    `schema:"flags"`
}

func BenchmarkSliceHeavyDecode(b *testing.B) {
	data := map[string][]string{
		"tags":   {"alpha", "beta", "gamma", "delta", "epsilon", "zeta", "eta", "theta", "iota", "kappa", "lambda", "mu", "nu", "xi", "omicron", "pi"},
		"ids":    {"1,2,3,4,5,6,7,8", "9,10,11,12,13,14,15,16"},
		"scores": {"1.5", "2.5", "3.5", "4.5", "5.5", "6.5", "7.5", "8.5"},
		"flags":  {"true", "false", "true", "false", "true", "false", "true", "false"},
	}
	decoder := NewDecoder()
	s := &sliceHeavyStruct{}
	b.ReportAllocs()
	for b.Loop() {
		if err := decoder.Decode(s, data); err != nil {
			b.Fatal(err)
		}
	}
}

type fan0 struct {
	V string `schema:"v"`
}
type fan1 struct {
	A fan0 `schema:"a"`
	B fan0 `schema:"b"`
}
type fan2 struct {
	A fan1 `schema:"a"`
	B fan1 `schema:"b"`
}
type fan3 struct {
	A fan2 `schema:"a"`
	B fan2 `schema:"b"`
}
type fan4 struct {
	A fan3 `schema:"a"`
	B fan3 `schema:"b"`
}
type fan5 struct {
	A fan4 `schema:"a"`
	B fan4 `schema:"b"`
}
type fan6 struct {
	A fan5 `schema:"a"`
	B fan5 `schema:"b"`
}
type fan7 struct {
	A fan6 `schema:"a"`
	B fan6 `schema:"b"`
}
type fan8 struct {
	A fan7 `schema:"a"`
	B fan7 `schema:"b"`
}
type fan9 struct {
	A fan8 `schema:"a"`
	B fan8 `schema:"b"`
}
type fan10 struct {
	A fan9 `schema:"a"`
	B fan9 `schema:"b"`
}
type fan11 struct {
	A fan10 `schema:"a"`
	B fan10 `schema:"b"`
}
type fan12 struct {
	A fan11 `schema:"a"`
	B fan11 `schema:"b"`
}
type fan13 struct {
	A fan12 `schema:"a"`
	B fan12 `schema:"b"`
}
type fan14 struct {
	A fan13 `schema:"a"`
	B fan13 `schema:"b"`
}
type fan15 struct {
	A fan14 `schema:"a"`
	B fan14 `schema:"b"`
}
type fan16 struct {
	A fan15 `schema:"a"`
	B fan15 `schema:"b"`
}
type fan17 struct {
	A fan16 `schema:"a"`
	B fan16 `schema:"b"`
}
type fan18 struct {
	A fan17 `schema:"a"`
	B fan17 `schema:"b"`
}
type fan19 struct {
	A fan18 `schema:"a"`
	B fan18 `schema:"b"`
}

// Deep fan-out nesting has exponentially many dotted paths; the direct map
// must stay capped so the first Decode neither stalls nor retains huge maps.
func TestDirectPathsCappedForDeepFanout(t *testing.T) {
	d := NewDecoder()
	var s fan19
	deepKey := strings.Repeat("a.", 10) + strings.Repeat("b.", 9) + "v"
	data := map[string][]string{deepKey: {"deep"}}
	done := make(chan error, 1)
	go func() { done <- d.Decode(&s, data) }()
	select {
	case err := <-done:
		if err != nil {
			t.Fatal(err)
		}
	case <-time.After(30 * time.Second):
		t.Fatal("Decode stalled: direct-path precomputation not capped")
	}
	if s.A.A.A.A.A.A.A.A.A.A.B.B.B.B.B.B.B.B.B.V != "deep" {
		t.Fatal("deep key not decoded")
	}
	info := d.cache.get(reflect.TypeOf(s))
	if len(info.direct) > maxDirectPaths+2 {
		t.Fatalf("direct map has %d entries, want <= %d", len(info.direct), maxDirectPaths+2)
	}
}

// Unlike the other encode benchmarks, this one encodes into a fresh map each
// iteration, matching how callers build url.Values per request.
func BenchmarkEncodeFreshDst(b *testing.B) {
	type S struct {
		A string  `schema:"a"`
		B int     `schema:"b"`
		C bool    `schema:"c"`
		D float64 `schema:"d"`
		E []int   `schema:"e"`
		F string  `schema:"f,omitempty"`
	}
	s := S{A: "abc", B: 123, C: true, D: 3.14, E: []int{1, 2, 3}}
	enc := NewEncoder()
	b.ReportAllocs()
	for b.Loop() {
		vals := make(map[string][]string, 8)
		if err := enc.Encode(&s, vals); err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkMixedCaseKeyDecode(b *testing.B) {
	type S struct {
		FirstName string `schema:"firstName"`
		LastName  string `schema:"lastName"`
		Age       int    `schema:"age"`
	}
	data := map[string][]string{
		"FirstName": {"Grace"},
		"LASTNAME":  {"Hopper"},
		"age":       {"85"},
	}
	decoder := NewDecoder()
	s := &S{}
	for b.Loop() {
		if err := decoder.Decode(s, data); err != nil {
			b.Fatal(err)
		}
	}
}
