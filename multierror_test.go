package schema

import (
	"errors"
	"strings"
	"testing"
)

func TestMultiErrorError(t *testing.T) {
	var m MultiError
	if got := m.Error(); got != "(0 errors)" {
		t.Fatalf("expected (0 errors), got %q", got)
	}

	errA := errors.New("a")
	m = MultiError{"a": errA}
	if got := m.Error(); got != errA.Error() {
		t.Fatalf("expected %q, got %q", errA.Error(), got)
	}

	errB := errors.New("b")
	m = MultiError{"a": errA, "b": errB}
	out := m.Error()
	if !strings.HasSuffix(out, "(and 1 other error)") {
		t.Fatalf("unexpected output %q", out)
	}
	if !strings.HasPrefix(out, errA.Error()) && !strings.HasPrefix(out, errB.Error()) {
		t.Fatalf("unexpected prefix %q", out)
	}

	errC := errors.New("c")
	m = MultiError{"a": errA, "b": errB, "c": errC}
	out = m.Error()
	if !strings.HasSuffix(out, "(and 2 other errors)") {
		t.Fatalf("unexpected output %q", out)
	}
}

func TestMultiErrorMerge(t *testing.T) {
	errA := errors.New("a")
	m1 := MultiError{"a": errA}
	errB := errors.New("b")
	m2 := MultiError{"a": errors.New("ignore"), "b": errB}
	m1.merge(m2)
	if len(m1) != 2 {
		t.Fatalf("expected len 2, got %d", len(m1))
	}
	if m1["a"] != errA {
		t.Errorf("existing key overwritten")
	}
	if m1["b"].Error() != errB.Error() {
		t.Errorf("missing merged error")
	}
}

func BenchmarkMultiErrorError(b *testing.B) {
	m := MultiError{"a": errors.New("a"), "b": errors.New("b"), "c": errors.New("c")}
	for b.Loop() {
		_ = m.Error()
	}
}
