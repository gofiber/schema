package schema

import (
	"fmt"
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

// setBuiltinKind must handle exactly the kinds builtinConverters covers, and
// with the same parse semantics; the two dispatch tables must not drift.
func TestSetBuiltinKindMatchesConverters(t *testing.T) {
	samples := map[reflect.Kind][]string{
		reflect.Bool:    {"true", "on", "0", "bogus"},
		reflect.Float32: {"1.5", "x"},
		reflect.Float64: {"2.25", "x"},
		reflect.Int:     {"42", "-7", "x", "9999999999999999999999"},
		reflect.Int8:    {"127", "128", "x"},
		reflect.Int16:   {"32767", "32768"},
		reflect.Int32:   {"1", "2147483648"},
		reflect.Int64:   {"12345678901", "x"},
		reflect.String:  {"anything", ""},
		reflect.Uint:    {"42", "-1", "x"},
		reflect.Uint8:   {"255", "256"},
		reflect.Uint16:  {"65535", "65536"},
		reflect.Uint32:  {"1", "4294967296"},
		reflect.Uint64:  {"12345678901", "-1"},
	}
	for k, conv := range builtinConverters {
		vals, ok := samples[k]
		if !ok {
			t.Fatalf("no samples for kind %v; extend this test alongside builtinConverters", k)
		}
		for _, s := range vals {
			target := reflect.New(typeForKind(k)).Elem()
			handled, setOK := setBuiltinKind(target, k, s)
			if !handled {
				t.Fatalf("setBuiltinKind does not handle kind %v present in builtinConverters", k)
			}
			convResult := conv(s)
			if setOK != convResult.IsValid() {
				t.Errorf("kind %v value %q: setBuiltinKind ok=%v but converter valid=%v", k, s, setOK, convResult.IsValid())
			}
			if setOK && convResult.IsValid() {
				if got, want := fmt.Sprint(target.Interface()), fmt.Sprint(convResult.Interface()); got != want {
					t.Errorf("kind %v value %q: setBuiltinKind stored %v, converter produced %v", k, s, got, want)
				}
			}
		}
	}
	for k := reflect.Invalid; k <= reflect.UnsafePointer; k++ {
		if builtinConverters[k] == nil {
			v := reflect.New(reflect.TypeOf(struct{}{})).Elem()
			if handled, _ := setBuiltinKind(v, k, "x"); handled {
				t.Errorf("setBuiltinKind handles kind %v that builtinConverters does not", k)
			}
		}
	}
}

func typeForKind(k reflect.Kind) reflect.Type {
	switch k {
	case reflect.Bool:
		return reflect.TypeOf(false)
	case reflect.Float32:
		return reflect.TypeOf(float32(0))
	case reflect.Float64:
		return reflect.TypeOf(float64(0))
	case reflect.Int:
		return reflect.TypeOf(int(0))
	case reflect.Int8:
		return reflect.TypeOf(int8(0))
	case reflect.Int16:
		return reflect.TypeOf(int16(0))
	case reflect.Int32:
		return reflect.TypeOf(int32(0))
	case reflect.Int64:
		return reflect.TypeOf(int64(0))
	case reflect.String:
		return reflect.TypeOf("")
	case reflect.Uint:
		return reflect.TypeOf(uint(0))
	case reflect.Uint8:
		return reflect.TypeOf(uint8(0))
	case reflect.Uint16:
		return reflect.TypeOf(uint16(0))
	case reflect.Uint32:
		return reflect.TypeOf(uint32(0))
	case reflect.Uint64:
		return reflect.TypeOf(uint64(0))
	}
	return reflect.TypeOf(struct{}{})
}

// decodeBuiltinSlice's sizing contract relies on no non-string builtin kind
// parsing a comma-containing value as a whole; pin that invariant.
func TestNoBuiltinKindParsesCommas(t *testing.T) {
	inputs := []string{"1,2", "true,false", "on,off", "1.5,2", ",", "1,"}
	for k, conv := range builtinConverters {
		if k == reflect.String {
			continue
		}
		for _, s := range inputs {
			if conv(s).IsValid() {
				t.Errorf("kind %v unexpectedly parses %q; decodeBuiltinSlice sizing depends on this failing", k, s)
			}
		}
	}
}
