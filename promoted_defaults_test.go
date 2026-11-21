package schema

import (
	"reflect"
	"strings"
	"testing"
)

// See https://github.com/gofiber/schema/issues/117.
func TestDefaultsOfPromotedFieldsAppliedOnlyWhenMissing(t *testing.T) {
	t.Parallel()

	type Paging struct {
		Page int `schema:"page,default:1"`
	}
	type ByValue struct {
		Paging
		Limit int `schema:"limit,default:50"`
	}
	type ByPointer struct {
		*Paging
		Limit int `schema:"limit,default:50"`
	}

	for _, key := range []string{"page", "Paging.page"} {
		src := map[string][]string{key: {"0"}, "limit": {"0"}}

		var v ByValue
		if err := NewDecoder().Decode(&v, src); err != nil {
			t.Fatal(err)
		}
		if v.Page != 0 || v.Limit != 0 {
			t.Errorf("%s=0 by value: got page=%d limit=%d, want 0 and 0", key, v.Page, v.Limit)
		}

		var p ByPointer
		if err := NewDecoder().Decode(&p, src); err != nil {
			t.Fatal(err)
		}
		if p.Paging == nil || p.Page != 0 || p.Limit != 0 {
			t.Errorf("%s=0 by pointer: got %+v limit=%d, want 0 and 0", key, p.Paging, p.Limit)
		}
	}

	var v ByValue
	if err := NewDecoder().Decode(&v, map[string][]string{}); err != nil {
		t.Fatal(err)
	}
	if v.Page != 1 || v.Limit != 50 {
		t.Errorf("missing: got page=%d limit=%d, want 1 and 50", v.Page, v.Limit)
	}
}

func TestDefaultsOfShadowedFieldsKeepDottedPath(t *testing.T) {
	t.Parallel()

	type Emb struct {
		P int `schema:"p,default:7"`
	}
	type Out struct {
		Emb
		P int `schema:"p,default:9"`
	}

	var a Out
	if err := NewDecoder().Decode(&a, map[string][]string{"p": {"0"}}); err != nil {
		t.Fatal(err)
	}
	if a.P != 0 || a.Emb.P != 7 {
		t.Errorf("p=0: got outer=%d emb=%d, want 0 and 7", a.P, a.Emb.P)
	}

	var b Out
	if err := NewDecoder().Decode(&b, map[string][]string{"Emb.p": {"0"}}); err != nil {
		t.Fatal(err)
	}
	if b.P != 9 || b.Emb.P != 0 {
		t.Errorf("Emb.p=0: got outer=%d emb=%d, want 9 and 0", b.P, b.Emb.P)
	}
}

func TestDefaultsOfStructNestedInEmbeddedRespectPromotedKey(t *testing.T) {
	t.Parallel()

	type Inner struct {
		X int `schema:"x,default:3"`
	}
	type Paging struct {
		Inner Inner `schema:"inner"`
	}
	type Outer struct {
		Paging
	}

	for _, key := range []string{"inner.x", "Paging.inner.x"} {
		var v Outer
		if err := NewDecoder().Decode(&v, map[string][]string{key: {"0"}}); err != nil {
			t.Fatal(err)
		}
		if v.Inner.X != 0 {
			t.Errorf("%s=0: got %d, want 0", key, v.Inner.X)
		}
	}

	var v Outer
	if err := NewDecoder().Decode(&v, map[string][]string{}); err != nil {
		t.Fatal(err)
	}
	if v.Inner.X != 3 {
		t.Errorf("missing: got %d, want 3", v.Inner.X)
	}
}

func TestDefaultsOfDoubleEmbeddedFieldAppliedOnlyWhenMissing(t *testing.T) {
	t.Parallel()

	type B struct {
		X int `schema:"x,default:5"`
	}
	type A struct{ B }
	type Outer struct{ A }

	for _, key := range []string{"x", "A.B.x"} {
		var v Outer
		if err := NewDecoder().Decode(&v, map[string][]string{key: {"0"}}); err != nil {
			t.Fatal(err)
		}
		if v.X != 0 {
			t.Errorf("%s=0: got %d, want 0", key, v.X)
		}
	}
}

func TestDefaultsOfFieldShadowedTwoLevelsDown(t *testing.T) {
	t.Parallel()

	type In struct {
		X int `schema:"x,default:3"`
	}
	type B struct {
		Inner In `schema:"inner"`
	}
	type A struct{ B }
	type Outer struct {
		A
		Inner In `schema:"inner"`
	}

	var v Outer
	if err := NewDecoder().Decode(&v, map[string][]string{"inner.x": {"0"}}); err != nil {
		t.Fatal(err)
	}
	if v.Inner.X != 0 || v.B.Inner.X != 3 {
		t.Errorf("inner.x=0: got outer=%d deep=%d, want 0 and 3", v.Inner.X, v.B.Inner.X)
	}
}

func TestDefaultsOfShadowedNamedStructIgnoreBareKey(t *testing.T) {
	t.Parallel()

	type In struct {
		X int `schema:"x,default:3"`
	}
	type Emb struct {
		Inner In `schema:"inner"`
	}
	type Out struct {
		Emb
		Inner In `schema:"inner"`
	}

	d := NewDecoder()
	d.IgnoreUnknownKeys(true)
	var v Out
	if err := d.Decode(&v, map[string][]string{"x": {"0"}}); err != nil {
		t.Fatal(err)
	}
	if v.Inner.X != 3 || v.Emb.Inner.X != 3 {
		t.Errorf("x=0: got outer=%d emb=%d, want 3 and 3", v.Inner.X, v.Emb.Inner.X)
	}
}

func TestDefaultsOfDeepEmbeddedFieldsNotPromotedByMiddle(t *testing.T) {
	t.Parallel()

	type B struct {
		X int `schema:"x,default:5"`
	}
	type C struct {
		X int `schema:"x,default:6"`
	}
	type ShadowedInA struct {
		B
		X int `schema:"x,default:7"`
	}
	type AmbiguousInA struct {
		B
		C
	}

	var s struct{ ShadowedInA }
	if err := NewDecoder().Decode(&s, map[string][]string{}); err != nil {
		t.Fatal(err)
	}
	if s.X != 7 || s.B.X != 5 {
		t.Errorf("shadowed: got a=%d b=%d, want 7 and 5", s.X, s.B.X)
	}

	var sx struct{ ShadowedInA }
	if err := NewDecoder().Decode(&sx, map[string][]string{"x": {"0"}}); err != nil {
		t.Fatal(err)
	}
	if sx.X != 0 || sx.B.X != 5 {
		t.Errorf("shadowed x=0: got a=%d b=%d, want 0 and 5", sx.X, sx.B.X)
	}

	var sb struct{ ShadowedInA }
	if err := NewDecoder().Decode(&sb, map[string][]string{"ShadowedInA.B.x": {"0"}}); err != nil {
		t.Fatal(err)
	}
	if sb.X != 7 || sb.B.X != 0 {
		t.Errorf("shadowed ShadowedInA.B.x=0: got a=%d b=%d, want 7 and 0", sb.X, sb.B.X)
	}

	var a struct{ AmbiguousInA }
	if err := NewDecoder().Decode(&a, map[string][]string{}); err != nil {
		t.Fatal(err)
	}
	if a.B.X != 5 || a.C.X != 6 {
		t.Errorf("ambiguous: got b=%d c=%d, want 5 and 6", a.B.X, a.C.X)
	}
}

func TestDefaultsOfDeepEmbeddedPointerAllocated(t *testing.T) {
	t.Parallel()

	type D struct {
		Z int `schema:"z,default:1"`
	}
	type B struct{ *D }
	type A struct{ B }
	type Outer struct{ A }

	var v Outer
	if err := NewDecoder().Decode(&v, map[string][]string{}); err != nil {
		t.Fatal(err)
	}
	if v.D == nil || v.Z != 1 {
		t.Errorf("got %+v, want allocated D with z=1", v.D)
	}
}

type defaultLeaf struct {
	index []int
	segs  []string
}

// collectDefaultLeaves lists the int fields with a default reachable through
// struct values and embedded pointers, with their declared segment chain.
func collectDefaultLeaves(t reflect.Type, index []int, segs []string, out *[]defaultLeaf) {
	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)
		name, opts, _ := strings.Cut(f.Tag.Get("schema"), ",")
		if name == "-" || !f.IsExported() {
			continue
		}
		if name == "" {
			name = f.Name
		}
		idx := append(append([]int{}, index...), i)
		sg := append(append([]string{}, segs...), name)
		ft := f.Type
		if f.Anonymous && ft.Kind() == reflect.Pointer {
			ft = ft.Elem()
		}
		switch {
		case ft.Kind() == reflect.Struct:
			collectDefaultLeaves(ft, idx, sg, out)
		case ft.Kind() == reflect.Int && strings.Contains(opts, "default:"):
			*out = append(*out, defaultLeaf{idx, sg})
		}
	}
}

func leafInts(v reflect.Value, leaves []defaultLeaf) []int64 {
	out := make([]int64, len(leaves))
	for i, l := range leaves {
		f, err := v.FieldByIndexErr(l.index)
		if err != nil {
			out[i] = -1
			continue
		}
		out[i] = f.Int()
	}
	return out
}

// Every key the decoder resolves to a defaulted field must keep a provided
// zero there, whichever embedded segments it spells out or skips.
func TestDefaultsKeepProvidedZeroForEveryResolvingKey(t *testing.T) {
	t.Parallel()

	type Paging struct {
		Page int `schema:"page,default:1"`
	}
	type In struct {
		X int `schema:"x,default:3"`
	}
	type B struct {
		X int `schema:"x,default:5"`
	}
	type C struct {
		X int `schema:"x,default:6"`
	}
	type BIn struct {
		Inner In `schema:"inner"`
	}
	type D struct {
		Z int `schema:"z,default:1"`
	}
	type BD struct{ *D }
	type Emb struct {
		P int `schema:"p,default:7"`
	}
	type ShadowedInA struct {
		B
		X int `schema:"x,default:7"`
	}
	type AmbiguousInA struct {
		B
		C
	}
	type AB struct{ B }
	type ABIn struct{ BIn }
	type ABD struct{ BD }

	shapes := []any{
		struct {
			Paging
			Limit int `schema:"limit,default:50"`
		}{},
		struct{ *Paging }{},
		struct {
			Emb
			P int `schema:"p,default:9"`
		}{},
		struct{ BIn }{},
		struct{ AB }{},
		struct{ ABIn }{},
		struct {
			ABIn
			Inner In `schema:"inner"`
		}{},
		struct {
			BIn
			Inner In `schema:"inner"`
		}{},
		struct{ ShadowedInA }{},
		struct{ AmbiguousInA }{},
		struct{ ABD }{},
	}

	for _, shape := range shapes {
		typ := reflect.TypeOf(shape)
		var leaves []defaultLeaf
		collectDefaultLeaves(typ, nil, nil, &leaves)

		decode := func(src map[string][]string) []int64 {
			d := NewDecoder()
			d.IgnoreUnknownKeys(true)
			v := reflect.New(typ)
			if err := d.Decode(v.Interface(), src); err != nil {
				t.Fatalf("%v %v: %v", typ, src, err)
			}
			return leafInts(v.Elem(), leaves)
		}
		defaults := decode(map[string][]string{})

		for _, l := range leaves {
			last := len(l.segs) - 1
			for mask := 0; mask < 1<<last; mask++ {
				var segs []string
				for i := range last {
					if mask&(1<<i) != 0 {
						segs = append(segs, l.segs[i])
					}
				}
				key := strings.Join(append(segs, l.segs[last]), ".")

				landed := -1
				for i, got := range decode(map[string][]string{key: {"42"}}) {
					if got == 42 {
						landed = i
					}
				}
				if landed < 0 {
					continue
				}
				got := decode(map[string][]string{key: {"0"}})
				for i := range leaves {
					want := defaults[i]
					if i == landed {
						want = 0
					}
					if got[i] != want {
						t.Errorf("%v %s=0: field %v got %d, want %d", typ, key, leaves[i].segs, got[i], want)
					}
				}
			}
		}
	}
}

func TestDefaultsBehindPromotedPointerRespectEveryKey(t *testing.T) {
	t.Parallel()

	type Node struct {
		V int `schema:"v,default:1"`
	}
	type Emb struct {
		Next *Node `schema:"next"`
	}
	type Out struct{ Emb }

	for _, key := range []string{"next.v", "Emb.next.v"} {
		var v Out
		if err := NewDecoder().Decode(&v, map[string][]string{key: {"0"}}); err != nil {
			t.Fatal(err)
		}
		if v.Next == nil || v.Next.V != 0 {
			t.Errorf("%s=0: got %+v, want v=0", key, v.Next)
		}
	}
}
