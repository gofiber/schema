package schema

import (
	"errors"
	"fmt"
	"reflect"
	"testing"
)

type cycleNode struct {
	Next *cycleNode `schema:"next"`
	V    int        `schema:"v,default:1"`
}

// A destination holding a cycle is walked for defaults once around it.
// Walking it again recursed until the stack ran out, a fatal error that kills
// the process, and would read each struct again under a longer prefix, under
// which the key that set it is missing.
func TestDefaultsWalkStopsAtCycles(t *testing.T) {
	t.Parallel()

	self := &cycleNode{}
	self.Next = self
	if err := NewDecoder().Decode(self, map[string][]string{"v": {"0"}}); err != nil {
		t.Fatal(err)
	}
	if self.V != 0 {
		t.Errorf("self loop: v=%d, want the 0 sent", self.V)
	}

	// A ring longer than the structs the walk keeps inline, back to the destination.
	ring := make([]*cycleNode, len(defaultsWalk{}.walked)+2)
	for i := range ring {
		ring[i] = &cycleNode{}
	}
	for i, n := range ring {
		n.Next = ring[(i+1)%len(ring)]
	}
	if err := NewDecoder().Decode(ring[0], map[string][]string{"next.v": {"0"}}); err != nil {
		t.Fatal(err)
	}
	for i, n := range ring {
		if want := map[bool]int{true: 0, false: 1}[i == 1]; n.V != want {
			t.Errorf("ring[%d]: v=%d, want %d", i, n.V, want)
		}
	}

	// A cycle that does not pass through the destination.
	type holder struct {
		Head *cycleNode `schema:"head"`
		V    int        `schema:"v,default:2"`
	}
	a, b := &cycleNode{}, &cycleNode{}
	a.Next, b.Next = b, a
	h := holder{Head: a}
	if err := NewDecoder().Decode(&h, map[string][]string{}); err != nil {
		t.Fatal(err)
	}
	if h.V != 2 || a.V != 1 || b.V != 1 {
		t.Errorf("got v=%d, a.v=%d, b.v=%d, want 2, 1 and 1", h.V, a.V, b.V)
	}

	// A cycle of two at the end of a chain longer than the structs kept inline.
	chain := make([]*cycleNode, len(defaultsWalk{}.walked)+2)
	for i := range chain {
		chain[i] = &cycleNode{}
	}
	for i, n := range chain[:len(chain)-1] {
		n.Next = chain[i+1]
	}
	chain[len(chain)-1].Next = chain[len(chain)-2]
	if err := NewDecoder().Decode(chain[0], map[string][]string{}); err != nil {
		t.Fatal(err)
	}
	for i, n := range chain {
		if n.V != 1 {
			t.Errorf("chain[%d]: v=%d, want 1", i, n.V)
		}
	}
}

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

type walkNoDefaults struct {
	X int `schema:"x"`
}

// A set struct pointer whose struct holds no default and no embedded pointer
// to allocate is not walked, so it is not recorded either: recording each
// made a destination with a handful of them allocate a map on every Decode.
func TestDefaultsWalkSkipsStructsWithNothingToDo(t *testing.T) {
	t.Parallel()

	type S struct {
		A, B, C, D, E *walkNoDefaults
		N             *cycleNode `schema:"n"`
		V             int        `schema:"v,default:1"`
	}
	d := NewDecoder()
	set := S{A: &walkNoDefaults{}, B: &walkNoDefaults{}, C: &walkNoDefaults{}, D: &walkNoDefaults{}, E: &walkNoDefaults{}, N: &cycleNode{}}
	v := reflect.ValueOf(&set).Elem()
	info := d.cache.get(v.Type())
	w := defaultsWalk{root: info, n: 1}
	w.walked[0] = structAt(v.Addr().UnsafePointer(), v.Type())
	if errs := d.setDefaults(info, v, &w, &source{m: map[string][]string{}}, "", "", 0); errs != nil {
		t.Fatal(errs)
	}
	// The destination, and the one struct that holds a default.
	if w.n != 2 || w.more.n != 0 || set.V != 1 || set.N.V != 1 {
		t.Errorf("recorded %d structs and %d more, v=%d, n.v=%d; want 2, none, 1 and 1", w.n, w.more.n, set.V, set.N.V)
	}

	var s S
	src := map[string][]string{"a.x": {"1"}, "b.x": {"1"}, "c.x": {"1"}, "d.x": {"1"}, "e.x": {"1"}, "n.next.v": {"0"}}
	if err := d.Decode(&s, src); err != nil {
		t.Fatal(err)
	}
	if s.V != 1 || s.E == nil || s.E.X != 1 || s.N == nil || s.N.V != 1 || s.N.Next == nil || s.N.Next.V != 0 {
		t.Errorf("got %+v, n %+v", s, s.N)
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

type onceLeaf struct {
	V int `schema:"v,default:1"`
}

// onceHolder holds a leaf by value.
type onceHolder struct {
	Leaf onceLeaf `schema:"leaf"`
	N    int      `schema:"n,default:2"`
}

// onceFinder points to a leaf that another struct holds.
type onceFinder struct {
	P *onceLeaf `schema:"p"`
}

// OnceBase is held through an embedded pointer, and points back to itself.
type OnceBase struct {
	Back *OnceBase `schema:"back"`
	W    int       `schema:"w,default:1"`
}

type onceSelf struct {
	Self *onceSelf `schema:"self"`
	V    int       `schema:"v,default:1"`
}

// OnceShared is embedded by pointer at two levels of one struct.
type OnceShared struct {
	W int `schema:"w,default:1"`
}

type onceSharedSub struct {
	*OnceShared
	X int `schema:"x"`
}

type onceOuter struct {
	In onceHolder `schema:"in"`
	M  int        `schema:"m,default:3"`
}

// The defaults walk comes to a struct once, also where it is held in two
// places: under a pointer and by value, under a pointer and behind an embedded
// one, or behind two embedded ones. The walk that comes first sees the keys
// that name the struct's fields under its own prefix. A second one would look
// under another prefix, find none, and replace a zero that a key set by the
// default.
func TestDefaultsWalkReachesEachStructOnce(t *testing.T) {
	t.Parallel()

	// Behind an embedded pointer, with a pointer back to it.
	type child struct {
		*OnceBase
		V int `schema:"v"`
	}
	base := &OnceBase{}
	base.Back = base
	if err := NewDecoder().Decode(&child{OnceBase: base}, map[string][]string{"w": {"0"}}); err != nil || base.W != 0 {
		t.Errorf("embedded pointer: got %v, w=%d, want nil and the 0 sent", err, base.W)
	}

	// Held by value, with a pointer to it from inside it.
	type selfRoot struct {
		In onceSelf `schema:"in"`
	}
	var self selfRoot
	self.In.Self = &self.In
	if err := NewDecoder().Decode(&self, map[string][]string{"in.v": {"0"}}); err != nil || self.In.V != 0 {
		t.Errorf("struct value: got %v, v=%d, want nil and the 0 sent", err, self.In.V)
	}

	// Behind two embedded pointers of one struct, which no pointer step
	// leads to.
	type sharedReq struct {
		*OnceShared
		Sub onceSharedSub `schema:"sub"`
	}
	shared := &OnceShared{}
	req := sharedReq{OnceShared: shared, Sub: onceSharedSub{OnceShared: shared}}
	if err := NewDecoder().Decode(&req, map[string][]string{"w": {"0"}}); err != nil || shared.W != 0 {
		t.Errorf("two embedded pointers: got %v, w=%d, want nil and the 0 sent", err, shared.W)
	}

	// Under two pointers.
	type both struct {
		A *onceLeaf `schema:"a"`
		B *onceLeaf `schema:"b"`
	}
	leaf := &onceLeaf{}
	if err := NewDecoder().Decode(&both{A: leaf, B: leaf}, map[string][]string{"a.v": {"0"}}); err != nil || leaf.V != 0 {
		t.Errorf("two pointers: got %v, v=%d, want nil and the 0 sent", err, leaf.V)
	}

	// A pointer followed before the walk of the struct that holds its target:
	// the struct is the pointer's then, and its holder leaves it alone.
	type later struct {
		First  *onceFinder `schema:"first"`
		Second *onceHolder `schema:"second"`
	}
	holder := &onceHolder{}
	if err := NewDecoder().Decode(&later{First: &onceFinder{P: &holder.Leaf}, Second: holder}, map[string][]string{"first.p.v": {"0"}}); err != nil {
		t.Fatal(err)
	}
	if holder.Leaf.V != 0 || holder.N != 2 {
		t.Errorf("pointer first: got v=%d n=%d, want the 0 sent and the default 2", holder.Leaf.V, holder.N)
	}

	// A pointer followed from the walk of the struct that holds its target:
	// the struct is its holder's from the start, though its fields come
	// after the pointer.
	type owner struct {
		Far  *onceFinder `schema:"far"`
		Leaf onceLeaf    `schema:"leaf"`
	}
	var own owner
	own.Far = &onceFinder{P: &own.Leaf}
	if err := NewDecoder().Decode(&own, map[string][]string{"leaf.v": {"0"}}); err != nil || own.Leaf.V != 0 {
		t.Errorf("holder first: got %v, v=%d, want nil and the 0 sent", err, own.Leaf.V)
	}

	// A struct another walk came to is left alone with all it holds, also a
	// struct no pointer can lead to, which is not recorded for itself.
	type two struct {
		B *onceHolder `schema:"b"`
		A *onceOuter  `schema:"a"`
	}
	outer := &onceOuter{}
	if err := NewDecoder().Decode(&two{B: &outer.In, A: outer}, map[string][]string{"b.leaf.v": {"0"}, "b.n": {"0"}}); err != nil {
		t.Fatal(err)
	}
	if outer.In.Leaf.V != 0 || outer.In.N != 0 || outer.M != 3 {
		t.Errorf("with what it holds: got v=%d n=%d m=%d, want the zeros sent and the default 3", outer.In.Leaf.V, outer.In.N, outer.M)
	}
}

type onceFirst struct {
	V int `schema:"v,default:1"`
}

// onceFirstHolder holds a onceFirst as its first field, so at its own address,
// and a pointer to one, which makes it a struct the walk records.
type onceFirstHolder struct {
	First onceFirst        `schema:"first"`
	Other *onceFirst       `schema:"other"`
	Next  *onceFirstHolder `schema:"next"`
}

// A struct and its first field lie at one address and are two structs: the
// walk that has come to the one has not come to the other. The chain is longer
// than the structs the walk keeps inline, where they are compared one way, so
// that the last ones are compared the other way.
func TestDefaultsWalkTellsStructFromItsFirstField(t *testing.T) {
	t.Parallel()

	var head *onceFirstHolder
	for range len(defaultsWalk{}.walked) + 4 {
		head = &onceFirstHolder{Next: head}
	}
	if err := NewDecoder().Decode(head, map[string][]string{}); err != nil {
		t.Fatal(err)
	}
	for at, depth := head, 0; at != nil; at, depth = at.Next, depth+1 {
		if at.First.V != 1 {
			t.Errorf("%d levels down: v=%d, want the default 1", depth, at.First.V)
		}
	}
}

// Nothing is recorded until the walk follows a pointer, and then only the
// structs a pointer can lead to: until then no struct can be reached twice,
// and the others never can.
func TestDefaultsWalkRecordsOnlyWhatCanBeReachedTwice(t *testing.T) {
	t.Parallel()

	type dest struct {
		In   onceHolder  `schema:"in"`
		Self onceSelf    `schema:"self"`
		P    *onceFinder `schema:"p"`
		Q    *onceSelf   `schema:"q"`
	}
	d := NewDecoder()
	walk := func(v *dest) *defaultsWalk {
		rv := reflect.ValueOf(v).Elem()
		info := d.cache.get(rv.Type())
		w := &defaultsWalk{root: info, n: 1}
		w.walked[0] = structAt(rv.Addr().UnsafePointer(), rv.Type())
		if errs := d.setDefaults(info, rv, w, &source{m: map[string][]string{}}, "", "", 0); errs != nil {
			t.Fatal(errs)
		}
		return w
	}
	// No pointer is set: the destination alone.
	if w := walk(&dest{}); w.tracking || w.n != 1 {
		t.Errorf("no pointer set: tracking=%v with %d structs recorded, want false and 1", w.tracking, w.n)
	}
	// A pointer is followed: the destination, the struct behind the pointer,
	// and of the two held by value the one a pointer can lead to. The leaf in
	// the other could be pointed to as well, and is recorded for itself.
	if w := walk(&dest{Q: &onceSelf{}}); !w.tracking || w.n != 4 {
		t.Errorf("pointer set: tracking=%v with %d structs recorded, want true and 4", w.tracking, w.n)
	}
}

type encodeNode struct {
	N    int         `schema:"n"`
	Next *encodeNode `schema:"next"`
}

// encodeChain links n nodes numbered from 0, the last one to the node at
// index back, or to none if back is negative.
func encodeChain(n, back int) []*encodeNode {
	nodes := make([]*encodeNode, n)
	for i := range nodes {
		nodes[i] = &encodeNode{N: i}
	}
	for i, node := range nodes[:n-1] {
		node.Next = nodes[i+1]
	}
	if back >= 0 {
		nodes[n-1].Next = nodes[back]
	}
	return nodes
}

type encodeInner struct {
	Up *encodeOuter `schema:"up"`
	I  int          `schema:"i"`
}

type encodeOuter struct {
	In encodeInner `schema:"in"`
	O  int         `schema:"o"`
}

type encodeHub struct {
	A, B, C *encodeHub
	N       int `schema:"n"`
}

type encodeFirst struct {
	N int `schema:"n"`
}

// encodeHolder holds an encodeFirst as its first field, so at its own address.
type encodeHolder struct {
	First encodeFirst   `schema:"first"`
	Self  *encodeFirst  `schema:"self"`
	Down  *encodeHolder `schema:"down"`
}

// A value in which a pointer leads back into a struct that holds it made
// Encode follow the pointer without end, until the stack ran out: a fatal
// error that kills the process. Encode ends with an error instead, once it
// has gone round.
func TestEncodeEndsOnPointerCycles(t *testing.T) {
	t.Parallel()

	const want = "schema: encountered a cycle via *schema.encodeNode"
	check := func(name string, src any, round int) {
		t.Helper()
		dst := map[string][]string{}
		err := NewEncoder().Encode(src, dst)
		var cyc *cycleError
		if !errors.As(err, &cyc) || err.Error() != want {
			t.Errorf("%s: got error %v, want %s", name, err, want)
		}
		// The round is found before the way down is three times as deep as
		// it is long, counted from where the trail starts, and a short one
		// where it closes.
		most := untrailed + 1 + 3*round
		if round <= len(encodeTrail{}.last) {
			most = untrailed + 1 + round
		}
		if got := len(dst["n"]); got < round || got > most {
			t.Errorf("%s: encoded %d structs for a round of %d, want at most %d", name, got, round, most)
		}
	}
	for _, size := range []int{1, 2, 3, 4, 5, 7, 16, 100} {
		check(fmt.Sprintf("ring of %d", size), encodeChain(size, 0)[0], size)
	}
	// A round that the value given leads to but is not on, and one behind a
	// struct given by value, which is a copy no pointer leads back to.
	check("chain into a ring", encodeChain(12, 9)[0], 3)
	check("chain into a long ring", encodeChain(90, 30)[0], 60)
	check("by value", *encodeChain(3, 0)[0], 3)

	// A pointer back from a struct held by value.
	outer := &encodeOuter{O: 1}
	outer.In.Up = outer
	dst := map[string][]string{}
	err := NewEncoder().Encode(outer, dst)
	if cyc := new(*cycleError); !errors.As(err, cyc) || len(dst["o"]) > untrailed+4 {
		t.Errorf("through a struct value: got %v after %d structs", err, len(dst["o"]))
	}

	// Every struct holds several pointers back: the first round found ends
	// the call, where trying the other pointers too would take without end.
	hubs := make([]*encodeHub, 6)
	for i := range hubs {
		hubs[i] = &encodeHub{N: i}
	}
	for i, h := range hubs {
		h.A, h.B, h.C = hubs[(i+1)%len(hubs)], hubs[(i+2)%len(hubs)], hubs[(i+3)%len(hubs)]
	}
	dst = map[string][]string{}
	if err := NewEncoder().Encode(hubs[0], dst); !errors.As(err, new(*cycleError)) || len(dst["n"]) > untrailed+1+3*len(hubs) {
		t.Errorf("several pointers back: got %v after %d structs", err, len(dst["n"]))
	}
}

// A struct that is held twice but does not hold itself is encoded twice, as it
// always was, and a chain of pointers is encoded to its end, however long: on
// the trail as well, which only the structs encode is inside of are on.
func TestEncodeFollowsPointersThatDoNotLeadBack(t *testing.T) {
	t.Parallel()

	type pair struct {
		A, B *encodeNode
	}
	shared := &encodeNode{N: 5}
	dst := map[string][]string{}
	if err := NewEncoder().Encode(&pair{A: shared, B: shared}, dst); err != nil || !reflect.DeepEqual(dst["n"], []string{"5", "5"}) {
		t.Errorf("held twice: got %v, %v, want nil and two values", err, dst["n"])
	}

	// The same once the trail has started: a struct is off it again when
	// encode has left it.
	leaf := &encodeHub{N: -1}
	top := &encodeHub{A: leaf, B: leaf, C: leaf}
	for range untrailed + 2 {
		top = &encodeHub{A: top}
	}
	dst = map[string][]string{}
	err := NewEncoder().Encode(top, dst)
	held := 0
	for _, n := range dst["n"] {
		if n == "-1" {
			held++
		}
	}
	if err != nil || held != 3 {
		t.Errorf("held three times, deep: got %v and %d values, want nil and 3", err, held)
	}

	// A pointer to the first field of a struct on the trail has its address,
	// but leads to another struct, which does not hold the one on the trail.
	var holder *encodeHolder
	for i := range untrailed + 3 {
		holder = &encodeHolder{First: encodeFirst{N: i}, Down: holder}
		holder.Self = &holder.First
	}
	dst = map[string][]string{}
	if err := NewEncoder().Encode(holder, dst); err != nil || len(dst["n"]) != 2*(untrailed+3) {
		t.Errorf("pointer to a first field: got %v and %d values, want nil and %d", err, len(dst["n"]), 2*(untrailed+3))
	}

	long := encodeChain(5000, -1)
	enc := NewEncoder()
	dst = map[string][]string{}
	if err := enc.Encode(long[0], dst); err != nil || len(dst["n"]) != len(long) || dst["n"][len(long)-1] != "4999" {
		t.Errorf("long chain: got %v and %d values, want nil and %d", err, len(dst["n"]), len(long))
	}

	// A round found leaves the encoder as it was.
	if err := enc.Encode(encodeChain(2, 0)[0], map[string][]string{}); err == nil {
		t.Error("ring: got no error")
	}
	dst = map[string][]string{}
	if err := enc.Encode(&encodeNode{N: 7}, dst); err != nil || !reflect.DeepEqual(dst, map[string][]string{"n": {"7"}, "next": {"null"}}) {
		t.Errorf("after a ring: got %v, %v", err, dst)
	}
}
