package schema

import (
	"fmt"
	"reflect"
	"strings"
	"testing"
	"time"
)

type E1 struct {
	F01 int     `schema:"f01"`
	F02 int     `schema:"-"`
	F03 string  `schema:"f03"`
	F04 string  `schema:"f04,omitempty"`
	F05 bool    `schema:"f05"`
	F06 bool    `schema:"f06"`
	F07 *string `schema:"f07"`
	F08 *int8   `schema:"f08"`
	F09 float64 `schema:"f09"`
	F10 func()  `schema:"f10"`
	F11 inner
}
type inner struct {
	F12 int
}

type SimpleStructForBenchmarkEncode struct {
	A string  `schema:"a"`
	B int     `schema:"b"`
	C bool    `schema:"c"`
	D float64 `schema:"d"`
	E struct {
		F float64 `schema:"f"`
	} `schema:"e"`
}

func TestFilled(t *testing.T) {
	f07 := "seven"
	var f08 int8 = 8
	s := &E1{
		F01: 1,
		F02: 2,
		F03: "three",
		F04: "four",
		F05: true,
		F06: false,
		F07: &f07,
		F08: &f08,
		F09: 1.618,
		F10: func() {},
		F11: inner{12},
	}

	vals := make(map[string][]string)
	errs := NewEncoder().Encode(s, vals)

	valExists(t, "f01", "1", vals)
	valNotExists(t, "f02", vals)
	valExists(t, "f03", "three", vals)
	valExists(t, "f05", "true", vals)
	valExists(t, "f06", "false", vals)
	valExists(t, "f07", "seven", vals)
	valExists(t, "f08", "8", vals)
	valExists(t, "f09", "1.618000", vals)
	valExists(t, "F12", "12", vals)

	emptyErr := MultiError{}
	if errs.Error() == emptyErr.Error() {
		t.Errorf("Expected error got %v", errs)
	}
}

type Aa int

type E3 struct {
	F01 bool    `schema:"f01"`
	F02 float32 `schema:"f02"`
	F03 float64 `schema:"f03"`
	F04 int     `schema:"f04"`
	F05 int8    `schema:"f05"`
	F06 int16   `schema:"f06"`
	F07 int32   `schema:"f07"`
	F08 int64   `schema:"f08"`
	F09 string  `schema:"f09"`
	F10 uint    `schema:"f10"`
	F11 uint8   `schema:"f11"`
	F12 uint16  `schema:"f12"`
	F13 uint32  `schema:"f13"`
	F14 uint64  `schema:"f14"`
	F15 Aa      `schema:"f15"`
}

// Test compatibility with default decoder types.
func TestCompat(t *testing.T) {
	src := &E3{
		F01: true,
		F02: 4.2,
		F03: 4.3,
		F04: -42,
		F05: -43,
		F06: -44,
		F07: -45,
		F08: -46,
		F09: "foo",
		F10: 42,
		F11: 43,
		F12: 44,
		F13: 45,
		F14: 46,
		F15: 1,
	}
	dst := &E3{}

	vals := make(map[string][]string)
	encoder := NewEncoder()
	decoder := NewDecoder()

	encoder.RegisterEncoder(src.F15, func(reflect.Value) string { return "1" })
	decoder.RegisterConverter(src.F15, func(string) reflect.Value { return reflect.ValueOf(1) })

	err := encoder.Encode(src, vals)
	if err != nil {
		t.Errorf("Encoder has non-nil error: %v", err)
	}
	err = decoder.Decode(dst, vals)
	if err != nil {
		t.Errorf("Decoder has non-nil error: %v", err)
	}

	if *src != *dst {
		t.Errorf("Decoder-Encoder compatibility: expected %v, got %v\n", src, dst)
	}
}

func TestEmpty(t *testing.T) {
	s := &E1{
		F01: 1,
		F02: 2,
		F03: "three",
	}

	estr := "schema: encoder not found for <nil>"
	vals := make(map[string][]string)
	err := NewEncoder().Encode(s, vals)
	if err.Error() != estr {
		t.Errorf("Expected: %s, got %v", estr, err)
	}

	valExists(t, "f03", "three", vals)
	valNotExists(t, "f04", vals)
}

func TestStruct(t *testing.T) {
	estr := "schema: interface must be a struct"
	vals := make(map[string][]string)
	err := NewEncoder().Encode("hello world", vals)

	if err.Error() != estr {
		t.Errorf("Expected: %s, got %v", estr, err)
	}
}

func TestSlices(t *testing.T) {
	type oneAsWord int
	ones := []oneAsWord{1, 2}
	s1 := &struct {
		ones     []oneAsWord `schema:"ones"`
		ints     []int       `schema:"ints"`
		nonempty []int       `schema:"nonempty"`
		empty    []int       `schema:"empty,omitempty"`
	}{ones, []int{1, 1}, []int{}, []int{}}
	vals := make(map[string][]string)

	encoder := NewEncoder()
	encoder.RegisterEncoder(ones[0], func(v reflect.Value) string { return "one" })
	err := encoder.Encode(s1, vals)
	if err != nil {
		t.Errorf("Encoder has non-nil error: %v", err)
	}

	valsExist(t, "ones", []string{"one", "one"}, vals)
	valsExist(t, "ints", []string{"1", "1"}, vals)
	valsExist(t, "nonempty", []string{}, vals)
	valNotExists(t, "empty", vals)
}

func TestCompatSlices(t *testing.T) {
	type oneAsWord int
	type s1 struct {
		Ones []oneAsWord `schema:"ones"`
		Ints []int       `schema:"ints"`
	}
	ones := []oneAsWord{1, 1}
	src := &s1{ones, []int{1, 1}}
	vals := make(map[string][]string)
	dst := &s1{}

	encoder := NewEncoder()
	encoder.RegisterEncoder(ones[0], func(v reflect.Value) string { return "one" })

	decoder := NewDecoder()
	decoder.RegisterConverter(ones[0], func(s string) reflect.Value {
		if s == "one" {
			return reflect.ValueOf(1)
		}
		return reflect.ValueOf(2)
	})

	err := encoder.Encode(src, vals)
	if err != nil {
		t.Errorf("Encoder has non-nil error: %v", err)
	}
	err = decoder.Decode(dst, vals)
	if err != nil {
		t.Errorf("Dncoder has non-nil error: %v", err)
	}

	if len(src.Ints) != len(dst.Ints) || len(src.Ones) != len(dst.Ones) {
		t.Fatalf("Expected %v, got %v", src, dst)
	}

	for i, v := range src.Ones {
		if dst.Ones[i] != v {
			t.Fatalf("Expected %v, got %v", v, dst.Ones[i])
		}
	}

	for i, v := range src.Ints {
		if dst.Ints[i] != v {
			t.Fatalf("Expected %v, got %v", v, dst.Ints[i])
		}
	}
}

func TestRegisterEncoder(t *testing.T) {
	type oneAsWord int
	type twoAsWord int
	type oneSliceAsWord []int

	s1 := &struct {
		oneAsWord
		twoAsWord
		oneSliceAsWord
	}{1, 2, []int{1, 1}}
	v1 := make(map[string][]string)

	encoder := NewEncoder()
	encoder.RegisterEncoder(s1.oneAsWord, func(v reflect.Value) string { return "one" })
	encoder.RegisterEncoder(s1.twoAsWord, func(v reflect.Value) string { return "two" })
	encoder.RegisterEncoder(s1.oneSliceAsWord, func(v reflect.Value) string { return "one" })

	err := encoder.Encode(s1, v1)
	if err != nil {
		t.Errorf("Encoder has non-nil error: %v", err)
	}

	valExists(t, "oneAsWord", "one", v1)
	valExists(t, "twoAsWord", "two", v1)
	valExists(t, "oneSliceAsWord", "one", v1)
}

func TestEncoderOrder(t *testing.T) {
	type builtinEncoderSimple int
	type builtinEncoderSimpleOverridden int
	type builtinEncoderSlice []int
	type builtinEncoderSliceOverridden []int
	type builtinEncoderStruct struct{ nr int }
	type builtinEncoderStructOverridden struct{ nr int }

	s1 := &struct {
		builtinEncoderSimple           `schema:"simple"`
		builtinEncoderSimpleOverridden `schema:"simple_overridden"`
		builtinEncoderSlice            `schema:"slice"`
		builtinEncoderSliceOverridden  `schema:"slice_overridden"`
		builtinEncoderStruct           `schema:"struct"`
		builtinEncoderStructOverridden `schema:"struct_overridden"`
	}{
		1,
		1,
		[]int{2},
		[]int{2},
		builtinEncoderStruct{3},
		builtinEncoderStructOverridden{3},
	}
	v1 := make(map[string][]string)

	encoder := NewEncoder()
	encoder.RegisterEncoder(s1.builtinEncoderSimpleOverridden, func(v reflect.Value) string { return "one" })
	encoder.RegisterEncoder(s1.builtinEncoderSliceOverridden, func(v reflect.Value) string { return "two" })
	encoder.RegisterEncoder(s1.builtinEncoderStructOverridden, func(v reflect.Value) string { return "three" })

	err := encoder.Encode(s1, v1)
	if err != nil {
		t.Errorf("Encoder has non-nil error: %v", err)
	}

	valExists(t, "simple", "1", v1)
	valExists(t, "simple_overridden", "one", v1)
	valExists(t, "slice", "2", v1)
	valExists(t, "slice_overridden", "two", v1)
	valExists(t, "nr", "3", v1)
	valExists(t, "struct_overridden", "three", v1)
}

func valExists(t *testing.T, key string, expect string, result map[string][]string) {
	valsExist(t, key, []string{expect}, result)
}

func valsExist(t *testing.T, key string, expect []string, result map[string][]string) {
	vals, ok := result[key]
	if !ok {
		t.Fatalf("Key not found. Expected: %s", key)
	}

	if len(expect) != len(vals) {
		t.Fatalf("Expected: %v, got: %v", expect, vals)
	}

	for i, v := range expect {
		if vals[i] != v {
			t.Fatalf("Unexpected value. Expected: %v, got %v", v, vals[i])
		}
	}
}

func valNotExists(t *testing.T, key string, result map[string][]string) {
	if val, ok := result[key]; ok {
		t.Error("Key not omitted. Expected: empty; got: " + val[0] + ".")
	}
}

func valsLength(t *testing.T, expectedLength int, result map[string][]string) {
	length := len(result)
	if length != expectedLength {
		t.Errorf("Expected length of %v, but got %v", expectedLength, length)
	}
}

func noError(t *testing.T, err error) {
	if err != nil {
		t.Errorf("Unexpected error. Got %v", err)
	}
}

type E4 struct {
	ID string `json:"id"`
}

func TestEncoderSetAliasTag(t *testing.T) {
	data := map[string][]string{}

	s := E4{
		ID: "foo",
	}
	encoder := NewEncoder()
	encoder.SetAliasTag("json")
	err := encoder.Encode(&s, data)
	if err != nil {
		t.Fatalf("Failed to encode: %v", err)
	}
	valExists(t, "id", "foo", data)
}

type E5 struct {
	F01 int      `schema:"f01,omitempty"`
	F02 string   `schema:"f02,omitempty"`
	F03 *string  `schema:"f03,omitempty"`
	F04 *int8    `schema:"f04,omitempty"`
	F05 float64  `schema:"f05,omitempty"`
	F06 E5F06    `schema:"f06,omitempty"`
	F07 E5F06    `schema:"f07,omitempty"`
	F08 []string `schema:"f08,omitempty"`
	F09 []string `schema:"f09,omitempty"`
}

type E5F06 struct {
	F0601 string `schema:"f0601,omitempty"`
}

func TestEncoderWithOmitempty(t *testing.T) {
	vals := map[string][]string{}

	s := E5{
		F02: "test",
		F07: E5F06{
			F0601: "test",
		},
		F09: []string{"test"},
	}

	encoder := NewEncoder()
	err := encoder.Encode(&s, vals)
	if err != nil {
		t.Fatalf("Failed to encode: %v", err)
	}

	valNotExists(t, "f01", vals)
	valExists(t, "f02", "test", vals)
	valNotExists(t, "f03", vals)
	valNotExists(t, "f04", vals)
	valNotExists(t, "f05", vals)
	valNotExists(t, "f06", vals)
	valExists(t, "f0601", "test", vals)
	valNotExists(t, "f08", vals)
	valsExist(t, "f09", []string{"test"}, vals)
}

type E6 struct {
	F01 *inner
	F02 *inner
	F03 *inner `schema:",omitempty"`
}

func TestStructPointer(t *testing.T) {
	vals := map[string][]string{}
	s := E6{
		F01: &inner{2},
	}

	encoder := NewEncoder()
	err := encoder.Encode(&s, vals)
	if err != nil {
		t.Fatalf("Failed to encode: %v", err)
	}
	valExists(t, "F12", "2", vals)
	valExists(t, "F02", "null", vals)
	valNotExists(t, "F03", vals)
}

func TestRegisterEncoderCustomArrayType(t *testing.T) {
	type CustomInt []int
	type S1 struct {
		SomeInts CustomInt `schema:",omitempty"`
	}

	ss := []S1{
		{},
		{CustomInt{}},
		{CustomInt{1, 2, 3}},
	}

	for s := range ss {
		vals := map[string][]string{}

		encoder := NewEncoder()
		encoder.RegisterEncoder(CustomInt{}, func(value reflect.Value) string {
			return fmt.Sprint(value.Interface())
		})

		err := encoder.Encode(ss[s], vals)
		if err != nil {
			t.Fatalf("Failed to encode: %v", err)
		}
	}
}

func TestRegisterEncoderStructIsZero(t *testing.T) {
	type S1 struct {
		SomeTime1 time.Time `schema:"tim1,omitempty"`
		SomeTime2 time.Time `schema:"tim2,omitempty"`
	}

	ss := []*S1{
		{
			SomeTime1: time.Date(2020, 8, 4, 13, 30, 1, 0, time.UTC),
		},
	}

	for s := range ss {
		vals := map[string][]string{}

		encoder := NewEncoder()
		encoder.RegisterEncoder(time.Time{}, func(value reflect.Value) string {
			tv, _ := reflect.TypeAssert[time.Time](value)
			return tv.Format(time.RFC3339Nano)
		})

		err := encoder.Encode(ss[s], vals)
		if err != nil {
			t.Errorf("Encoder has non-nil error: %v", err)
		}

		ta, ok := vals["tim1"]
		if !ok {
			t.Error("expected tim1 to be present")
		}

		if len(ta) != 1 {
			t.Error("expected tim1 to be present")
		}

		if ta[0] != "2020-08-04T13:30:01Z" {
			t.Error("expected correct tim1 time")
		}

		_, ok = vals["tim2"]
		if ok {
			t.Error("expected tim1 not to be present")
		}
	}
}

func TestRegisterEncoderWithPtrType(t *testing.T) {
	type CustomTime struct {
		time time.Time
	}

	type S1 struct {
		DateStart *CustomTime
		DateEnd   *CustomTime
		Empty     *CustomTime `schema:"empty,omitempty"`
	}

	ss := S1{
		DateStart: &CustomTime{time: time.Now()},
		DateEnd:   nil,
	}

	encoder := NewEncoder()
	encoder.RegisterEncoder(&CustomTime{}, func(value reflect.Value) string {
		if value.IsNil() {
			return ""
		}

		custom, _ := reflect.TypeAssert[*CustomTime](value)
		return custom.time.String()
	})

	vals := map[string][]string{}
	err := encoder.Encode(ss, vals)

	noError(t, err)
	valsLength(t, 2, vals)
	valExists(t, "DateStart", ss.DateStart.time.String(), vals)
	valExists(t, "DateEnd", "", vals)
}

func TestTimeDurationEncoding(t *testing.T) {
	type DurationStruct struct {
		Timeout time.Duration `schema:"timeout"`
	}

	vals := map[string][]string{}
	testData := DurationStruct{
		Timeout: 3 * time.Minute,
	}

	enc := NewEncoder()
	enc.RegisterEncoder(time.Duration(0), func(v reflect.Value) string {
		d, _ := reflect.TypeAssert[time.Duration](v)
		return d.String() // "3m0s"
	})

	err := enc.Encode(&testData, vals)
	if err != nil {
		t.Fatalf("Failed to encode time.Duration: %v", err)
	}

	got, ok := vals["timeout"]
	if !ok || len(got) < 1 {
		t.Fatalf("Encoded map missing key 'timeout'")
	}
	if got[0] != (3 * time.Minute).String() {
		t.Errorf("Expected %q, got %q", (3 * time.Minute).String(), got[0])
	}
}

// Test for omitempty with zero time.Duration.
func TestTimeDurationOmitEmpty(t *testing.T) {
	type DurationStruct struct {
		Timeout time.Duration `schema:"timeout,omitempty"`
	}

	vals := map[string][]string{}
	testData := DurationStruct{
		Timeout: 0,
	}

	enc := NewEncoder()
	enc.RegisterEncoder(time.Duration(0), func(v reflect.Value) string {
		d, _ := reflect.TypeAssert[time.Duration](v)
		return d.String()
	})

	err := enc.Encode(&testData, vals)
	if err != nil {
		t.Fatalf("Failed to encode time.Duration: %v", err)
	}
	// Should be omitted since 0 for time.Duration is "zero" and tagged as omitempty
	if _, found := vals["timeout"]; found {
		t.Errorf("Expected 'timeout' to be omitted, but it was present: %v", vals["timeout"])
	}
}

func TestEncoderZeroAndNonZeroFields(t *testing.T) {
	type ZeroTestStruct struct {
		A string  `schema:"a,omitempty"`
		B int     `schema:"b,omitempty"`
		C float64 `schema:"c,omitempty"`
		D bool    `schema:"d,omitempty"`
		E *int    `schema:"e,omitempty"`
		F *string `schema:"f,omitempty"`
		G string  `schema:"g"` // no omitempty
	}

	vals := map[string][]string{}
	intVal := 42
	strVal := "Hello"
	s := ZeroTestStruct{
		A: "",
		B: 0,
		C: 0.0,
		D: false,
		E: &intVal,
		F: &strVal,
		G: "MustEncode",
	}

	enc := NewEncoder()
	err := enc.Encode(&s, vals)
	if err != nil {
		t.Fatalf("Encoding error: %v", err)
	}

	// Fields A, B, C, D are zero and should be omitted
	if _, found := vals["a"]; found {
		t.Errorf("Expected 'a' to be omitted for zero string")
	}
	if _, found := vals["b"]; found {
		t.Errorf("Expected 'b' to be omitted for zero int")
	}
	if _, found := vals["c"]; found {
		t.Errorf("Expected 'c' to be omitted for zero float")
	}
	if _, found := vals["d"]; found {
		t.Errorf("Expected 'd' to be omitted for false bool")
	}

	// E is a pointer to an int, so it should appear
	gotE, found := vals["e"]
	if !found {
		t.Error("Expected 'e' to be present")
	} else if len(gotE) != 1 || gotE[0] != "42" {
		t.Errorf("Expected '42', got %v", gotE)
	}

	// F is a pointer to string, so it should appear
	gotF, found := vals["f"]
	if !found {
		t.Error("Expected 'f' to be present")
	} else if len(gotF) != 1 || gotF[0] != "Hello" {
		t.Errorf("Expected 'Hello', got %v", gotF)
	}

	// G has no omitempty tag and must be encoded
	gotG, found := vals["g"]
	if !found {
		t.Error("Expected 'g' to be present")
	} else if len(gotG) != 1 || gotG[0] != "MustEncode" {
		t.Errorf("Expected 'MustEncode', got %v", gotG)
	}
}

func BenchmarkSimpleStructEncode(b *testing.B) {
	s := SimpleStructForBenchmarkEncode{
		A: "abc",
		B: 123,
		C: true,
		D: 3.14,
		E: struct {
			F float64 `schema:"f"`
		}{F: 6.28},
	}
	enc := NewEncoder()

	vals := map[string][]string{}
	for b.Loop() {
		_ = enc.Encode(&s, vals)
	}
}

func BenchmarkSimpleStructEncodeParallel(b *testing.B) {
	s := SimpleStructForBenchmarkEncode{
		A: "abc",
		B: 123,
		C: true,
		D: 3.14,
		E: struct {
			F float64 `schema:"f"`
		}{F: 6.28},
	}
	enc := NewEncoder()

	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		vals := map[string][]string{}
		for pb.Next() {
			_ = enc.Encode(&s, vals)
		}
	})
}

type LargeStructForBenchmarkEncode struct {
	F1 string   `schema:"f1"`
	F2 string   `schema:"f2"`
	F3 int      `schema:"f3"`
	F4 int      `schema:"f4"`
	F5 []string `schema:"f5"`
	F6 []int    `schema:"f6"`
	F7 float64  `schema:"f7"`
	F8 bool     `schema:"f8"`
	F9 struct {
		N1 time.Time `schema:"n1"`
		N2 string    `schema:"n2"`
	} `schema:"f9"`
}

func BenchmarkLargeStructEncode(b *testing.B) {
	s := LargeStructForBenchmarkEncode{
		F1: "Lorem", F2: "Ipsum", F3: 123, F4: 456,
		F5: []string{"A", "B", "C", "D"},
		F6: []int{10, 20, 30, 40},
		F7: 3.14159, F8: true,
		F9: struct {
			N1 time.Time `schema:"n1"`
			N2 string    `schema:"n2"`
		}{
			N1: time.Date(2025, 1, 1, 12, 0, 0, 0, time.UTC),
			N2: "NestedStringValue",
		},
	}
	enc := NewEncoder()

	// Optionally register a custom encoder for time.Time
	enc.RegisterEncoder(time.Time{}, func(v reflect.Value) string {
		tVal, _ := reflect.TypeAssert[time.Time](v)
		return tVal.Format(time.RFC3339)
	})

	vals := map[string][]string{}
	for b.Loop() {
		_ = enc.Encode(&s, vals)
	}
}

func BenchmarkLargeStructEncodeParallel(b *testing.B) {
	s := LargeStructForBenchmarkEncode{
		F1: "Lorem", F2: "Ipsum", F3: 123, F4: 456,
		F5: []string{"A", "B", "C", "D"},
		F6: []int{10, 20, 30, 40},
		F7: 3.14159, F8: true,
		F9: struct {
			N1 time.Time `schema:"n1"`
			N2 string    `schema:"n2"`
		}{
			N1: time.Date(2025, 1, 1, 12, 0, 0, 0, time.UTC),
			N2: "NestedStringValue",
		},
	}
	enc := NewEncoder()
	enc.RegisterEncoder(time.Time{}, func(v reflect.Value) string {
		tVal, _ := reflect.TypeAssert[time.Time](v)
		return tVal.Format(time.RFC3339)
	})

	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		vals := map[string][]string{}
		for pb.Next() {
			_ = enc.Encode(&s, vals)
		}
	})
}

func BenchmarkTimeDurationEncoding(b *testing.B) {
	type DurationStruct struct {
		Timeout time.Duration `schema:"timeout"`
	}

	testData := DurationStruct{
		Timeout: 5 * time.Second,
	}

	enc := NewEncoder()
	enc.RegisterEncoder(time.Duration(0), func(v reflect.Value) string {
		d, _ := reflect.TypeAssert[time.Duration](v)
		return d.String()
	})

	vals := map[string][]string{}
	for b.Loop() {
		_ = enc.Encode(&testData, vals)
	}
}

// A plan build racing SetAliasTag/RegisterEncoder must not re-insert a
// stale plan after the invalidation; once the reconfiguration call returns,
// every subsequent Encode must observe it.
func TestEncoderReconfigureDuringEncode(t *testing.T) {
	type T1 struct {
		A string `schema:"schemaname" url:"urlname"`
	}
	for i := 0; i < 300; i++ {
		e := NewEncoder()
		done := make(chan struct{})
		go func() {
			defer close(done)
			dst := map[string][]string{}
			_ = e.Encode(T1{A: "x"}, dst)
		}()
		e.SetAliasTag("url")
		<-done

		dst := map[string][]string{}
		if err := e.Encode(T1{A: "x"}, dst); err != nil {
			t.Fatal(err)
		}
		if _, ok := dst["urlname"]; !ok {
			t.Fatalf("iteration %d: encode after SetAliasTag returned stale plan: %v", i, dst)
		}
	}
}

// Errors inside recursed structs (value and pointer) must be collected under
// the nested type's name.
func TestEncoderNestedErrors(t *testing.T) {
	type Bad struct {
		M map[string]string `schema:"m"`
	}
	type S struct {
		V Bad  `schema:"v"`
		P *Bad `schema:"p"`
	}
	enc := NewEncoder()
	dst := map[string][]string{}
	err := enc.Encode(S{V: Bad{M: map[string]string{"a": "b"}}, P: &Bad{}}, dst)
	if err == nil {
		t.Fatal("expected error for unsupported nested field")
	}
	if !strings.Contains(err.Error(), "encoder not found") {
		t.Errorf("unexpected error: %v", err)
	}
}

// An encoding plan stored by a build that raced a reconfiguration carries
// the old generation and must never be served after the reconfiguration
// returned, even if it lands in the cache.
func TestEncoderStaleCacheEntryIgnored(t *testing.T) {
	type T1 struct {
		A string `schema:"schemaname" url:"urlname"`
	}
	e := NewEncoder()
	typ := reflect.TypeOf(T1{})

	dst := map[string][]string{}
	if err := e.Encode(T1{A: "x"}, dst); err != nil {
		t.Fatal(err)
	}
	stalePlan, ok := e.encCache.Load(typ)
	if !ok {
		t.Fatal("expected cached plan")
	}

	e.SetAliasTag("url")
	// Simulate a delayed store from a build that raced the reconfiguration.
	e.encCache.Store(typ, stalePlan)

	dst = map[string][]string{}
	if err := e.Encode(T1{A: "y"}, dst); err != nil {
		t.Fatal(err)
	}
	if _, ok := dst["urlname"]; !ok {
		t.Fatalf("stale plan served: %v", dst)
	}
}

// Pins the four cases the removed isValidStructPointer helper used to
// classify, now decided by the cached plan's recurseStructPtr flag plus a
// runtime nil check: valid struct pointers recurse, nil struct pointers
// encode as "null", struct values recurse, and pointers to non-structs
// encode their element (or "null" when nil).
func TestEncodeStructPointerClassification(t *testing.T) {
	type Inner struct {
		X string `schema:"x"`
	}
	type S struct {
		SP  *Inner `schema:"sp"`  // valid struct pointer -> recursed
		NP  *Inner `schema:"np"`  // nil struct pointer -> "null"
		SV  Inner  `schema:"sv"`  // struct value -> recursed
		IP  *int   `schema:"ip"`  // pointer to non-struct -> element
		NIP *int   `schema:"nip"` // nil pointer to non-struct -> "null"
	}

	seven := 7
	src := S{
		SP: &Inner{X: "from-ptr"},
		SV: Inner{X: "from-val"},
		IP: &seven,
	}
	// SP and SV both recurse and write under the inner field's alias; SV is
	// encoded after SP overwrote nothing (values append under "x").
	dst := map[string][]string{}
	if err := NewEncoder().Encode(src, dst); err != nil {
		t.Fatal(err)
	}

	if got := dst["x"]; len(got) != 2 || got[0] != "from-ptr" || got[1] != "from-val" {
		t.Errorf("struct pointer/value recursion: expected [from-ptr from-val] under \"x\", got %v", got)
	}
	if _, ok := dst["sp"]; ok {
		t.Error("valid struct pointer must recurse, not encode under its own alias")
	}
	if got := dst["np"]; len(got) != 1 || got[0] != "null" {
		t.Errorf("nil struct pointer: expected [null], got %v", got)
	}
	if got := dst["ip"]; len(got) != 1 || got[0] != "7" {
		t.Errorf("pointer to non-struct: expected [7], got %v", got)
	}
	if got := dst["nip"]; len(got) != 1 || got[0] != "null" {
		t.Errorf("nil pointer to non-struct: expected [null], got %v", got)
	}
}

// Encoding into a nil destination map used to panic and crash the caller.
func TestEncodeNilDstReturnsError(t *testing.T) {
	type S struct {
		A string `schema:"a"`
	}
	err := NewEncoder().Encode(S{A: "x"}, nil)
	if err == nil || !strings.Contains(err.Error(), "dst map must not be nil") {
		t.Fatalf("expected nil-dst error, got %v", err)
	}
}

// Non-nil pointers to unsupported element types used to invoke a nil inner
// encoder and panic, crashing the caller; they must error instead, while
// nil pointers keep encoding as "null".
func TestEncodePointerToUnsupported(t *testing.T) {
	type S struct {
		M *map[string]string `schema:"m"`
		L *[]int             `schema:"l"`
		A string             `schema:"a"`
	}

	m := map[string]string{"k": "v"}
	l := []int{1}
	dst := map[string][]string{}
	err := NewEncoder().Encode(S{M: &m, L: &l, A: "x"}, dst)
	if err == nil || !strings.Contains(err.Error(), "encoder not found") {
		t.Fatalf("expected encoder-not-found error, got %v", err)
	}
	if got := dst["a"]; len(got) != 1 || got[0] != "x" {
		t.Errorf("supported sibling must still encode: %v", dst)
	}

	// Nil pointers to unsupported types keep the historical "null" output.
	dst = map[string][]string{}
	if err := NewEncoder().Encode(S{A: "x"}, dst); err != nil {
		t.Fatal(err)
	}
	if got := dst["m"]; len(got) != 1 || got[0] != "null" {
		t.Errorf("nil *map: expected [null], got %v", dst["m"])
	}
	if got := dst["l"]; len(got) != 1 || got[0] != "null" {
		t.Errorf("nil *slice: expected [null], got %v", dst["l"])
	}

	// omitempty suppresses the null.
	type SO struct {
		M *map[string]string `schema:"m,omitempty"`
	}
	dst = map[string][]string{}
	if err := NewEncoder().Encode(SO{}, dst); err != nil {
		t.Fatal(err)
	}
	if _, ok := dst["m"]; ok {
		t.Errorf("omitempty nil pointer must be skipped: %v", dst)
	}
}

// Panics from user-registered encoders (or reflection) must surface as
// errors, mirroring Decode's recovery, instead of crashing the caller.
func TestEncodeRecoversEncoderPanic(t *testing.T) {
	type PT struct{ V int }
	type S struct {
		P PT `schema:"p"`
	}
	enc := NewEncoder()
	enc.RegisterEncoder(PT{}, func(reflect.Value) string { panic("boom") })
	err := enc.Encode(S{P: PT{V: 1}}, map[string][]string{})
	if err == nil || !strings.Contains(err.Error(), "panic while encoding: boom") {
		t.Fatalf("expected recovered panic error, got %v", err)
	}
}

// An empty slice whose element type has no encoder (e.g. []*Struct) must be
// skipped under omitempty and emitted empty otherwise, not spuriously error;
// only a non-empty such slice errors. Non-slice unencodable fields still error.
func TestEncodeEmptySliceUnencodableElem(t *testing.T) {
	type Inner struct {
		A string `schema:"a"`
	}

	// omitempty + nil slice -> skipped, no error.
	type SOmit struct {
		Children []*Inner `schema:"children,omitempty"`
		Name     string   `schema:"name"`
	}
	dst := map[string][]string{}
	if err := NewEncoder().Encode(SOmit{Name: "bob"}, dst); err != nil {
		t.Fatalf("omitempty empty slice should not error: %v", err)
	}
	if _, ok := dst["children"]; ok {
		t.Errorf("omitempty empty slice should be skipped, got %v", dst["children"])
	}
	if got := dst["name"]; len(got) != 1 || got[0] != "bob" {
		t.Errorf("sibling not encoded: %v", dst)
	}

	// no omitempty + nil slice -> emitted empty, no error.
	type SPlain struct {
		Children []*Inner `schema:"children"`
		Name     string   `schema:"name"`
	}
	dst = map[string][]string{}
	if err := NewEncoder().Encode(SPlain{Name: "bob"}, dst); err != nil {
		t.Fatalf("empty slice should not error: %v", err)
	}
	if got, ok := dst["children"]; !ok || len(got) != 0 {
		t.Errorf("empty slice should emit empty, got %v (present=%v)", got, ok)
	}

	// non-empty slice of unencodable element -> error.
	dst = map[string][]string{}
	err := NewEncoder().Encode(SPlain{Children: []*Inner{{A: "x"}}, Name: "bob"}, dst)
	if err == nil || !strings.Contains(err.Error(), "encoder not found") {
		t.Errorf("non-empty unencodable-elem slice should error, got %v", err)
	}

	// non-slice unencodable field (map) still errors regardless of emptiness.
	type SMap struct {
		M map[string]string `schema:"m"`
	}
	if err := NewEncoder().Encode(SMap{}, map[string][]string{}); err == nil {
		t.Errorf("empty map field should still error 'encoder not found'")
	}
	if err := NewEncoder().Encode(SMap{M: map[string]string{"k": "v"}}, map[string][]string{}); err == nil {
		t.Errorf("non-empty map field should error 'encoder not found'")
	}
}
