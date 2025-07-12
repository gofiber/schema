package schema

import (
	"reflect"
	"testing"
)

func TestConverters(t *testing.T) {
	tests := []struct {
		name  string
		v     reflect.Value
		want  interface{}
		valid bool
	}{
		{"boolTrue", convertBool("true"), true, true},
		{"boolOn", convertBool("on"), true, true},
		{"boolInvalid", convertBool("x"), nil, false},
		{"float32", convertFloat32("1.5"), float32(1.5), true},
		{"float32Invalid", convertFloat32("x"), nil, false},
		{"float64", convertFloat64("2.5"), 2.5, true},
		{"float64Invalid", convertFloat64("x"), nil, false},
		{"int", convertInt("10"), int(10), true},
		{"intInvalid", convertInt("x"), nil, false},
		{"uint", convertUint("5"), uint(5), true},
		{"uintInvalid", convertUint("-1"), nil, false},
		{"string", convertString("abc"), "abc", true},
	}
	for _, tt := range tests {
		if tt.valid {
			if !tt.v.IsValid() {
				t.Errorf("%s: expected valid value", tt.name)
				continue
			}
			if got := tt.v.Interface(); got != tt.want {
				t.Errorf("%s: expected %v, got %v", tt.name, tt.want, got)
			}
		} else if tt.v.IsValid() {
			t.Errorf("%s: expected invalid value", tt.name)
		}
	}
}

func TestBuiltinConverters(t *testing.T) {
	kinds := []reflect.Kind{
		reflect.Bool, reflect.Float32, reflect.Float64,
		reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64,
		reflect.String,
		reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64,
	}
	for _, k := range kinds {
		if builtinConverters[k] == nil {
			t.Errorf("missing converter for %v", k)
		}
	}
}

func BenchmarkConvertBool(b *testing.B) {
	for b.Loop() {
		convertBool("true")
	}
}

func BenchmarkConvertInt(b *testing.B) {
	for b.Loop() {
		convertInt("42")
	}
}
