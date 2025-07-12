package schema

import (
	"reflect"
	"testing"
)

func TestParseTagAndOptions(t *testing.T) {
	alias, opts := parseTag("name,omitempty,default:foo")
	if alias != "name" {
		t.Fatalf("expected alias name, got %s", alias)
	}
	if !opts.Contains("omitempty") {
		t.Fatalf("expected omitempty option")
	}
	if val := opts.getDefaultOptionValue(); val != "foo" {
		t.Fatalf("expected default foo, got %s", val)
	}
}

func TestFieldAlias(t *testing.T) {
	type S struct {
		Field string `json:"custom,omitempty"`
	}
	f, ok := reflect.TypeOf(S{}).FieldByName("Field")
	if !ok {
		t.Fatal("field not found")
	}
	alias, opts := fieldAlias(f, "json")
	if alias != "custom" {
		t.Fatalf("expected alias custom, got %s", alias)
	}
	if !opts.Contains("omitempty") {
		t.Fatalf("expected omitempty option")
	}
}

func TestTagOptionsContains(t *testing.T) {
	opts := tagOptions{"a", "b", "default:val"}
	if !opts.Contains("a") || opts.Contains("c") {
		t.Fatalf("contains failed")
	}
	if val := opts.getDefaultOptionValue(); val != "val" {
		t.Fatalf("expected default val, got %s", val)
	}
}

func TestIsValidStructPointer(t *testing.T) {
	type S struct{}
	if !isValidStructPointer(reflect.ValueOf(&S{})) {
		t.Errorf("expected true for struct pointer")
	}
	if isValidStructPointer(reflect.ValueOf(S{})) {
		t.Errorf("expected false for struct value")
	}
	var sp *S
	if isValidStructPointer(reflect.ValueOf(sp)) {
		t.Errorf("expected false for nil pointer")
	}
	var i int
	if isValidStructPointer(reflect.ValueOf(&i)) {
		t.Errorf("expected false for pointer to non-struct")
	}
}

func TestConvertPointer(t *testing.T) {
	v := convertPointer(reflect.Bool, "true")
	if !v.IsValid() || !v.Elem().Bool() {
		t.Fatalf("expected true, got %v", v)
	}

	v = convertPointer(reflect.Int, "10")
	if !v.IsValid() || v.Elem().Int() != 10 {
		t.Fatalf("expected 10, got %v", v)
	}

	v = convertPointer(reflect.String, "abc")
	if !v.IsValid() || v.Elem().String() != "abc" {
		t.Fatalf("expected abc, got %v", v)
	}

	v = convertPointer(reflect.Complex64, "1")
	if v.IsValid() {
		t.Fatalf("expected invalid value for unsupported kind")
	}
}

func BenchmarkParseTag(b *testing.B) {
	for i := 0; i < b.N; i++ {
		parseTag("field,omitempty,default:value")
	}
}

func BenchmarkIsZero(b *testing.B) {
	type S struct{ A int }
	v := reflect.ValueOf(S{})
	for i := 0; i < b.N; i++ {
		isZero(v)
	}
}

func BenchmarkConvertPointer(b *testing.B) {
	for i := 0; i < b.N; i++ {
		convertPointer(reflect.Int, "42")
	}
}

type customZero struct{ A int }

func (c customZero) IsZero() bool { return c.A == 0 }

func TestIsZeroCases(t *testing.T) {
	var sl []int
	if !isZero(reflect.ValueOf(sl)) {
		t.Errorf("nil slice should be zero")
	}
	sl = []int{}
	if !isZero(reflect.ValueOf(sl)) {
		t.Errorf("empty slice should be zero")
	}
	sl = []int{1}
	if isZero(reflect.ValueOf(sl)) {
		t.Errorf("non-empty slice considered zero")
	}

	arr := [2]int{}
	if !isZero(reflect.ValueOf(arr)) {
		t.Errorf("zero array should be zero")
	}
	arr = [2]int{0, 1}
	if isZero(reflect.ValueOf(arr)) {
		t.Errorf("non-zero array considered zero")
	}

	type S struct {
		A int
		B string
	}
	if !isZero(reflect.ValueOf(S{})) {
		t.Errorf("zero struct should be zero")
	}
	if isZero(reflect.ValueOf(S{A: 1})) {
		t.Errorf("non-zero struct considered zero")
	}

	if !isZero(reflect.ValueOf(customZero{})) {
		t.Errorf("IsZero method not used for zero value")
	}
	if isZero(reflect.ValueOf(customZero{A: 1})) {
		t.Errorf("IsZero method not used for non-zero value")
	}
}

func TestIsZeroFuncAndMap(t *testing.T) {
	tests := map[string]func(){
		"nil":     nil,
		"non-nil": func() {},
	}
	for name, fn := range tests {
		t.Run(name, func(t *testing.T) {
			defer func() {
				if r := recover(); r == nil {
					t.Errorf("expected panic for %s func", name)
				}
			}()
			isZero(reflect.ValueOf(fn))
		})
	}

	var m map[string]int
	if !isZero(reflect.ValueOf(m)) {
		t.Errorf("nil map should be zero")
	}
	m = map[string]int{}
	if !isZero(reflect.ValueOf(m)) {
		t.Errorf("empty map should be zero")
	}
	m["a"] = 1
	if isZero(reflect.ValueOf(m)) {
		t.Errorf("non-empty map considered zero")
	}
}
