package schema

import (
	"reflect"
	"testing"
)

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
