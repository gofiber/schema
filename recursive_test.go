package schema

import (
	"errors"
	"fmt"
	"reflect"
	"testing"
)

// SelfEmbedded embeds a pointer to its own type. Building its metadata
// expanded that pointer without end, until memory ran out.
type SelfEmbedded struct {
	*SelfEmbedded `schema:"parent"`
	V             int `schema:"v,default:3"`
	W             int `schema:"w,default:4"`
}

// An embedded pointer leading back to its struct is a named struct pointer:
// nothing is promoted through it or allocated for it, keys name its fields
// under its alias, and the defaults walk enters it once it is set.
func TestSelfEmbeddedPointerIsNamedField(t *testing.T) {
	t.Parallel()

	var r SelfEmbedded
	if err := NewDecoder().Decode(&r, map[string][]string{"v": {"1"}}); err != nil {
		t.Fatal(err)
	}
	if r.V != 1 || r.W != 4 || r.SelfEmbedded != nil {
		t.Errorf("got v=%d w=%d parent=%+v, want 1, 4 and nil", r.V, r.W, r.SelfEmbedded)
	}

	r = SelfEmbedded{}
	if err := NewDecoder().Decode(&r, map[string][]string{"parent.v": {"5"}}); err != nil {
		t.Fatal(err)
	}
	p := r.SelfEmbedded
	if r.V != 3 || r.W != 4 || p == nil || p.V != 5 || p.W != 4 || p.SelfEmbedded != nil {
		t.Errorf("got v=%d w=%d parent=%+v, want 3, 4 and {v=5 w=4 parent=nil}", r.V, r.W, p)
	}

	type requiredSelf struct {
		*requiredSelf
		Name string `schema:"name,required"`
	}
	var rs requiredSelf
	err := NewDecoder().Decode(&rs, map[string][]string{})
	var me MultiError
	if !errors.As(err, &me) || len(me) != 1 || !errors.As(me["name"], new(EmptyFieldError)) || rs.requiredSelf != nil {
		t.Errorf("got %v, embedded %v, want only name reported missing and nil", err, rs.requiredSelf)
	}

	// Unexported, so neither allocated nor reachable by key.
	type hiddenSelf struct {
		*hiddenSelf
		V int `schema:"v,default:1"`
	}
	var hs hiddenSelf
	if err := NewDecoder().Decode(&hs, map[string][]string{}); err != nil || hs.V != 1 || hs.hiddenSelf != nil {
		t.Errorf("got %v, v=%d, embedded=%v, want nil, 1 and nil", err, hs.V, hs.hiddenSelf)
	}
	// A key naming a field below it cannot allocate it: the key is passed
	// over, as one naming an unexported field is, and the others decode.
	hs = hiddenSelf{}
	if err := NewDecoder().Decode(&hs, map[string][]string{"hiddenSelf.v": {"5"}, "v": {"2"}}); err != nil || hs.V != 2 || hs.hiddenSelf != nil {
		t.Errorf("got %v, v=%d, embedded=%v, want nil, 2 and nil", err, hs.V, hs.hiddenSelf)
	}
}

type unexportedPtrTarget struct {
	V int `schema:"v"`
}

// A key below a nil unexported struct pointer, named or embedded, names what
// cannot be allocated. Allocating it panicked, which ended the decode with
// that panic as its error.
func TestKeyBelowUnexportedNilPointer(t *testing.T) {
	t.Parallel()

	type named struct {
		inner *unexportedPtrTarget
		X     int `schema:"x"`
	}
	var n named
	if err := NewDecoder().Decode(&n, map[string][]string{"inner.v": {"1"}, "x": {"2"}}); err != nil || n.X != 2 || n.inner != nil {
		t.Errorf("named: got %v, x=%d, inner=%v, want nil, 2 and nil", err, n.X, n.inner)
	}

	type embedded struct {
		*unexportedPtrTarget
		X int `schema:"x"`
	}
	var e embedded
	if err := NewDecoder().Decode(&e, map[string][]string{"unexportedPtrTarget.v": {"1"}, "x": {"2"}}); err != nil || e.X != 2 || e.unexportedPtrTarget != nil {
		t.Errorf("embedded: got %v, x=%d, embedded=%v, want nil, 2 and nil", err, e.X, e.unexportedPtrTarget)
	}

	// Below an unexported named pointer the caller set, nothing is settable:
	// not the field a key names, nor the embedded pointer on the way to it.
	type Part struct {
		W int `schema:"w"`
	}
	type holder struct {
		*Part
		V int `schema:"v"`
	}
	type element struct {
		inner *holder
	}
	type list struct {
		Items []element `schema:"items"`
		X     int       `schema:"x"`
	}
	l := list{Items: []element{{inner: &holder{}}}}
	src := map[string][]string{"items.0.inner.v": {"1"}, "items.0.inner.w": {"1"}, "x": {"2"}}
	if err := NewDecoder().Decode(&l, src); err != nil || l.X != 2 || l.Items[0].inner.V != 0 || l.Items[0].inner.Part != nil {
		t.Errorf("set by the caller: got %v, x=%d, inner=%+v, want nil, 2 and {Part:<nil> V:0}", err, l.X, l.Items[0].inner)
	}
}

type unsettableDefaults struct {
	V int   `schema:"v,default:1"`
	L []int `schema:"l,default:1|2"`
	P *int  `schema:"p,default:3"`
}

// UnsettablePart is exported, so a nil embedded pointer to it is allocated
// wherever it can be set.
type UnsettablePart struct {
	W int `schema:"w"`
}

type unsettableHolder struct {
	*UnsettablePart
	V int `schema:"v"`
}

// The defaults walk reaches structs below unexported named fields too: one
// the caller set under an unexported pointer, and an unexported struct value.
// Nothing can be set there, neither a default nor a nil embedded pointer, and
// trying to panicked, which ended the decode with that panic as its error.
func TestDefaultsWalkBelowUnexportedFields(t *testing.T) {
	t.Parallel()

	type named struct {
		inner *unsettableDefaults
		X     int `schema:"x,default:2"`
	}
	n := named{inner: &unsettableDefaults{}}
	if err := NewDecoder().Decode(&n, map[string][]string{}); err != nil || n.X != 2 || !reflect.DeepEqual(n.inner, &unsettableDefaults{}) {
		t.Errorf("pointer set by the caller: got %v, x=%d, inner=%+v, want nil, 2 and zero", err, n.X, n.inner)
	}

	type ptrHolder struct {
		inner *unsettableHolder
		X     int `schema:"x,default:2"`
	}
	p := ptrHolder{inner: &unsettableHolder{}}
	if err := NewDecoder().Decode(&p, map[string][]string{}); err != nil || p.X != 2 || p.inner.UnsettablePart != nil {
		t.Errorf("embedded pointer below a set pointer: got %v, x=%d, inner=%+v, want nil, 2 and {<nil> 0}", err, p.X, p.inner)
	}

	type valueHolder struct {
		inner unsettableHolder
		X     int `schema:"x,default:2"`
	}
	var v valueHolder
	if err := NewDecoder().Decode(&v, map[string][]string{"x": {"3"}}); err != nil || v.X != 3 || v.inner.UnsettablePart != nil {
		t.Errorf("embedded pointer below a struct value: got %v, x=%d, inner=%+v, want nil, 3 and {<nil> 0}", err, v.X, v.inner)
	}

	type direct struct {
		v int `schema:"v,default:1"`
		X int `schema:"x,default:2"`
	}
	var d direct
	if err := NewDecoder().Decode(&d, map[string][]string{}); err != nil || d.v != 0 || d.X != 2 {
		t.Errorf("unexported field: got %v, v=%d, x=%d, want nil, 0 and 2", err, d.v, d.X)
	}

	// An unexported embedded struct is another matter: the fields it promotes
	// can be set, so they take their defaults as ever, also through a pointer
	// the caller set.
	type promoted struct {
		V int `schema:"v,default:1"`
	}
	type viaValue struct {
		promoted
		X int `schema:"x"`
	}
	var vv viaValue
	if err := NewDecoder().Decode(&vv, map[string][]string{}); err != nil || vv.V != 1 {
		t.Errorf("embedded struct value: got %v, v=%d, want nil and 1", err, vv.V)
	}
	type viaPointer struct {
		*promoted
		X int `schema:"x"`
	}
	vp := viaPointer{promoted: &promoted{}}
	if err := NewDecoder().Decode(&vp, map[string][]string{}); err != nil || vp.V != 1 {
		t.Errorf("embedded pointer set by the caller: got %v, v=%d, want nil and 1", err, vp.V)
	}
}

// SelfEmbeddedPlain embeds a pointer to its own type and holds no default.
type SelfEmbeddedPlain struct {
	*SelfEmbeddedPlain `schema:"parent"`
	X                  int `schema:"x"`
}

// A struct that embeds a pointer to its own type and holds no default decodes
// as well: keys name the fields below the pointer under its alias, level by
// level, and nothing is allocated below the levels they name.
func TestSelfEmbeddedPointerWithoutDefaults(t *testing.T) {
	t.Parallel()

	var s SelfEmbeddedPlain
	if err := NewDecoder().Decode(&s, map[string][]string{"x": {"1"}, "parent.parent.x": {"2"}}); err != nil {
		t.Fatal(err)
	}
	if p := s.SelfEmbeddedPlain; s.X != 1 || p == nil || p.SelfEmbeddedPlain == nil || p.SelfEmbeddedPlain.X != 2 || p.SelfEmbeddedPlain.SelfEmbeddedPlain != nil {
		t.Errorf("got x=%d parent=%+v", s.X, p)
	}
}

// MutualA and MutualB embed pointers to each other.
type MutualA struct {
	*MutualB
	X int `schema:"x,default:1"`
}

type MutualB struct {
	*MutualA
	Y int `schema:"y,default:2"`
}

// Each of two structs embedding the other is built without end too. The
// metadata, and so what a key decodes to, must not depend on which of them
// is built first.
func TestMutuallyEmbeddedPointers(t *testing.T) {
	t.Parallel()

	var showA func(*MutualA) string
	showB := func(b *MutualB) string {
		if b == nil {
			return "nil"
		}
		return fmt.Sprintf("{y=%d a=%s}", b.Y, showA(b.MutualA))
	}
	showA = func(a *MutualA) string {
		if a == nil {
			return "nil"
		}
		return fmt.Sprintf("{x=%d b=%s}", a.X, showB(a.MutualB))
	}
	decodeA := func(d *Decoder) string {
		var a MutualA
		err := d.Decode(&a, map[string][]string{"x": {"5"}, "MutualB.y": {"6"}})
		return fmt.Sprintf("%v %s", err, showA(&a))
	}
	decodeB := func(d *Decoder) string {
		var b MutualB
		err := d.Decode(&b, map[string][]string{"MutualA.MutualB.y": {"0"}})
		return fmt.Sprintf("%v %s", err, showB(&b))
	}

	aFirst := NewDecoder()
	gotA, gotB := decodeA(aFirst), decodeB(aFirst)
	bFirst := NewDecoder()
	if b, a := decodeB(bFirst), decodeA(bFirst); a != gotA || b != gotB {
		t.Errorf("built A first: %s, %s; built B first: %s, %s", gotA, gotB, a, b)
	}
	if want := "<nil> {x=5 b={y=6 a=nil}}"; gotA != want {
		t.Errorf("into MutualA: %s, want %s", gotA, want)
	}
	// The nested MutualB is walked under its own prefix, where y was sent.
	if want := "<nil> {y=2 a={x=1 b={y=0 a=nil}}}"; gotB != want {
		t.Errorf("into MutualB: %s, want %s", gotB, want)
	}
}

// BackOuter embeds a pointer to BackInner, which holds a BackOuter value: a
// build comes back to BackOuter through a named field.
type BackOuter struct {
	*BackInner
	X int `schema:"x"`
}

type BackInner struct {
	Outer BackOuter `schema:"outer"`
	Y     int       `schema:"y,default:7"`
}

// ViaValue embeds a struct value that embeds a pointer back to ViaValue.
type ViaValue struct {
	ViaValueInner
	Z int `schema:"z"`
}

type ViaValueInner struct {
	*ViaValue `schema:"up"`
	Q         int `schema:"q,default:9"`
}

func TestEmbeddedPointerLeadingBackThroughOtherFields(t *testing.T) {
	t.Parallel()

	var o BackOuter
	if err := NewDecoder().Decode(&o, map[string][]string{"x": {"1"}, "BackInner.outer.x": {"2"}}); err != nil {
		t.Fatal(err)
	}
	if inner := o.BackInner; o.X != 1 || inner == nil || inner.Outer.X != 2 || inner.Y != 7 {
		t.Errorf("got %+v, inner %+v", o, inner)
	}

	var in BackInner
	if err := NewDecoder().Decode(&in, map[string][]string{"outer.x": {"3"}}); err != nil {
		t.Fatal(err)
	}
	if in.Outer.X != 3 || in.Outer.BackInner != nil || in.Y != 7 {
		t.Errorf("got %+v", in)
	}

	// The struct value is expanded, so its fields are promoted, the pointer
	// back among them as a named field.
	var v ViaValue
	if err := NewDecoder().Decode(&v, map[string][]string{"z": {"1"}, "up.z": {"2"}}); err != nil {
		t.Fatal(err)
	}
	if v.Z != 1 || v.Q != 9 || v.ViaValue == nil || v.ViaValue.Z != 2 || v.ViaValue.Q != 9 {
		t.Errorf("got z=%d q=%d up=%+v", v.Z, v.Q, v.ViaValue)
	}
}

// WholeSelf embeds a pointer to its own type and unmarshals itself, so a
// field of its type is decoded as a whole.
type WholeSelf struct {
	*WholeSelf
	V int `schema:"v,default:1"`
}

func (w *WholeSelf) UnmarshalText(text []byte) error {
	w.V = len(text)
	return nil
}

// WholeRingA holds a WholeRingB value, which unmarshals itself and embeds a
// pointer back to WholeRingA.
type WholeRingA struct {
	B WholeRingB `schema:"b"`
	X int        `schema:"x"`
}

type WholeRingB struct {
	*WholeRingA
	N int `schema:"n,default:2"`
}

func (b *WholeRingB) UnmarshalText(text []byte) error {
	b.N = len(text)
	return nil
}

// Nothing is promoted through a struct decoded as a whole, but the defaults
// plan descends into one, so an embedded pointer leading back through one was
// expanded without end as well, until the stack ran out.
func TestEmbeddedPointerLeadingBackThroughStructsDecodedWhole(t *testing.T) {
	t.Parallel()

	var w WholeSelf
	if err := NewDecoder().Decode(&w, map[string][]string{}); err != nil || w.V != 1 || w.WholeSelf != nil {
		t.Errorf("nothing sent: got %v, v=%d, embedded=%v, want nil, 1 and nil", err, w.V, w.WholeSelf)
	}
	// Its own key still decodes the embedded field as a whole.
	w = WholeSelf{}
	if err := NewDecoder().Decode(&w, map[string][]string{"WholeSelf": {"abc"}}); err != nil {
		t.Fatal(err)
	}
	if in := w.WholeSelf; w.V != 1 || in == nil || in.V != 3 || in.WholeSelf != nil {
		t.Errorf("key sent: got v=%d, embedded=%+v, want 1 and {v=3, embedded=nil}", w.V, in)
	}

	// Whichever of the two types is built first.
	for _, bFirst := range []bool{false, true} {
		d := NewDecoder()
		if bFirst {
			if err := d.Decode(&WholeRingB{}, map[string][]string{}); err != nil {
				t.Fatal(err)
			}
		}
		var a WholeRingA
		if err := d.Decode(&a, map[string][]string{"x": {"1"}}); err != nil {
			t.Fatal(err)
		}
		if a.X != 1 || a.B.N != 2 || a.B.WholeRingA != nil {
			t.Errorf("B first %v, into A: got x=%d, b=%+v, want 1 and {n=2, embedded=nil}", bFirst, a.X, a.B)
		}
		var b WholeRingB
		if err := d.Decode(&b, map[string][]string{"WholeRingA.x": {"5"}}); err != nil {
			t.Fatal(err)
		}
		if in := b.WholeRingA; b.N != 2 || in == nil || in.X != 5 || in.B.N != 2 || in.B.WholeRingA != nil {
			t.Errorf("B first %v, into B: got n=%d, embedded=%+v, want 2 and {x=5, b={n=2, embedded=nil}}", bFirst, b.N, in)
		}
	}
}

// unplannedInner holds an unplannedWhole value, which unmarshals itself and
// embeds a pointer back to unplannedInner. Neither holds a default or a
// pointer to allocate, so neither has a defaults plan.
type unplannedInner struct {
	B unplannedWhole `schema:"b"`
	X int            `schema:"x"`
}

type unplannedWhole struct {
	*unplannedInner
	N int `schema:"n"`
}

func (w *unplannedWhole) UnmarshalText(text []byte) error {
	w.N = len(text)
	return nil
}

// wholeSelfPlain unmarshals itself and embeds a pointer to its own type, one
// that is not allocated as the type is not exported. It holds no default.
type wholeSelfPlain struct {
	*wholeSelfPlain
	N int `schema:"n"`
}

func (w *wholeSelfPlain) UnmarshalText(text []byte) error {
	w.N = len(text)
	return nil
}

// Only the defaults plan descends into a struct decoded as a whole. Where
// there is none, a build never comes back through one, so the embedded
// pointer before it is expanded, and promotes its fields, as in the releases.
func TestEmbeddedPointerBeforeWholeStructWithoutDefaultsPlan(t *testing.T) {
	t.Parallel()

	w := unplannedWhole{unplannedInner: &unplannedInner{}}
	if err := NewDecoder().Decode(&w, map[string][]string{"x": {"1"}, "n": {"2"}, "b": {"abc"}}); err != nil {
		t.Fatal(err)
	}
	if w.N != 2 || w.X != 1 || w.B.N != 3 {
		t.Errorf("got n=%d x=%d b.n=%d, want 2, 1 and 3", w.N, w.X, w.B.N)
	}
	var in unplannedInner
	if err := NewDecoder().Decode(&in, map[string][]string{"x": {"1"}, "b": {"ab"}}); err != nil || in.X != 1 || in.B.N != 2 {
		t.Errorf("got %v, %+v, want nil, x=1 and b.n=2", err, in)
	}
	// Nor does a build come back through an embedded pointer that is itself
	// decoded as a whole.
	var self wholeSelfPlain
	if err := NewDecoder().Decode(&self, map[string][]string{"n": {"2"}}); err != nil || self.N != 2 || self.wholeSelfPlain != nil {
		t.Errorf("got %v, n=%d, embedded=%v, want nil, 2 and nil", err, self.N, self.wholeSelfPlain)
	}
}

// defaultedWholeInner holds a defaultedWhole value under a default option,
// which no struct takes. defaultedWhole unmarshals itself and embeds a pointer
// back to defaultedWholeInner.
type defaultedWholeInner struct {
	B defaultedWhole `schema:"b,default:x"`
	X int            `schema:"x,default:1"`
}

type defaultedWhole struct {
	*defaultedWholeInner
	N int `schema:"n"`
}

func (w *defaultedWhole) UnmarshalText(text []byte) error {
	w.N = len(text)
	return nil
}

// The defaults plan leaves a named struct that carries a default option alone
// as well, so one decoded as a whole is no way back even where there is a
// plan, and the embedded pointer before it is expanded as in the releases.
func TestEmbeddedPointerBeforeDefaultedWholeStruct(t *testing.T) {
	t.Parallel()

	w := defaultedWhole{defaultedWholeInner: &defaultedWholeInner{}}
	if err := NewDecoder().Decode(&w, map[string][]string{"x": {"2"}, "n": {"3"}, "b": {"abcd"}}); err != nil {
		t.Fatal(err)
	}
	if w.X != 2 || w.N != 3 || w.B.N != 4 {
		t.Errorf("got x=%d n=%d b.n=%d, want 2, 3 and 4", w.X, w.N, w.B.N)
	}
}

// The embedded pointers that lead back are found from the type alone.
func TestLeadsBack(t *testing.T) {
	t.Parallel()

	type hasSelf struct{ *hasSelf }
	type plain struct{ V int }
	type embedsPlain struct{ *plain }
	c := newCache()
	for _, tc := range []struct {
		t    reflect.Type
		want bool
	}{
		{reflect.TypeFor[SelfEmbedded](), true},
		{reflect.TypeFor[hasSelf](), true},
		{reflect.TypeFor[MutualA](), true},
		{reflect.TypeFor[MutualB](), true},
		{reflect.TypeFor[BackOuter](), true},
		{reflect.TypeFor[ViaValueInner](), true},
		{reflect.TypeFor[WholeSelf](), true},
		{reflect.TypeFor[WholeRingB](), true},
		{reflect.TypeFor[unplannedWhole](), false},
		{reflect.TypeFor[wholeSelfPlain](), false},
		{reflect.TypeFor[defaultedWhole](), false},
		{reflect.TypeFor[embedsPlain](), false},
	} {
		f := c.createField(tc.t.Field(0), "", "schema")
		planned := c.needsDefaultsWalk(tc.t, "schema", map[reflect.Type]bool{})
		if got := c.leadsBack(f, tc.t, "schema", planned); got != tc.want {
			t.Errorf("%v: leadsBack = %v, want %v", tc.t, got, tc.want)
		}
	}
}
