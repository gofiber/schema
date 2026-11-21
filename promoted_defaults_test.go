package schema

import "testing"

// Defaults on fields promoted from an embedded struct must be applied only
// when the promoted key is missing, just like a field declared directly.
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

	provided := map[string][]string{"page": {"0"}, "limit": {"0"}}

	var v ByValue
	if err := NewDecoder().Decode(&v, provided); err != nil {
		t.Fatal(err)
	}
	if v.Page != 0 || v.Limit != 0 {
		t.Errorf("embedded by value: got page=%d limit=%d, want 0 and 0", v.Page, v.Limit)
	}

	var p ByPointer
	if err := NewDecoder().Decode(&p, provided); err != nil {
		t.Fatal(err)
	}
	if p.Paging == nil || p.Page != 0 || p.Limit != 0 {
		t.Errorf("embedded by pointer: got %+v limit=%d, want 0 and 0", p.Paging, p.Limit)
	}
}

// A shadowed field keeps its own dotted-path semantics: the promoted copy in
// the outer struct is unreachable, so the embedded walk's default applies.
func TestDefaultsOfShadowedFieldsKeepDottedPath(t *testing.T) {
	t.Parallel()

	type Emb struct {
		P int `schema:"p,default:7"`
	}
	type Out struct {
		Emb
		P int `schema:"p,default:9"`
	}

	// p=0: the outer field is provided; the shadowed embedded field is not.
	var a Out
	if err := NewDecoder().Decode(&a, map[string][]string{"p": {"0"}}); err != nil {
		t.Fatal(err)
	}
	if a.P != 0 || a.Emb.P != 7 {
		t.Errorf("p=0: got outer=%d emb=%d, want 0 and 7", a.P, a.Emb.P)
	}

	// Emb.p=0: the shadowed embedded field is provided; the outer is not.
	var b Out
	if err := NewDecoder().Decode(&b, map[string][]string{"Emb.p": {"0"}}); err != nil {
		t.Fatal(err)
	}
	if b.P != 9 || b.Emb.P != 0 {
		t.Errorf("Emb.p=0: got outer=%d emb=%d, want 9 and 0", b.P, b.Emb.P)
	}
}

// A struct nested inside the embedded one is promoted too: its key is inner.x,
// not Paging.inner.x.
func TestDefaultsOfFieldsNestedInEmbeddedRespectPromotedKey(t *testing.T) {
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

	var v Outer
	if err := NewDecoder().Decode(&v, map[string][]string{"inner.x": {"0"}}); err != nil {
		t.Fatal(err)
	}
	if v.Paging.Inner.X != 0 {
		t.Errorf("inner.x=0: got %d, want 0", v.Paging.Inner.X)
	}
}

// A field promoted two levels deep is still provided by its bare key.
func TestDefaultsOfDoubleEmbeddedFieldAppliedOnlyWhenMissing(t *testing.T) {
	t.Parallel()

	type B struct {
		X int `schema:"x,default:5"`
	}
	type A struct {
		B
	}
	type Outer struct {
		A
	}

	var v Outer
	if err := NewDecoder().Decode(&v, map[string][]string{"x": {"0"}}); err != nil {
		t.Fatal(err)
	}
	if v.A.B.X != 0 {
		t.Errorf("x=0: got %d, want 0", v.A.B.X)
	}
}
