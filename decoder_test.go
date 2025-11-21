// Copyright 2012 The Gorilla Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package schema

import (
	"encoding/hex"
	"errors"
	"fmt"
	"mime/multipart"
	"reflect"
	"strings"
	"testing"
	"time"
)

type IntAlias int

type rudeBool bool

func (id *rudeBool) UnmarshalText(text []byte) error {
	value := string(text)
	switch {
	case strings.EqualFold("Yup", value):
		*id = true
	case strings.EqualFold("Nope", value):
		*id = false
	default:
		return errors.New("value must be yup or nope")
	}
	return nil
}

// All cases we want to cover, in a nutshell.
type S1 struct {
	F01 int         `schema:"f1"`
	F02 *int        `schema:"f2"`
	F03 []int       `schema:"f3"`
	F04 []*int      `schema:"f4"`
	F05 *[]int      `schema:"f5"`
	F06 *[]*int     `schema:"f6"`
	F07 S2          `schema:"f7"`
	F08 *S1         `schema:"f8"`
	F09 int         `schema:"-"`
	F10 []S1        `schema:"f10"`
	F11 []*S1       `schema:"f11"`
	F12 *[]S1       `schema:"f12"`
	F13 *[]*S1      `schema:"f13"`
	F14 int         `schema:"f14"`
	F15 IntAlias    `schema:"f15"`
	F16 []IntAlias  `schema:"f16"`
	F17 S19         `schema:"f17"`
	F18 rudeBool    `schema:"f18"`
	F19 *rudeBool   `schema:"f19"`
	F20 []rudeBool  `schema:"f20"`
	F21 []*rudeBool `schema:"f21"`
}

type LargeStructForBenchmark struct {
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

// A simple struct for demonstration benchmarks
type SimpleStructForBenchmark struct {
	A string  `schema:"a"`
	B int     `schema:"b"`
	C bool    `schema:"c"`
	D float64 `schema:"d"`
	E struct {
		F float64 `schema:"f"`
	} `schema:"e"`
}

type S2 struct {
	F01 *[]*int `schema:"f1"`
}

type S19 [2]byte

func (id *S19) UnmarshalText(text []byte) error {
	buf, err := hex.DecodeString(string(text))
	if err != nil {
		return err
	}
	if len(buf) > len(*id) {
		return errors.New("out of range")
	}
	copy((*id)[:], buf)
	return nil
}

func TestAll(t *testing.T) {
	v := map[string][]string{
		"f1":             {"1"},
		"f2":             {"2"},
		"f3":             {"31", "32"},
		"f4":             {"41", "42"},
		"f5":             {"51", "52"},
		"f6":             {"61", "62"},
		"f7.f1":          {"71", "72"},
		"f8.f8.f7.f1":    {"81", "82"},
		"f9":             {"9"},
		"f10.0.f10.0.f6": {"101", "102"},
		"f10.0.f10.1.f6": {"103", "104"},
		"f11.0.f11.0.f6": {"111", "112"},
		"f11.0.f11.1.f6": {"113", "114"},
		"f12.0.f12.0.f6": {"121", "122"},
		"f12.0.f12.1.f6": {"123", "124"},
		"f13.0.f13.0.f6": {"131", "132"},
		"f13.0.f13.1.f6": {"133", "134"},
		"f14":            {},
		"f15":            {"151"},
		"f16":            {"161", "162"},
		"f17":            {"1a2b"},
		"f18":            {"yup"},
		"f19":            {"nope"},
		"f20":            {"nope", "yup"},
		"f21":            {"yup", "nope"},
	}
	f2 := 2
	f41, f42 := 41, 42
	f61, f62 := 61, 62
	f71, f72 := 71, 72
	f81, f82 := 81, 82
	f101, f102, f103, f104 := 101, 102, 103, 104
	f111, f112, f113, f114 := 111, 112, 113, 114
	f121, f122, f123, f124 := 121, 122, 123, 124
	f131, f132, f133, f134 := 131, 132, 133, 134
	var f151 IntAlias = 151
	var f161, f162 IntAlias = 161, 162
	var f152, f153 rudeBool = true, false
	e := S1{
		F01: 1,
		F02: &f2,
		F03: []int{31, 32},
		F04: []*int{&f41, &f42},
		F05: &[]int{51, 52},
		F06: &[]*int{&f61, &f62},
		F07: S2{
			F01: &[]*int{&f71, &f72},
		},
		F08: &S1{
			F08: &S1{
				F07: S2{
					F01: &[]*int{&f81, &f82},
				},
			},
		},
		F09: 0,
		F10: []S1{
			{
				F10: []S1{
					{F06: &[]*int{&f101, &f102}},
					{F06: &[]*int{&f103, &f104}},
				},
			},
		},
		F11: []*S1{
			{
				F11: []*S1{
					{F06: &[]*int{&f111, &f112}},
					{F06: &[]*int{&f113, &f114}},
				},
			},
		},
		F12: &[]S1{
			{
				F12: &[]S1{
					{F06: &[]*int{&f121, &f122}},
					{F06: &[]*int{&f123, &f124}},
				},
			},
		},
		F13: &[]*S1{
			{
				F13: &[]*S1{
					{F06: &[]*int{&f131, &f132}},
					{F06: &[]*int{&f133, &f134}},
				},
			},
		},
		F14: 0,
		F15: f151,
		F16: []IntAlias{f161, f162},
		F17: S19{0x1a, 0x2b},
		F18: f152,
		F19: &f153,
		F20: []rudeBool{f153, f152},
		F21: []*rudeBool{&f152, &f153},
	}

	s := &S1{}
	_ = NewDecoder().Decode(s, v)

	vals := func(values []*int) []int {
		r := make([]int, len(values))
		for k, v := range values {
			r[k] = *v
		}
		return r
	}

	if s.F01 != e.F01 {
		t.Errorf("f1: expected %v, got %v", e.F01, s.F01)
	}
	if s.F02 == nil {
		t.Errorf("f2: expected %v, got nil", *e.F02)
	} else if *s.F02 != *e.F02 {
		t.Errorf("f2: expected %v, got %v", *e.F02, *s.F02)
	}
	if s.F03 == nil {
		t.Errorf("f3: expected %v, got nil", e.F03)
	} else if len(s.F03) != 2 || s.F03[0] != e.F03[0] || s.F03[1] != e.F03[1] {
		t.Errorf("f3: expected %v, got %v", e.F03, s.F03)
	}
	if s.F04 == nil {
		t.Errorf("f4: expected %v, got nil", e.F04)
	} else {
		if len(s.F04) != 2 || *(s.F04)[0] != *(e.F04)[0] || *(s.F04)[1] != *(e.F04)[1] {
			t.Errorf("f4: expected %v, got %v", vals(e.F04), vals(s.F04))
		}
	}
	if s.F05 == nil {
		t.Errorf("f5: expected %v, got nil", e.F05)
	} else {
		sF05, eF05 := *s.F05, *e.F05
		if len(sF05) != 2 || sF05[0] != eF05[0] || sF05[1] != eF05[1] {
			t.Errorf("f5: expected %v, got %v", eF05, sF05)
		}
	}
	if s.F06 == nil {
		t.Errorf("f6: expected %v, got nil", vals(*e.F06))
	} else {
		sF06, eF06 := *s.F06, *e.F06
		if len(sF06) != 2 || *(sF06)[0] != *(eF06)[0] || *(sF06)[1] != *(eF06)[1] {
			t.Errorf("f6: expected %v, got %v", vals(eF06), vals(sF06))
		}
	}
	if s.F07.F01 == nil {
		t.Errorf("f7.f1: expected %v, got nil", vals(*e.F07.F01))
	} else {
		sF07, eF07 := *s.F07.F01, *e.F07.F01
		if len(sF07) != 2 || *(sF07)[0] != *(eF07)[0] || *(sF07)[1] != *(eF07)[1] {
			t.Errorf("f7.f1: expected %v, got %v", vals(eF07), vals(sF07))
		}
	}
	if s.F08 == nil {
		t.Errorf("f8: got nil")
	} else if s.F08.F08 == nil {
		t.Errorf("f8.f8: got nil")
	} else if s.F08.F08.F07.F01 == nil {
		t.Errorf("f8.f8.f7.f1: expected %v, got nil", vals(*e.F08.F08.F07.F01))
	} else {
		sF08, eF08 := *s.F08.F08.F07.F01, *e.F08.F08.F07.F01
		if len(sF08) != 2 || *(sF08)[0] != *(eF08)[0] || *(sF08)[1] != *(eF08)[1] {
			t.Errorf("f8.f8.f7.f1: expected %v, got %v", vals(eF08), vals(sF08))
		}
	}
	if s.F09 != e.F09 {
		t.Errorf("f9: expected %v, got %v", e.F09, s.F09)
	}
	if s.F10 == nil {
		t.Errorf("f10: got nil")
	} else if len(s.F10) != 1 {
		t.Errorf("f10: expected 1 element, got %v", s.F10)
	} else {
		if len(s.F10[0].F10) != 2 {
			t.Errorf("f10.0.f10: expected 1 element, got %v", s.F10[0].F10)
		} else {
			sF10, eF10 := *s.F10[0].F10[0].F06, *e.F10[0].F10[0].F06
			if sF10 == nil {
				t.Errorf("f10.0.f10.0.f6: expected %v, got nil", vals(eF10))
			} else {
				if len(sF10) != 2 || *(sF10)[0] != *(eF10)[0] || *(sF10)[1] != *(eF10)[1] {
					t.Errorf("f10.0.f10.0.f6: expected %v, got %v", vals(eF10), vals(sF10))
				}
			}
			sF10, eF10 = *s.F10[0].F10[1].F06, *e.F10[0].F10[1].F06
			if sF10 == nil {
				t.Errorf("f10.0.f10.0.f6: expected %v, got nil", vals(eF10))
			} else {
				if len(sF10) != 2 || *(sF10)[0] != *(eF10)[0] || *(sF10)[1] != *(eF10)[1] {
					t.Errorf("f10.0.f10.0.f6: expected %v, got %v", vals(eF10), vals(sF10))
				}
			}
		}
	}
	if s.F11 == nil {
		t.Errorf("f11: got nil")
	} else if len(s.F11) != 1 {
		t.Errorf("f11: expected 1 element, got %v", s.F11)
	} else {
		if len(s.F11[0].F11) != 2 {
			t.Errorf("f11.0.f11: expected 1 element, got %v", s.F11[0].F11)
		} else {
			sF11, eF11 := *s.F11[0].F11[0].F06, *e.F11[0].F11[0].F06
			if sF11 == nil {
				t.Errorf("f11.0.f11.0.f6: expected %v, got nil", vals(eF11))
			} else {
				if len(sF11) != 2 || *(sF11)[0] != *(eF11)[0] || *(sF11)[1] != *(eF11)[1] {
					t.Errorf("f11.0.f11.0.f6: expected %v, got %v", vals(eF11), vals(sF11))
				}
			}
			sF11, eF11 = *s.F11[0].F11[1].F06, *e.F11[0].F11[1].F06
			if sF11 == nil {
				t.Errorf("f11.0.f11.0.f6: expected %v, got nil", vals(eF11))
			} else {
				if len(sF11) != 2 || *(sF11)[0] != *(eF11)[0] || *(sF11)[1] != *(eF11)[1] {
					t.Errorf("f11.0.f11.0.f6: expected %v, got %v", vals(eF11), vals(sF11))
				}
			}
		}
	}
	if s.F12 == nil {
		t.Errorf("f12: got nil")
	} else if len(*s.F12) != 1 {
		t.Errorf("f12: expected 1 element, got %v", *s.F12)
	} else {
		sF12, eF12 := *(s.F12), *(e.F12)
		if len(*sF12[0].F12) != 2 {
			t.Errorf("f12.0.f12: expected 1 element, got %v", *sF12[0].F12)
		} else {
			sF122, eF122 := *(*sF12[0].F12)[0].F06, *(*eF12[0].F12)[0].F06
			if sF122 == nil {
				t.Errorf("f12.0.f12.0.f6: expected %v, got nil", vals(eF122))
			} else {
				if len(sF122) != 2 || *(sF122)[0] != *(eF122)[0] || *(sF122)[1] != *(eF122)[1] {
					t.Errorf("f12.0.f12.0.f6: expected %v, got %v", vals(eF122), vals(sF122))
				}
			}
			sF122, eF122 = *(*sF12[0].F12)[1].F06, *(*eF12[0].F12)[1].F06
			if sF122 == nil {
				t.Errorf("f12.0.f12.0.f6: expected %v, got nil", vals(eF122))
			} else {
				if len(sF122) != 2 || *(sF122)[0] != *(eF122)[0] || *(sF122)[1] != *(eF122)[1] {
					t.Errorf("f12.0.f12.0.f6: expected %v, got %v", vals(eF122), vals(sF122))
				}
			}
		}
	}
	if s.F13 == nil {
		t.Errorf("f13: got nil")
	} else if len(*s.F13) != 1 {
		t.Errorf("f13: expected 1 element, got %v", *s.F13)
	} else {
		sF13, eF13 := *(s.F13), *(e.F13)
		if len(*sF13[0].F13) != 2 {
			t.Errorf("f13.0.f13: expected 1 element, got %v", *sF13[0].F13)
		} else {
			sF132, eF132 := *(*sF13[0].F13)[0].F06, *(*eF13[0].F13)[0].F06
			if sF132 == nil {
				t.Errorf("f13.0.f13.0.f6: expected %v, got nil", vals(eF132))
			} else {
				if len(sF132) != 2 || *(sF132)[0] != *(eF132)[0] || *(sF132)[1] != *(eF132)[1] {
					t.Errorf("f13.0.f13.0.f6: expected %v, got %v", vals(eF132), vals(sF132))
				}
			}
			sF132, eF132 = *(*sF13[0].F13)[1].F06, *(*eF13[0].F13)[1].F06
			if sF132 == nil {
				t.Errorf("f13.0.f13.0.f6: expected %v, got nil", vals(eF132))
			} else {
				if len(sF132) != 2 || *(sF132)[0] != *(eF132)[0] || *(sF132)[1] != *(eF132)[1] {
					t.Errorf("f13.0.f13.0.f6: expected %v, got %v", vals(eF132), vals(sF132))
				}
			}
		}
	}
	if s.F14 != e.F14 {
		t.Errorf("f14: expected %v, got %v", e.F14, s.F14)
	}
	if s.F15 != e.F15 {
		t.Errorf("f15: expected %v, got %v", e.F15, s.F15)
	}
	if s.F16 == nil {
		t.Errorf("f16: nil")
	} else if len(s.F16) != len(e.F16) {
		t.Errorf("f16: expected len %d, got %d", len(e.F16), len(s.F16))
	} else if !reflect.DeepEqual(s.F16, e.F16) {
		t.Errorf("f16: expected %v, got %v", e.F16, s.F16)
	}
	if s.F17 != e.F17 {
		t.Errorf("f17: expected %v, got %v", e.F17, s.F17)
	}
	if s.F18 != e.F18 {
		t.Errorf("f18: expected %v, got %v", e.F18, s.F18)
	}
	if *s.F19 != *e.F19 {
		t.Errorf("f19: expected %v, got %v", *e.F19, *s.F19)
	}
	if s.F20 == nil {
		t.Errorf("f20: nil")
	} else if len(s.F20) != len(e.F20) {
		t.Errorf("f20: expected %v, got %v", e.F20, s.F20)
	} else if !reflect.DeepEqual(s.F20, e.F20) {
		t.Errorf("f20: expected %v, got %v", e.F20, s.F20)
	}
	if s.F21 == nil {
		t.Errorf("f21: nil")
	} else if len(s.F21) != len(e.F21) {
		t.Errorf("f21: expected length %d, got %d", len(e.F21), len(s.F21))
	} else if !reflect.DeepEqual(s.F21, e.F21) {
		t.Errorf("f21: expected %v, got %v", e.F21, s.F21)
	}
}

func BenchmarkAll(b *testing.B) {
	v := map[string][]string{
		"f1":             {"1"},
		"f2":             {"2"},
		"f3":             {"31", "32"},
		"f4":             {"41", "42"},
		"f5":             {"51", "52"},
		"f6":             {"61", "62"},
		"f7.f1":          {"71", "72"},
		"f8.f8.f7.f1":    {"81", "82"},
		"f9":             {"9"},
		"f10.0.f10.0.f6": {"101", "102"},
		"f10.0.f10.1.f6": {"103", "104"},
		"f11.0.f11.0.f6": {"111", "112"},
		"f11.0.f11.1.f6": {"113", "114"},
		"f12.0.f12.0.f6": {"121", "122"},
		"f12.0.f12.1.f6": {"123", "124"},
		"f13.0.f13.0.f6": {"131", "132"},
		"f13.0.f13.1.f6": {"133", "134"},
	}

	decoder := NewDecoder()
	for b.Loop() {
		_ = decoder.Decode(S1{}, v)
	}
}

// ----------------------------------------------------------------------------

type S3 struct {
	F01 bool
	F02 float32
	F03 float64
	F04 int
	F05 int8
	F06 int16
	F07 int32
	F08 int64
	F09 string
	F10 uint
	F11 uint8
	F12 uint16
	F13 uint32
	F14 uint64
}

func TestDefaultConverters(t *testing.T) {
	v := map[string][]string{
		"F01": {"true"},
		"F02": {"4.2"},
		"F03": {"4.3"},
		"F04": {"-42"},
		"F05": {"-43"},
		"F06": {"-44"},
		"F07": {"-45"},
		"F08": {"-46"},
		"F09": {"foo"},
		"F10": {"42"},
		"F11": {"43"},
		"F12": {"44"},
		"F13": {"45"},
		"F14": {"46"},
	}
	e := S3{
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
	}
	s := &S3{}
	_ = NewDecoder().Decode(s, v)
	if s.F01 != e.F01 {
		t.Errorf("F01: expected %v, got %v", e.F01, s.F01)
	}
	if s.F02 != e.F02 {
		t.Errorf("F02: expected %v, got %v", e.F02, s.F02)
	}
	if s.F03 != e.F03 {
		t.Errorf("F03: expected %v, got %v", e.F03, s.F03)
	}
	if s.F04 != e.F04 {
		t.Errorf("F04: expected %v, got %v", e.F04, s.F04)
	}
	if s.F05 != e.F05 {
		t.Errorf("F05: expected %v, got %v", e.F05, s.F05)
	}
	if s.F06 != e.F06 {
		t.Errorf("F06: expected %v, got %v", e.F06, s.F06)
	}
	if s.F07 != e.F07 {
		t.Errorf("F07: expected %v, got %v", e.F07, s.F07)
	}
	if s.F08 != e.F08 {
		t.Errorf("F08: expected %v, got %v", e.F08, s.F08)
	}
	if s.F09 != e.F09 {
		t.Errorf("F09: expected %v, got %v", e.F09, s.F09)
	}
	if s.F10 != e.F10 {
		t.Errorf("F10: expected %v, got %v", e.F10, s.F10)
	}
	if s.F11 != e.F11 {
		t.Errorf("F11: expected %v, got %v", e.F11, s.F11)
	}
	if s.F12 != e.F12 {
		t.Errorf("F12: expected %v, got %v", e.F12, s.F12)
	}
	if s.F13 != e.F13 {
		t.Errorf("F13: expected %v, got %v", e.F13, s.F13)
	}
	if s.F14 != e.F14 {
		t.Errorf("F14: expected %v, got %v", e.F14, s.F14)
	}
}

func TestOn(t *testing.T) {
	v := map[string][]string{
		"F01": {"on"},
	}
	s := S3{}
	err := NewDecoder().Decode(&s, v)
	if err != nil {
		t.Fatal(err)
	}
	if !s.F01 {
		t.Fatal("Value was not set to true")
	}
}

// ----------------------------------------------------------------------------

func TestInlineStruct(t *testing.T) {
	s1 := &struct {
		F01 bool
	}{}
	s2 := &struct {
		F01 int
	}{}
	v1 := map[string][]string{
		"F01": {"true"},
	}
	v2 := map[string][]string{
		"F01": {"42"},
	}
	decoder := NewDecoder()
	_ = decoder.Decode(s1, v1)
	if s1.F01 != true {
		t.Errorf("s1: expected %v, got %v", true, s1.F01)
	}
	_ = decoder.Decode(s2, v2)
	if s2.F01 != 42 {
		t.Errorf("s2: expected %v, got %v", 42, s2.F01)
	}
}

// ----------------------------------------------------------------------------

type Foo struct {
	F01 int
	F02 Bar
	Bif []Baz
}

type Bar struct {
	F01 string
	F02 string
	F03 string
	F14 string
	S05 string
	Str string
}

type Baz struct {
	F99 []string
}

func TestSimpleExample(t *testing.T) {
	data := map[string][]string{
		"F01":       {"1"},
		"F02.F01":   {"S1"},
		"F02.F02":   {"S2"},
		"F02.F03":   {"S3"},
		"F02.F14":   {"S4"},
		"F02.S05":   {"S5"},
		"F02.Str":   {"Str"},
		"Bif.0.F99": {"A", "B", "C"},
	}

	e := &Foo{
		F01: 1,
		F02: Bar{
			F01: "S1",
			F02: "S2",
			F03: "S3",
			F14: "S4",
			S05: "S5",
			Str: "Str",
		},
		Bif: []Baz{
			{
				F99: []string{"A", "B", "C"},
			},
		},
	}

	s := &Foo{}
	_ = NewDecoder().Decode(s, data)

	if s.F01 != e.F01 {
		t.Errorf("F01: expected %v, got %v", e.F01, s.F01)
	}
	if s.F02.F01 != e.F02.F01 {
		t.Errorf("F02.F01: expected %v, got %v", e.F02.F01, s.F02.F01)
	}
	if s.F02.F02 != e.F02.F02 {
		t.Errorf("F02.F02: expected %v, got %v", e.F02.F02, s.F02.F02)
	}
	if s.F02.F03 != e.F02.F03 {
		t.Errorf("F02.F03: expected %v, got %v", e.F02.F03, s.F02.F03)
	}
	if s.F02.F14 != e.F02.F14 {
		t.Errorf("F02.F14: expected %v, got %v", e.F02.F14, s.F02.F14)
	}
	if s.F02.S05 != e.F02.S05 {
		t.Errorf("F02.S05: expected %v, got %v", e.F02.S05, s.F02.S05)
	}
	if s.F02.Str != e.F02.Str {
		t.Errorf("F02.Str: expected %v, got %v", e.F02.Str, s.F02.Str)
	}
	if len(s.Bif) != len(e.Bif) {
		t.Errorf("Bif len: expected %d, got %d", len(e.Bif), len(s.Bif))
	} else {
		if len(s.Bif[0].F99) != len(e.Bif[0].F99) {
			t.Errorf("Bif[0].F99 len: expected %d, got %d", len(e.Bif[0].F99), len(s.Bif[0].F99))
		}
	}
}

// ----------------------------------------------------------------------------

type S4 struct {
	F01 int64
	F02 float64
	F03 bool
	F04 rudeBool
}

func TestConversionError(t *testing.T) {
	data := map[string][]string{
		"F01": {"foo"},
		"F02": {"bar"},
		"F03": {"baz"},
		"F04": {"not-a-yes-or-nope"},
	}
	s := &S4{}
	e := NewDecoder().Decode(s, data)

	m := e.(MultiError)
	if len(m) != 4 {
		t.Errorf("Expected 3 errors, got %v", m)
	}
}

// ----------------------------------------------------------------------------

type S5 struct {
	F01 []string
}

func TestEmptyValue(t *testing.T) {
	data := map[string][]string{
		"F01": {"", "foo"},
	}
	s := &S5{}
	err := NewDecoder().Decode(s, data)
	if err != nil {
		t.Fatalf("Failed to decode: %v", err)
	}
	if len(s.F01) != 1 {
		t.Errorf("Expected 1 values in F01")
	}
}

func TestEmptyValueZeroEmpty(t *testing.T) {
	data := map[string][]string{
		"F01": {"", "foo"},
	}
	s := S5{}
	d := NewDecoder()
	d.ZeroEmpty(true)
	err := d.Decode(&s, data)
	if err != nil {
		t.Fatal(err)
	}
	if len(s.F01) != 2 {
		t.Errorf("Expected 1 values in F01")
	}
}

// ----------------------------------------------------------------------------

type S6 struct {
	id string
}

func TestUnexportedField(t *testing.T) {
	data := map[string][]string{
		"id": {"identifier"},
	}
	s := &S6{}
	err := NewDecoder().Decode(s, data)
	if err != nil {
		t.Fatalf("Failed to decode: %v", err)
	}
	if s.id != "" {
		t.Errorf("Unexported field expected to be ignored")
	}
}

// ----------------------------------------------------------------------------

type S7 struct {
	ID string
}

func TestMultipleValues(t *testing.T) {
	data := map[string][]string{
		"ID": {"0", "1"},
	}

	s := S7{}
	err := NewDecoder().Decode(&s, data)
	if err != nil {
		t.Fatalf("Failed to decode: %v", err)
	}
	if s.ID != "1" {
		t.Errorf("Last defined value must be used when multiple values for same field are provided")
	}
}

type S8 struct {
	ID string `json:"id"`
}

func TestSetAliasTag(t *testing.T) {
	data := map[string][]string{
		"id": {"foo"},
	}

	s := S8{}
	dec := NewDecoder()
	dec.SetAliasTag("json")
	err := dec.Decode(&s, data)
	if err != nil {
		t.Fatalf("Failed to decode: %v", err)
	}
	if s.ID != "foo" {
		t.Fatalf("Bad value: got %q, want %q", s.ID, "foo")
	}
}

func TestZeroEmpty(t *testing.T) {
	data := map[string][]string{
		"F01": {""},
		"F03": {"true"},
	}
	s := S4{1, 1, false, false}
	d := NewDecoder()
	d.ZeroEmpty(true)

	err := d.Decode(&s, data)
	if err != nil {
		t.Fatal(err)
	}
	if s.F01 != 0 {
		t.Errorf("F01: got %v, want %v", s.F01, 0)
	}
	if s.F02 != 1 {
		t.Errorf("F02: got %v, want %v", s.F02, 1)
	}
	if s.F03 != true {
		t.Errorf("F03: got %v, want %v", s.F03, true)
	}
}

func TestNoZeroEmpty(t *testing.T) {
	data := map[string][]string{
		"F01": {""},
		"F03": {"true"},
	}
	s := S4{1, 1, false, false}
	d := NewDecoder()
	d.ZeroEmpty(false)
	err := d.Decode(&s, data)
	if err != nil {
		t.Fatal(err)
	}
	if s.F01 != 1 {
		t.Errorf("F01: got %v, want %v", s.F01, 1)
	}
	if s.F02 != 1 {
		t.Errorf("F02: got %v, want %v", s.F02, 1)
	}
	if s.F03 != true {
		t.Errorf("F03: got %v, want %v", s.F03, true)
	}
	if s.F04 != false {
		t.Errorf("F04: got %v, want %v", s.F04, false)
	}
}

// ----------------------------------------------------------------------------

type S9 struct {
	Id string
}

type S10 struct {
	S9
}

func TestEmbeddedField(t *testing.T) {
	data := map[string][]string{
		"Id": {"identifier"},
	}
	s := &S10{}
	err := NewDecoder().Decode(s, data)
	if err != nil {
		t.Fatalf("Failed to decode: %v", err)
	}
	if s.Id != "identifier" {
		t.Errorf("Missing support for embedded fields")
	}
}

type S11 struct {
	S10
}

func TestMultipleLevelEmbeddedField(t *testing.T) {
	data := map[string][]string{
		"Id": {"identifier"},
	}
	s := &S11{}
	err := NewDecoder().Decode(s, data)
	if s.Id != "identifier" {
		t.Errorf("Missing support for multiple-level embedded fields (%v)", err)
	}
}

func TestInvalidPath(t *testing.T) {
	data := map[string][]string{
		"Foo.Bar": {"baz"},
	}
	s := S9{}
	err := NewDecoder().Decode(&s, data)
	expectedErr := `schema: invalid path "Foo.Bar"`
	if err.Error() != expectedErr {
		t.Fatalf("got %q, want %q", err, expectedErr)
	}
}

func TestInvalidPathIgnoreUnknownKeys(t *testing.T) {
	data := map[string][]string{
		"Foo.Bar": {"baz"},
	}
	s := S9{}
	dec := NewDecoder()
	dec.IgnoreUnknownKeys(true)
	err := dec.Decode(&s, data)
	if err != nil {
		t.Fatal(err)
	}
}

// ----------------------------------------------------------------------------

type S1NT struct {
	F1  int
	F2  *int
	F3  []int
	F4  []*int
	F5  *[]int
	F6  *[]*int
	F7  S2
	F8  *S1
	F9  int `schema:"-"`
	F10 []S1
	F11 []*S1
	F12 *[]S1
	F13 *[]*S1
}

func TestAllNT(t *testing.T) {
	v := map[string][]string{
		"f1":             {"1"},
		"f2":             {"2"},
		"f3":             {"31", "32"},
		"f4":             {"41", "42"},
		"f5":             {"51", "52"},
		"f6":             {"61", "62"},
		"f7.f1":          {"71", "72"},
		"f8.f8.f7.f1":    {"81", "82"},
		"f9":             {"9"},
		"f10.0.f10.0.f6": {"101", "102"},
		"f10.0.f10.1.f6": {"103", "104"},
		"f11.0.f11.0.f6": {"111", "112"},
		"f11.0.f11.1.f6": {"113", "114"},
		"f12.0.f12.0.f6": {"121", "122"},
		"f12.0.f12.1.f6": {"123", "124"},
		"f13.0.f13.0.f6": {"131", "132"},
		"f13.0.f13.1.f6": {"133", "134"},
	}
	f2 := 2
	f41, f42 := 41, 42
	f61, f62 := 61, 62
	f71, f72 := 71, 72
	f81, f82 := 81, 82
	f101, f102, f103, f104 := 101, 102, 103, 104
	f111, f112, f113, f114 := 111, 112, 113, 114
	f121, f122, f123, f124 := 121, 122, 123, 124
	f131, f132, f133, f134 := 131, 132, 133, 134
	e := S1NT{
		F1: 1,
		F2: &f2,
		F3: []int{31, 32},
		F4: []*int{&f41, &f42},
		F5: &[]int{51, 52},
		F6: &[]*int{&f61, &f62},
		F7: S2{
			F01: &[]*int{&f71, &f72},
		},
		F8: &S1{
			F08: &S1{
				F07: S2{
					F01: &[]*int{&f81, &f82},
				},
			},
		},
		F9: 0,
		F10: []S1{
			{
				F10: []S1{
					{F06: &[]*int{&f101, &f102}},
					{F06: &[]*int{&f103, &f104}},
				},
			},
		},
		F11: []*S1{
			{
				F11: []*S1{
					{F06: &[]*int{&f111, &f112}},
					{F06: &[]*int{&f113, &f114}},
				},
			},
		},
		F12: &[]S1{
			{
				F12: &[]S1{
					{F06: &[]*int{&f121, &f122}},
					{F06: &[]*int{&f123, &f124}},
				},
			},
		},
		F13: &[]*S1{
			{
				F13: &[]*S1{
					{F06: &[]*int{&f131, &f132}},
					{F06: &[]*int{&f133, &f134}},
				},
			},
		},
	}

	s := &S1NT{}
	_ = NewDecoder().Decode(s, v)

	vals := func(values []*int) []int {
		r := make([]int, len(values))
		for k, v := range values {
			r[k] = *v
		}
		return r
	}

	if s.F1 != e.F1 {
		t.Errorf("f1: expected %v, got %v", e.F1, s.F1)
	}
	if s.F2 == nil {
		t.Errorf("f2: expected %v, got nil", *e.F2)
	} else if *s.F2 != *e.F2 {
		t.Errorf("f2: expected %v, got %v", *e.F2, *s.F2)
	}
	if s.F3 == nil {
		t.Errorf("f3: expected %v, got nil", e.F3)
	} else if len(s.F3) != 2 || s.F3[0] != e.F3[0] || s.F3[1] != e.F3[1] {
		t.Errorf("f3: expected %v, got %v", e.F3, s.F3)
	}
	if s.F4 == nil {
		t.Errorf("f4: expected %v, got nil", e.F4)
	} else {
		if len(s.F4) != 2 || *(s.F4)[0] != *(e.F4)[0] || *(s.F4)[1] != *(e.F4)[1] {
			t.Errorf("f4: expected %v, got %v", vals(e.F4), vals(s.F4))
		}
	}
	if s.F5 == nil {
		t.Errorf("f5: expected %v, got nil", e.F5)
	} else {
		sF5, eF5 := *s.F5, *e.F5
		if len(sF5) != 2 || sF5[0] != eF5[0] || sF5[1] != eF5[1] {
			t.Errorf("f5: expected %v, got %v", eF5, sF5)
		}
	}
	if s.F6 == nil {
		t.Errorf("f6: expected %v, got nil", vals(*e.F6))
	} else {
		sF6, eF6 := *s.F6, *e.F6
		if len(sF6) != 2 || *(sF6)[0] != *(eF6)[0] || *(sF6)[1] != *(eF6)[1] {
			t.Errorf("f6: expected %v, got %v", vals(eF6), vals(sF6))
		}
	}
	if s.F7.F01 == nil {
		t.Errorf("f7.f1: expected %v, got nil", vals(*e.F7.F01))
	} else {
		sF7, eF7 := *s.F7.F01, *e.F7.F01
		if len(sF7) != 2 || *(sF7)[0] != *(eF7)[0] || *(sF7)[1] != *(eF7)[1] {
			t.Errorf("f7.f1: expected %v, got %v", vals(eF7), vals(sF7))
		}
	}
	if s.F8 == nil {
		t.Errorf("f8: got nil")
	} else if s.F8.F08 == nil {
		t.Errorf("f8.f8: got nil")
	} else if s.F8.F08.F07.F01 == nil {
		t.Errorf("f8.f8.f7.f1: expected %v, got nil", vals(*e.F8.F08.F07.F01))
	} else {
		sF8, eF8 := *s.F8.F08.F07.F01, *e.F8.F08.F07.F01
		if len(sF8) != 2 || *(sF8)[0] != *(eF8)[0] || *(sF8)[1] != *(eF8)[1] {
			t.Errorf("f8.f8.f7.f1: expected %v, got %v", vals(eF8), vals(sF8))
		}
	}
	if s.F9 != e.F9 {
		t.Errorf("f9: expected %v, got %v", e.F9, s.F9)
	}
	if s.F10 == nil {
		t.Errorf("f10: got nil")
	} else if len(s.F10) != 1 {
		t.Errorf("f10: expected 1 element, got %v", s.F10)
	} else {
		if len(s.F10[0].F10) != 2 {
			t.Errorf("f10.0.f10: expected 1 element, got %v", s.F10[0].F10)
		} else {
			sF10, eF10 := *s.F10[0].F10[0].F06, *e.F10[0].F10[0].F06
			if sF10 == nil {
				t.Errorf("f10.0.f10.0.f6: expected %v, got nil", vals(eF10))
			} else {
				if len(sF10) != 2 || *(sF10)[0] != *(eF10)[0] || *(sF10)[1] != *(eF10)[1] {
					t.Errorf("f10.0.f10.0.f6: expected %v, got %v", vals(eF10), vals(sF10))
				}
			}
			sF10, eF10 = *s.F10[0].F10[1].F06, *e.F10[0].F10[1].F06
			if sF10 == nil {
				t.Errorf("f10.0.f10.0.f6: expected %v, got nil", vals(eF10))
			} else {
				if len(sF10) != 2 || *(sF10)[0] != *(eF10)[0] || *(sF10)[1] != *(eF10)[1] {
					t.Errorf("f10.0.f10.0.f6: expected %v, got %v", vals(eF10), vals(sF10))
				}
			}
		}
	}
	if s.F11 == nil {
		t.Errorf("f11: got nil")
	} else if len(s.F11) != 1 {
		t.Errorf("f11: expected 1 element, got %v", s.F11)
	} else {
		if len(s.F11[0].F11) != 2 {
			t.Errorf("f11.0.f11: expected 1 element, got %v", s.F11[0].F11)
		} else {
			sF11, eF11 := *s.F11[0].F11[0].F06, *e.F11[0].F11[0].F06
			if sF11 == nil {
				t.Errorf("f11.0.f11.0.f6: expected %v, got nil", vals(eF11))
			} else {
				if len(sF11) != 2 || *(sF11)[0] != *(eF11)[0] || *(sF11)[1] != *(eF11)[1] {
					t.Errorf("f11.0.f11.0.f6: expected %v, got %v", vals(eF11), vals(sF11))
				}
			}
			sF11, eF11 = *s.F11[0].F11[1].F06, *e.F11[0].F11[1].F06
			if sF11 == nil {
				t.Errorf("f11.0.f11.0.f6: expected %v, got nil", vals(eF11))
			} else {
				if len(sF11) != 2 || *(sF11)[0] != *(eF11)[0] || *(sF11)[1] != *(eF11)[1] {
					t.Errorf("f11.0.f11.0.f6: expected %v, got %v", vals(eF11), vals(sF11))
				}
			}
		}
	}
	if s.F12 == nil {
		t.Errorf("f12: got nil")
	} else if len(*s.F12) != 1 {
		t.Errorf("f12: expected 1 element, got %v", *s.F12)
	} else {
		sF12, eF12 := *(s.F12), *(e.F12)
		if len(*sF12[0].F12) != 2 {
			t.Errorf("f12.0.f12: expected 1 element, got %v", *sF12[0].F12)
		} else {
			sF122, eF122 := *(*sF12[0].F12)[0].F06, *(*eF12[0].F12)[0].F06
			if sF122 == nil {
				t.Errorf("f12.0.f12.0.f6: expected %v, got nil", vals(eF122))
			} else {
				if len(sF122) != 2 || *(sF122)[0] != *(eF122)[0] || *(sF122)[1] != *(eF122)[1] {
					t.Errorf("f12.0.f12.0.f6: expected %v, got %v", vals(eF122), vals(sF122))
				}
			}
			sF122, eF122 = *(*sF12[0].F12)[1].F06, *(*eF12[0].F12)[1].F06
			if sF122 == nil {
				t.Errorf("f12.0.f12.0.f6: expected %v, got nil", vals(eF122))
			} else {
				if len(sF122) != 2 || *(sF122)[0] != *(eF122)[0] || *(sF122)[1] != *(eF122)[1] {
					t.Errorf("f12.0.f12.0.f6: expected %v, got %v", vals(eF122), vals(sF122))
				}
			}
		}
	}
	if s.F13 == nil {
		t.Errorf("f13: got nil")
	} else if len(*s.F13) != 1 {
		t.Errorf("f13: expected 1 element, got %v", *s.F13)
	} else {
		sF13, eF13 := *(s.F13), *(e.F13)
		if len(*sF13[0].F13) != 2 {
			t.Errorf("f13.0.f13: expected 1 element, got %v", *sF13[0].F13)
		} else {
			sF132, eF132 := *(*sF13[0].F13)[0].F06, *(*eF13[0].F13)[0].F06
			if sF132 == nil {
				t.Errorf("f13.0.f13.0.f6: expected %v, got nil", vals(eF132))
			} else {
				if len(sF132) != 2 || *(sF132)[0] != *(eF132)[0] || *(sF132)[1] != *(eF132)[1] {
					t.Errorf("f13.0.f13.0.f6: expected %v, got %v", vals(eF132), vals(sF132))
				}
			}
			sF132, eF132 = *(*sF13[0].F13)[1].F06, *(*eF13[0].F13)[1].F06
			if sF132 == nil {
				t.Errorf("f13.0.f13.0.f6: expected %v, got nil", vals(eF132))
			} else {
				if len(sF132) != 2 || *(sF132)[0] != *(eF132)[0] || *(sF132)[1] != *(eF132)[1] {
					t.Errorf("f13.0.f13.0.f6: expected %v, got %v", vals(eF132), vals(sF132))
				}
			}
		}
	}
}

// ----------------------------------------------------------------------------

type S12A struct {
	ID []int
}

func TestCSVSlice(t *testing.T) {
	data := map[string][]string{
		"ID": {"0,1"},
	}

	s := S12A{}
	err := NewDecoder().Decode(&s, data)
	if err != nil {
		t.Fatalf("Failed to decode: %v", err)
	}
	if len(s.ID) != 2 {
		t.Errorf("Expected two values in the result list, got %+v", s.ID)
	}
	if s.ID[0] != 0 || s.ID[1] != 1 {
		t.Errorf("Expected []{0, 1} got %+v", s)
	}
}

type S12B struct {
	ID []string
}

// Decode should not split on , into a slice for string only
func TestCSVStringSlice(t *testing.T) {
	data := map[string][]string{
		"ID": {"0,1"},
	}

	s := S12B{}
	err := NewDecoder().Decode(&s, data)
	if err != nil {
		t.Fatalf("Failed to decode: %v", err)
	}
	if len(s.ID) != 1 {
		t.Errorf("Expected one value in the result list, got %+v", s.ID)
	}
	if s.ID[0] != "0,1" {
		t.Errorf("Expected []{0, 1} got %+v", s)
	}
}

// Invalid data provided by client should not panic (github issue 33)
func TestInvalidDataProvidedByClient(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			t.Errorf("Panicked calling decoder.Decode: %v", r)
		}
	}()

	type S struct {
		f string // nolint:unused
	}

	data := map[string][]string{
		"f.f": {"v"},
	}

	err := NewDecoder().Decode(new(S), data)
	if err == nil {
		t.Errorf("invalid path in decoder.Decode should return an error.")
	}
}

// underlying cause of error in issue 33
func TestInvalidPathInCacheParsePath(t *testing.T) {
	type S struct {
		f string // nolint:unused
	}

	typ := reflect.ValueOf(new(S)).Elem().Type()
	c := newCache()
	_, err := c.parsePath("f.f", typ)
	if err == nil {
		t.Errorf("invalid path in cache.parsePath should return an error.")
	}
}

// issue 32
func TestDecodeToTypedField(t *testing.T) {
	type Aa bool
	s1 := &struct{ Aa }{}
	v1 := map[string][]string{"Aa": {"true"}}
	err := NewDecoder().Decode(s1, v1)
	if err != nil {
		t.Fatalf("Failed to decode: %v", err)
	}
	if s1.Aa != Aa(true) {
		t.Errorf("s1: expected %v, got %v", true, s1.Aa)
	}
}

// issue 37
func TestRegisterConverter(t *testing.T) {
	type Aa int
	type Bb int
	s1 := &struct {
		Aa
		Bb
	}{}
	decoder := NewDecoder()

	decoder.RegisterConverter(s1.Aa, func(s string) reflect.Value { return reflect.ValueOf(1) })
	decoder.RegisterConverter(s1.Bb, func(s string) reflect.Value { return reflect.ValueOf(2) })

	v1 := map[string][]string{"Aa": {"4"}, "Bb": {"5"}}
	err := decoder.Decode(s1, v1)
	if err != nil {
		t.Fatalf("Failed to decode: %v", err)
	}

	if s1.Aa != Aa(1) {
		t.Errorf("s1.Aa: expected %v, got %v", 1, s1.Aa)
	}
	if s1.Bb != Bb(2) {
		t.Errorf("s1.Bb: expected %v, got %v", 2, s1.Bb)
	}
}

// Issue #40
func TestRegisterConverterSlice(t *testing.T) {
	decoder := NewDecoder()
	decoder.RegisterConverter([]string{}, func(input string) reflect.Value {
		return reflect.ValueOf(strings.Split(input, ","))
	})

	result := struct {
		Multiple []string `schema:"multiple"`
	}{}

	expected := []string{"one", "two", "three"}
	err := decoder.Decode(&result, map[string][]string{
		"multiple": {"one,two,three"},
	})
	if err != nil {
		t.Fatalf("Failed to decode: %v", err)
	}
	for i := range expected {
		if got, want := expected[i], result.Multiple[i]; got != want {
			t.Errorf("%d: got %s, want %s", i, got, want)
		}
	}
}

func TestRegisterConverterMap(t *testing.T) {
	decoder := NewDecoder()
	decoder.IgnoreUnknownKeys(false)
	decoder.RegisterConverter(map[string]string{}, func(input string) reflect.Value {
		m := make(map[string]string)
		for _, pair := range strings.Split(input, ",") {
			parts := strings.Split(pair, ":")
			switch len(parts) {
			case 2:
				m[parts[0]] = parts[1]
			}
		}
		return reflect.ValueOf(m)
	})

	result := struct {
		Multiple map[string]string `schema:"multiple"`
	}{}

	err := decoder.Decode(&result, map[string][]string{
		"multiple": {"a:one,b:two"},
	})
	if err != nil {
		t.Fatal(err)
	}
	expected := map[string]string{"a": "one", "b": "two"}
	for k, v := range expected {
		got, ok := result.Multiple[k]
		if !ok {
			t.Fatalf("got %v, want %v", result.Multiple, expected)
		}
		if got != v {
			t.Errorf("got %s, want %s", got, v)
		}
	}
}

type S13 struct {
	Value []S14
}

type S14 struct {
	F1 string
	F2 string
}

func (n *S14) UnmarshalText(text []byte) error {
	textParts := strings.Split(string(text), " ")
	if len(textParts) < 2 {
		return errors.New("Not a valid name!")
	}

	n.F1, n.F2 = textParts[0], textParts[len(textParts)-1]
	return nil
}

type S15 struct {
	Value []S16
}

type S16 struct {
	F1 string
	F2 string
}

func TestCustomTypeSlice(t *testing.T) {
	data := map[string][]string{
		"Value.0": {"Louisa May Alcott"},
		"Value.1": {"Florence Nightingale"},
		"Value.2": {"Clara Barton"},
	}

	s := S13{}
	decoder := NewDecoder()

	if err := decoder.Decode(&s, data); err != nil {
		t.Fatal(err)
	}

	if len(s.Value) != 3 {
		t.Fatalf("Expected 3 values in the result list, got %+v", s.Value)
	}
	if s.Value[0].F1 != "Louisa" || s.Value[0].F2 != "Alcott" {
		t.Errorf("Expected S14{'Louisa', 'Alcott'} got %+v", s.Value[0])
	}
	if s.Value[1].F1 != "Florence" || s.Value[1].F2 != "Nightingale" {
		t.Errorf("Expected S14{'Florence', 'Nightingale'} got %+v", s.Value[1])
	}
	if s.Value[2].F1 != "Clara" || s.Value[2].F2 != "Barton" {
		t.Errorf("Expected S14{'Clara', 'Barton'} got %+v", s.Value[2])
	}
}

func TestCustomTypeSliceWithError(t *testing.T) {
	data := map[string][]string{
		"Value.0": {"Louisa May Alcott"},
		"Value.1": {"Florence Nightingale"},
		"Value.2": {"Clara"},
	}

	s := S13{}
	decoder := NewDecoder()

	if err := decoder.Decode(&s, data); err == nil {
		t.Error("Not detecting error in conversion")
	}
}

func TestNoTextUnmarshalerTypeSlice(t *testing.T) {
	data := map[string][]string{
		"Value.0": {"Louisa May Alcott"},
		"Value.1": {"Florence Nightingale"},
		"Value.2": {"Clara Barton"},
	}

	s := S15{}
	decoder := NewDecoder()

	if err := decoder.Decode(&s, data); err == nil {
		t.Error("Not detecting when there's no converter")
	}
}

// ----------------------------------------------------------------------------

type S17 struct {
	Value S14
}

type S18 struct {
	Value S16
}

func TestCustomType(t *testing.T) {
	data := map[string][]string{
		"Value": {"Louisa May Alcott"},
	}

	s := S17{}
	decoder := NewDecoder()

	if err := decoder.Decode(&s, data); err != nil {
		t.Fatal(err)
	}

	if s.Value.F1 != "Louisa" || s.Value.F2 != "Alcott" {
		t.Errorf("Expected S14{'Louisa', 'Alcott'} got %+v", s.Value)
	}
}

func TestCustomTypeWithError(t *testing.T) {
	data := map[string][]string{
		"Value": {"Louisa"},
	}

	s := S17{}
	decoder := NewDecoder()

	if err := decoder.Decode(&s, data); err == nil {
		t.Error("Not detecting error in conversion")
	}
}

func TestNoTextUnmarshalerType(t *testing.T) {
	data := map[string][]string{
		"Value": {"Louisa May Alcott"},
	}

	s := S18{}
	decoder := NewDecoder()

	if err := decoder.Decode(&s, data); err == nil {
		t.Error("Not detecting when there's no converter")
	}
}

func TestExpectedType(t *testing.T) {
	data := map[string][]string{
		"bools":   {"1", "a"},
		"date":    {"invalid"},
		"Foo.Bar": {"a", "b"},
	}

	type B struct {
		Bar *int
	}
	type A struct {
		Bools []bool    `schema:"bools"`
		Date  time.Time `schema:"date"`
		Foo   B
	}

	a := A{}

	err := NewDecoder().Decode(&a, data)

	e := err.(MultiError)["bools"].(ConversionError)
	if e.Type != reflect.TypeOf(false) && e.Index == 1 {
		t.Errorf("Expected bool, index: 1 got %+v, index: %d", e.Type, e.Index)
	}
	e = err.(MultiError)["date"].(ConversionError)
	if e.Type != reflect.TypeOf(time.Time{}) {
		t.Errorf("Expected time.Time got %+v", e.Type)
	}
	e = err.(MultiError)["Foo.Bar"].(ConversionError)
	if e.Type != reflect.TypeOf(0) {
		t.Errorf("Expected int got %+v", e.Type)
	}
}

type R1 struct {
	A string `schema:"a,required"`
	B struct {
		C int     `schema:"c,required"`
		D float64 `schema:"d"`
		E string  `schema:"e,required"`
	} `schema:"b"`
	F []string `schema:"f,required"`
	G []int    `schema:"g,othertag"`
	H bool     `schema:"h,required"`
}

func TestRequiredField(t *testing.T) {
	var a R1
	v := map[string][]string{
		"a":   {"bbb"},
		"b.c": {"88"},
		"b.d": {"9"},
		"f":   {""},
		"h":   {"true"},
	}
	err := NewDecoder().Decode(&a, v)
	if err == nil {
		t.Errorf("error nil, b.e is empty expect")
		return
	}
	// b.e empty
	v["b.e"] = []string{""} // empty string
	err = NewDecoder().Decode(&a, v)
	if err == nil {
		t.Errorf("error nil, b.e is empty expect")
		return
	}
	if expected := `b.e is empty`; err.Error() != expected {
		t.Errorf("got %q, want %q", err, expected)
	}

	// all fields ok
	v["b.e"] = []string{"nonempty"}
	err = NewDecoder().Decode(&a, v)
	if err != nil {
		t.Errorf("error: %v", err)
		return
	}

	// set f empty
	v["f"] = []string{}
	err = NewDecoder().Decode(&a, v)
	if err == nil {
		t.Errorf("error nil, f is empty expect")
		return
	}
	if expected := `f is empty`; err.Error() != expected {
		t.Errorf("got %q, want %q", err, expected)
	}
	v["f"] = []string{"nonempty"}

	// b.c type int with empty string
	v["b.c"] = []string{""}
	err = NewDecoder().Decode(&a, v)
	if err == nil {
		t.Errorf("error nil, b.c is empty expect")
		return
	}
	v["b.c"] = []string{"3"}

	// h type bool with empty string
	v["h"] = []string{""}
	err = NewDecoder().Decode(&a, v)
	if err == nil {
		t.Errorf("error nil, h is empty expect")
		return
	}
	if expected := `h is empty`; err.Error() != expected {
		t.Errorf("got %q, want %q", err, expected)
	}
}

type R2 struct {
	A struct {
		B int `schema:"b"`
	} `schema:"a,required"`
}

func TestRequiredStructFiled(t *testing.T) {
	v := map[string][]string{
		"a.b": {"3"},
	}
	var a R2
	err := NewDecoder().Decode(&a, v)
	if err != nil {
		t.Errorf("error: %v", err)
	}
}

type Node struct {
	Value int   `schema:"val,required"`
	Next  *Node `schema:"next,required"`
}

func TestRecursiveStruct(t *testing.T) {
	v := map[string][]string{
		"val":      {"1"},
		"next.val": {"2"},
	}
	var a Node
	err := NewDecoder().Decode(&a, v)
	if err != nil {
		t.Errorf("error: %v", err)
	}
}

func TestRequiredFieldIsMissingCorrectError(t *testing.T) {
	type RM1S struct {
		A string `schema:"rm1aa,required"`
		B string `schema:"rm1bb,required"`
	}
	type RM1 struct {
		RM1S
	}

	var a RM1
	v := map[string][]string{
		"rm1aa": {"aaa"},
	}
	expectedError := "RM1S.rm1bb is empty"
	err := NewDecoder().Decode(&a, v)
	if err.Error() != expectedError {
		t.Errorf("expected %v, got %v", expectedError, err)
	}
}

type AS1 struct {
	A int32 `schema:"a,required"`
	E int32 `schema:"e,required"`
}
type AS2 struct {
	AS1
	B string `schema:"b,required"`
}
type AS3 struct {
	C int32 `schema:"c"`
}

type AS4 struct {
	AS3
	D string `schema:"d"`
}

func TestAnonymousStructField(t *testing.T) {
	patterns := []map[string][]string{
		{
			"a": {"1"},
			"e": {"2"},
			"b": {"abc"},
		},
		{
			"AS1.a": {"1"},
			"AS1.e": {"2"},
			"b":     {"abc"},
		},
	}
	for _, v := range patterns {
		a := AS2{}
		err := NewDecoder().Decode(&a, v)
		if err != nil {
			t.Errorf("Decode failed %s, %#v", err, v)
			continue
		}
		if a.A != 1 {
			t.Errorf("A: expected %v, got %v", 1, a.A)
		}
		if a.E != 2 {
			t.Errorf("E: expected %v, got %v", 2, a.E)
		}
		if a.B != "abc" {
			t.Errorf("B: expected %v, got %v", "abc", a.B)
		}
		embedded := a.AS1
		if embedded.A != 1 {
			t.Errorf("AS1.A: expected %v, got %v", 1, embedded.A)
		}
		if embedded.E != 2 {
			t.Errorf("AS1.E: expected %v, got %v", 2, embedded.E)
		}
	}
	a := AS2{}
	err := NewDecoder().Decode(&a, map[string][]string{
		"e": {"2"},
		"b": {"abc"},
	})
	if err == nil {
		t.Errorf("error nil, a is empty expect")
	}
	patterns = []map[string][]string{
		{
			"c": {"1"},
			"d": {"abc"},
		},
		{
			"AS3.c": {"1"},
			"d":     {"abc"},
		},
	}
	for _, v := range patterns {
		a := AS4{}
		err := NewDecoder().Decode(&a, v)
		if err != nil {
			t.Errorf("Decode failed %s, %#v", err, v)
			continue
		}
		if a.C != 1 {
			t.Errorf("C: expected %v, got %v", 1, a.C)
		}
		if a.D != "abc" {
			t.Errorf("D: expected %v, got %v", "abc", a.D)
		}
		if embedded := a.AS3; embedded.C != 1 {
			t.Errorf("AS3.C: expected %v, got %v", 1, embedded.C)
		}
	}
}

func TestAmbiguousStructField(t *testing.T) {
	type I1 struct {
		X int
	}
	type I2 struct {
		I1
	}
	type B1 struct {
		X bool
	}
	type B2 struct {
		B1
	}
	type IB struct {
		I1
		B1
	}
	type S struct {
		I1
		I2
		B1
		B2
		IB
	}
	dst := S{}
	src := map[string][]string{
		"X":    {"123"},
		"IB.X": {"123"},
	}
	dec := NewDecoder()
	dec.IgnoreUnknownKeys(false)
	err := dec.Decode(&dst, src)
	e, ok := err.(MultiError)
	if !ok || len(e) != 2 {
		t.Errorf("Expected 2 errors, got %#v", err)
	}
	if expected := (UnknownKeyError{Key: "X"}); e["X"] != expected {
		t.Errorf("X: expected %#v, got %#v", expected, e["X"])
	}
	if expected := (UnknownKeyError{Key: "IB.X"}); e["IB.X"] != expected {
		t.Errorf("X: expected %#v, got %#v", expected, e["IB.X"])
	}
	dec.IgnoreUnknownKeys(true)
	err = dec.Decode(&dst, src)
	if err != nil {
		t.Errorf("Decode failed %v", err)
	}

	expected := S{
		I1: I1{X: 123},
		I2: I2{I1: I1{X: 234}},
		B1: B1{X: true},
		B2: B2{B1: B1{X: true}},
		IB: IB{I1: I1{X: 345}, B1: B1{X: true}},
	}
	patterns := []map[string][]string{
		{
			"I1.X":    {"123"},
			"I2.X":    {"234"},
			"B1.X":    {"true"},
			"B2.X":    {"1"},
			"IB.I1.X": {"345"},
			"IB.B1.X": {"on"},
		},
		{
			"I1.X":    {"123"},
			"I2.I1.X": {"234"},
			"B1.X":    {"true"},
			"B2.B1.X": {"1"},
			"IB.I1.X": {"345"},
			"IB.B1.X": {"on"},
		},
	}
	for _, src := range patterns {
		dst := S{}
		dec := NewDecoder()
		dec.IgnoreUnknownKeys(false)
		err := dec.Decode(&dst, src)
		if err != nil {
			t.Errorf("Decode failed %v, %#v", err, src)
		}
		if !reflect.DeepEqual(expected, dst) {
			t.Errorf("Expected %+v, got %+v", expected, dst)
		}
	}
}

func TestComprehensiveDecodingErrors(t *testing.T) {
	type I1 struct {
		V int  `schema:",required"`
		P *int `schema:",required"`
	}
	type I2 struct {
		I1
		J I1
	}
	type S1 struct {
		V string  `schema:"v,required"`
		P *string `schema:"p,required"`
	}
	type S2 struct {
		S1 `schema:"s"`
		T  S1 `schema:"t"`
	}
	type D struct {
		I2
		X S2 `schema:"x"`
		Y S2 `schema:"-"`
	}
	patterns := []map[string][]string{
		{
			"V":       {"invalid"}, // invalid
			"I2.I1.P": {},          // empty
			"I2.J.V":  {""},        // empty
			"I2.J.P":  {"123"},     // ok
			"x.s.v":   {""},        // empty
			"x.s.p":   {""},        // ok
			"x.t.v":   {"abc"},     // ok
			"x.t.p":   {},          // empty
			"Y.s.v":   {"ignored"}, // unknown
		},
		{
			"V":     {"invalid"}, // invalid
			"P":     {},          // empty
			"J.V":   {""},        // empty
			"J.P":   {"123"},     // ok
			"x.v":   {""},        // empty
			"x.p":   {""},        // ok
			"x.t.v": {"abc"},     // ok
			"x.t.p": {},          // empty
			"Y.s.v": {"ignored"}, // unknown
		},
	}
	for _, src := range patterns {
		dst := D{}
		dec := NewDecoder()
		dec.IgnoreUnknownKeys(false)
		err := dec.Decode(&dst, src)
		e, ok := err.(MultiError)
		if !ok || len(e) != 6 {
			t.Errorf("Expected 6 errors, got %#v", err)
		}
		if cerr, ok := e["V"].(ConversionError); !ok {
			t.Errorf("%s: expected %#v, got %#v", "I2.I1.V", ConversionError{Key: "V"}, cerr)
		}
		if key, expected := "I2.I1.P", (EmptyFieldError{Key: "I2.I1.P"}); e[key] != expected {
			t.Errorf("%s: expected %#v, got %#v", key, expected, e[key])
		}
		if key, expected := "I2.J.V", (EmptyFieldError{Key: "I2.J.V"}); e[key] != expected {
			t.Errorf("%s: expected %#v, got %#v", key, expected, e[key])
		}
		if key, expected := "x.s.v", (EmptyFieldError{Key: "x.s.v"}); e[key] != expected {
			t.Errorf("%s: expected %#v, got %#v", key, expected, e[key])
		}
		if key, expected := "x.t.p", (EmptyFieldError{Key: "x.t.p"}); e[key] != expected {
			t.Errorf("%s: expected %#v, got %#v", key, expected, e[key])
		}
		if key, expected := "Y.s.v", (UnknownKeyError{Key: "Y.s.v"}); e[key] != expected {
			t.Errorf("%s: expected %#v, got %#v", key, expected, e[key])
		}
		if expected := 123; dst.J.P == nil || *dst.J.P != expected {
			t.Errorf("I2.J.P: expected %#v, got %#v", expected, dst.J.P)
		}
		if expected := ""; dst.X.P == nil || *dst.X.P != expected {
			t.Errorf("X.S1.P: expected %#v, got %#v", expected, dst.X.P)
		}
		if expected := "abc"; dst.X.T.V != expected {
			t.Errorf("X.T.V: expected %#v, got %#v", expected, dst.X.T.V)
		}
	}
}

// Test to ensure that a registered converter overrides the default text unmarshaler.
func TestRegisterConverterOverridesTextUnmarshaler(t *testing.T) {
	type MyTime time.Time
	s1 := &struct {
		MyTime
	}{}
	decoder := NewDecoder()

	ts := time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC)
	decoder.RegisterConverter(s1.MyTime, func(s string) reflect.Value { return reflect.ValueOf(ts) })

	v1 := map[string][]string{"MyTime": {"4"}}
	err := decoder.Decode(s1, v1)
	if err != nil {
		t.Fatalf("Failed to decode: %v", err)
	}

	if s1.MyTime != MyTime(ts) {
		t.Errorf("s1.Aa: expected %v, got %v", ts, s1.MyTime)
	}
}

type S20E string

func (e *S20E) UnmarshalText(text []byte) error {
	*e = S20E("x")
	return nil
}

type S20 []S20E

func (s *S20) UnmarshalText(text []byte) error {
	*s = S20{"a", "b", "c"}
	return nil
}

// Test to ensure that when a custom type based on a slice implements an
// encoding.TextUnmarshaler interface that it takes precedence over any
// implementations by its elements.
func TestTextUnmarshalerTypeSlice(t *testing.T) {
	data := map[string][]string{
		"Value": {"a,b,c"},
	}
	s := struct {
		Value S20
	}{}
	decoder := NewDecoder()
	if err := decoder.Decode(&s, data); err != nil {
		t.Fatal("Error while decoding:", err)
	}
	expected := S20{"a", "b", "c"}
	if !reflect.DeepEqual(expected, s.Value) {
		t.Errorf("Expected %v errors, got %v", expected, s.Value)
	}
}

type S21E struct{ ElementValue string }

func (e *S21E) UnmarshalText(text []byte) error {
	*e = S21E{"x"}
	return nil
}

type S21 []S21E

func (s *S21) UnmarshalText(text []byte) error {
	*s = S21{{"a"}}
	return nil
}

type S21B []S21E

// Test to ensure that if custom type base on a slice of structs implements an
// encoding.TextUnmarshaler interface it is unaffected by the special path
// requirements imposed on a slice of structs.
func TestTextUnmarshalerTypeSliceOfStructs(t *testing.T) {
	data := map[string][]string{
		"Value": {"raw a"},
	}
	// Implements encoding.TextUnmarshaler, should not throw invalid path
	// error.
	s := struct {
		Value S21
	}{}
	decoder := NewDecoder()
	if err := decoder.Decode(&s, data); err != nil {
		t.Fatal("Error while decoding:", err)
	}
	expected := S21{{"a"}}
	if !reflect.DeepEqual(expected, s.Value) {
		t.Errorf("Expected %v errors, got %v", expected, s.Value)
	}
	// Does not implement encoding.TextUnmarshaler, should throw invalid
	// path error.
	sb := struct {
		Value S21B
	}{}
	if err := decoder.Decode(&sb, data); err == errInvalidPath {
		t.Fatal("Expecting invalid path error", err)
	}
}

type S22 string

func (s *S22) UnmarshalText(text []byte) error {
	*s = S22("a")
	return nil
}

// Test to ensure that when a field that should be decoded into a type
// implementing the encoding.TextUnmarshaler interface is set to an empty value
// that the UnmarshalText method is utilized over other methods of decoding,
// especially including simply setting the zero value.
func TestTextUnmarshalerEmpty(t *testing.T) {
	data := map[string][]string{
		"Value": {""}, // empty value
	}
	// Implements encoding.TextUnmarshaler, should use the type's
	// UnmarshalText method.
	s := struct {
		Value S22
	}{}
	decoder := NewDecoder()
	if err := decoder.Decode(&s, data); err != nil {
		t.Fatal("Error while decoding:", err)
	}
	expected := S22("a")
	if expected != s.Value {
		t.Errorf("Expected %v errors, got %v", expected, s.Value)
	}
}

type S23n struct {
	F2 string `schema:"F2"`
	F3 string `schema:"F3"`
}

type S23e struct {
	*S23n
	F1 string `schema:"F1"`
}

type S23 []*S23e

func TestUnmashalPointerToEmbedded(t *testing.T) {
	data := map[string][]string{
		"A.0.F2": {"raw a"},
		"A.0.F3": {"raw b"},
	}

	// Implements encoding.TextUnmarshaler, should not throw invalid path
	// error.
	s := struct {
		Value S23 `schema:"A"`
	}{}
	decoder := NewDecoder()

	if err := decoder.Decode(&s, data); err != nil {
		t.Fatal("Error while decoding:", err)
	}

	expected := S23{
		&S23e{
			S23n: &S23n{"raw a", "raw b"},
		},
	}
	if !reflect.DeepEqual(expected, s.Value) {
		t.Errorf("Expected %v errors, got %v", expected, s.Value)
	}
}

type S24 struct {
	F1 string `schema:"F1"`
}

type S24e struct {
	*S24
	F2 string `schema:"F2"`
}

func TestUnmarshallToEmbeddedNoData(t *testing.T) {
	data := map[string][]string{
		"F3": {"raw a"},
	}

	s := &S24e{}

	decoder := NewDecoder()
	err := decoder.Decode(s, data)

	expectedErr := `schema: invalid path "F3"`
	if err.Error() != expectedErr {
		t.Fatalf("got %q, want %q", err, expectedErr)
	}
}

type S25ee struct {
	F3 string `schema:"F3"`
}

type S25e struct {
	S25ee
	F2 string `schema:"F2"`
}

type S25 struct {
	S25e
	F1 string `schema:"F1"`
}

func TestDoubleEmbedded(t *testing.T) {
	data := map[string][]string{
		"F1": {"raw a"},
		"F2": {"raw b"},
		"F3": {"raw c"},
	}

	s := S25{}
	decoder := NewDecoder()

	if err := decoder.Decode(&s, data); err != nil {
		t.Fatal("Error while decoding:", err)
	}

	expected := S25{
		F1: "raw a",
		S25e: S25e{
			F2: "raw b",
			S25ee: S25ee{
				F3: "raw c",
			},
		},
	}
	if !reflect.DeepEqual(expected, s) {
		t.Errorf("Expected %v errors, got %v", expected, s)
	}
}

func TestDefaultValuesAreSet(t *testing.T) {
	type N struct {
		S1 string    `schema:"s1,default:test1"`
		I2 int       `schema:"i2,default:22"`
		R2 []float64 `schema:"r2,default:2|3.5|11.01"`
	}

	type D struct {
		N
		S string   `schema:"s,default:test1"`
		I int      `schema:"i,default:21"`
		J int8     `schema:"j,default:2"`
		K int16    `schema:"k,default:-455"`
		L int32    `schema:"l,default:899"`
		M int64    `schema:"m,default:12455"`
		B bool     `schema:"b,default:false"`
		F float64  `schema:"f,default:3.14"`
		G float32  `schema:"g,default:19.12"`
		U uint     `schema:"u,default:1"`
		V uint8    `schema:"v,default:190"`
		W uint16   `schema:"w,default:20000"`
		Y uint32   `schema:"y,default:156666666"`
		Z uint64   `schema:"z,default:1545465465465546"`
		X []string `schema:"x,default:x1|x2"`
	}

	data := map[string][]string{}

	d := D{}

	decoder := NewDecoder()

	if err := decoder.Decode(&d, data); err != nil {
		t.Fatal("Error while decoding:", err)
	}

	expected := D{
		N: N{
			S1: "test1",
			I2: 22,
			R2: []float64{2, 3.5, 11.01},
		},
		S: "test1",
		I: 21,
		J: 2,
		K: -455,
		L: 899,
		M: 12455,
		B: false,
		F: 3.14,
		G: 19.12,
		U: 1,
		V: 190,
		W: 20000,
		Y: 156666666,
		Z: 1545465465465546,
		X: []string{"x1", "x2"},
	}

	if !reflect.DeepEqual(expected, d) {
		t.Errorf("Expected %v, got %v", expected, d)
	}

	type P struct {
		*N
		S *string  `schema:"s,default:test1"`
		I *int     `schema:"i,default:21"`
		J *int8    `schema:"j,default:2"`
		K *int16   `schema:"k,default:-455"`
		L *int32   `schema:"l,default:899"`
		M *int64   `schema:"m,default:12455"`
		B *bool    `schema:"b,default:false"`
		F *float64 `schema:"f,default:3.14"`
		G *float32 `schema:"g,default:19.12"`
		U *uint    `schema:"u,default:1"`
		V *uint8   `schema:"v,default:190"`
		W *uint16  `schema:"w,default:20000"`
		Y *uint32  `schema:"y,default:156666666"`
		Z *uint64  `schema:"z,default:1545465465465546"`
		X []string `schema:"x,default:x1|x2"`
	}

	p := P{N: &N{}}

	if err := decoder.Decode(&p, data); err != nil {
		t.Fatal("Error while decoding:", err)
	}

	vExpected := reflect.ValueOf(expected)
	vActual := reflect.ValueOf(p)

	i := 0

	for i < vExpected.NumField() {
		if !reflect.DeepEqual(vExpected.Field(i).Interface(), reflect.Indirect(vActual.Field(i)).Interface()) {
			t.Errorf("Expected %v, got %v", vExpected.Field(i).Interface(), reflect.Indirect(vActual.Field(i)).Interface())
		}
		i++
	}
}

func TestDefaultValuesAreIgnoredIfValuesAreProvided(t *testing.T) {
	type D struct {
		S string  `schema:"s,default:test1"`
		I int     `schema:"i,default:21"`
		B bool    `schema:"b,default:false"`
		F float64 `schema:"f,default:3.14"`
		U uint    `schema:"u,default:1"`
	}

	data := map[string][]string{"s": {"s"}, "i": {"1"}, "b": {"true"}, "f": {"0.22"}, "u": {"14"}}

	d := D{}

	decoder := NewDecoder()

	if err := decoder.Decode(&d, data); err != nil {
		t.Fatal("Error while decoding:", err)
	}

	expected := D{
		S: "s",
		I: 1,
		B: true,
		F: 0.22,
		U: 14,
	}

	if !reflect.DeepEqual(expected, d) {
		t.Errorf("Expected %v, got %v", expected, d)
	}
}

func TestRequiredFieldsCannotHaveDefaults(t *testing.T) {
	type D struct {
		S string  `schema:"s,required,default:test1"`
		I int     `schema:"i,required,default:21"`
		B bool    `schema:"b,required,default:false"`
		F float64 `schema:"f,required,default:3.14"`
		U uint    `schema:"u,required,default:1"`
	}

	data := map[string][]string{"s": {"s"}, "i": {"1"}, "b": {"true"}, "f": {"0.22"}, "u": {"14"}}

	d := D{}

	decoder := NewDecoder()

	err := decoder.Decode(&d, data)

	expected := "required fields cannot have a default value"

	if err == nil || !strings.Contains(err.Error(), expected) {
		t.Errorf("decoding should fail with error msg %s got %q", expected, err)
	}
}

func TestInvalidDefaultElementInSliceRaiseError(t *testing.T) {
	type D struct {
		A []int  `schema:"a,default:0|notInt"`
		B []bool `schema:"b,default:true|notInt"`
		// //uint types
		D []uint   `schema:"d,default:1|notInt"`
		E []uint8  `schema:"e,default:2|notInt"`
		F []uint16 `schema:"f,default:3|notInt"`
		G []uint32 `schema:"g,default:4|notInt"`
		H []uint64 `schema:"h,default:5|notInt"`
		// // int types
		N []int   `schema:"n,default:11|notInt"`
		O []int8  `schema:"o,default:12|notInt"`
		P []int16 `schema:"p,default:13|notInt"`
		Q []int32 `schema:"q,default:14|notInt"`
		R []int64 `schema:"r,default:15|notInt"`
		// // float
		X []float32 `schema:"c,default:2.2|notInt"`
		Y []float64 `schema:"c,default:3.3|notInt"`
	}
	d := D{}

	data := map[string][]string{}

	decoder := NewDecoder()

	err := decoder.Decode(&d, data)

	if err == nil {
		t.Error("if a different type exists, error should be raised")
	}

	dType := reflect.TypeOf(d)

	e, ok := err.(MultiError)
	if !ok || len(e) != dType.NumField() {
		t.Errorf("Expected %d errors, got %#v", dType.NumField(), err)
	}

	for i := 0; i < dType.NumField(); i++ {
		v := dType.Field(i)
		fieldKey := "default-" + string(v.Name)
		errMsg := fmt.Sprintf("failed setting default: notInt is not compatible with field %s type", string(v.Name))
		ferr := e[fieldKey]
		if strings.Compare(ferr.Error(), errMsg) != 0 {
			t.Errorf("%s: expected %s, got %#v\n", fieldKey, ferr.Error(), errMsg)
		}
	}
}

func TestInvalidDefaultsValuesHaveNoEffect(t *testing.T) {
	type D struct {
		B bool     `schema:"b,default:invalid"`
		C *float32 `schema:"c,default:notAFloat"`
		// uint types
		D uint   `schema:"d,default:notUint"`
		E uint8  `schema:"e,default:notUint"`
		F uint16 `schema:"f,default:notUint"`
		G uint32 `schema:"g,default:notUint"`
		H uint64 `schema:"h,default:notUint"`
		// uint types pointers
		I *uint   `schema:"i,default:notUint"`
		J *uint8  `schema:"j,default:notUint"`
		K *uint16 `schema:"k,default:notUint"`
		L *uint32 `schema:"l,default:notUint"`
		M *uint64 `schema:"m,default:notUint"`
		// int types
		N int   `schema:"n,default:notInt"`
		O int8  `schema:"o,default:notInt"`
		P int16 `schema:"p,default:notInt"`
		Q int32 `schema:"q,default:notInt"`
		R int64 `schema:"r,default:notInt"`
		// int types pointers
		S *int   `schema:"s,default:notInt"`
		T *int8  `schema:"t,default:notInt"`
		U *int16 `schema:"u,default:notInt"`
		V *int32 `schema:"v,default:notInt"`
		W *int64 `schema:"w,default:notInt"`
		// float
		X float32  `schema:"c,default:notAFloat"`
		Y float64  `schema:"c,default:notAFloat"`
		Z *float64 `schema:"c,default:notAFloat"`
	}

	d := D{}

	expected := D{}

	data := map[string][]string{}

	decoder := NewDecoder()

	err := decoder.Decode(&d, data)
	if err != nil {
		t.Errorf("decoding should succeed but got error: %q", err)
	}

	if !reflect.DeepEqual(expected, d) {
		t.Errorf("expected %v but got %v", expected, d)
	}
}

func TestDefaultsAreNotSupportedForStructsAndStructSlices(t *testing.T) {
	type C struct {
		C string `schema:"c"`
	}

	type D struct {
		S S1     `schema:"s,default:{f1:0}"`
		A []C    `schema:"a,default:{c:test1}|{c:test2}"`
		B []*int `schema:"b,default:12"`
		E *C     `schema:"e,default:{c:test3}"`
	}

	d := D{}

	data := map[string][]string{}

	decoder := NewDecoder()

	err := decoder.Decode(&d, data)

	expected := "default option is supported only on: bool, float variants, string, unit variants types or their corresponding pointers or slices"

	if err == nil || !strings.Contains(err.Error(), expected) {
		t.Errorf("decoding should fail with error msg %s got %q", expected, err)
	}
}

func TestDefaultValueWithColon(t *testing.T) {
	t.Parallel()
	type D struct {
		URL string `schema:"url,default:http://localhost:8080"`
	}

	var d D
	decoder := NewDecoder()
	if err := decoder.Decode(&d, map[string][]string{}); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if d.URL != "http://localhost:8080" {
		t.Errorf("expected default url to be http://localhost:8080, got %s", d.URL)
	}
}

func TestDecoder_MaxSize(t *testing.T) {
	t.Parallel()

	type Nested struct {
		Val          int
		NestedValues []struct {
			NVal int
		}
	}
	type NestedSlices struct {
		Values []Nested
	}

	testcases := []struct {
		name            string
		maxSize         int
		decoderInput    func() (dst NestedSlices, src map[string][]string)
		expectedDecoded NestedSlices
		expectedErr     MultiError
	}{
		{
			name:    "no error on decoding under max size",
			maxSize: 10,
			decoderInput: func() (dst NestedSlices, src map[string][]string) {
				return dst, map[string][]string{
					"Values.1.Val":                 {"132"},
					"Values.1.NestedValues.1.NVal": {"1"},
					"Values.1.NestedValues.2.NVal": {"2"},
					"Values.1.NestedValues.3.NVal": {"3"},
				}
			},
			expectedDecoded: NestedSlices{
				Values: []Nested{
					{
						Val:          0,
						NestedValues: nil,
					},
					{
						Val: 132, NestedValues: []struct{ NVal int }{
							{NVal: 0},
							{NVal: 1},
							{NVal: 2},
							{NVal: 3},
						},
					},
				},
			},
			expectedErr: nil,
		},
		{
			name:    "error on decoding above max size",
			maxSize: 1,
			decoderInput: func() (dst NestedSlices, src map[string][]string) {
				return dst, map[string][]string{
					"Values.1.Val":                 {"132"},
					"Values.1.NestedValues.1.NVal": {"1"},
					"Values.1.NestedValues.2.NVal": {"2"},
					"Values.1.NestedValues.3.NVal": {"3"},
				}
			},
			expectedErr: MultiError{
				"Values.1.NestedValues.2.NVal": errors.New("slice index 2 is larger than the configured maxSize 1"),
				"Values.1.NestedValues.3.NVal": errors.New("slice index 3 is larger than the configured maxSize 1"),
			},
		},
	}

	for _, tc := range testcases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			dec := NewDecoder()
			dec.MaxSize(tc.maxSize)
			dst, src := tc.decoderInput()
			err := dec.Decode(&dst, src)

			if tc.expectedErr != nil {
				var gotErr MultiError
				if !errors.As(err, &gotErr) {
					t.Errorf("decoder error is not of type %T", gotErr)
				}
				if !reflect.DeepEqual(gotErr, tc.expectedErr) {
					t.Errorf("expected %v, got %v", tc.expectedErr, gotErr)
				}
			} else {
				if !reflect.DeepEqual(dst, tc.expectedDecoded) {
					t.Errorf("expected %v, got %v", tc.expectedDecoded, dst)
				}
			}
		})
	}
}

func TestDefaultsAppliedOnlyWhenMissing(t *testing.T) {
	t.Parallel()

	type Data struct {
		B bool    `schema:"b,default:true"`
		I int     `schema:"i,default:5"`
		F float64 `schema:"f,default:1.5"`
		S []int   `schema:"s,default:1|2"`
	}

	dec := NewDecoder()

	// Values are explicitly set – no defaults should be applied
	withVals := Data{}
	if err := dec.Decode(&withVals, map[string][]string{
		"b": {"false"},
		"i": {"0"},
		"f": {"0"},
		"s": {},
	}); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if withVals.B {
		t.Errorf("B should be false when the value is set")
	}
	if withVals.I != 0 {
		t.Errorf("I should be 0 when the value is set, got %d", withVals.I)
	}
	if withVals.F != 0 {
		t.Errorf("F should be 0 when the value is set, got %f", withVals.F)
	}
	if len(withVals.S) != 0 {
		t.Errorf("S should be empty when the value is set, got %v", withVals.S)
	}

	// No values provided – defaults should be applied
	withoutVals := Data{}
	if err := dec.Decode(&withoutVals, map[string][]string{}); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !withoutVals.B {
		t.Errorf("B should default to true when missing")
	}
	if withoutVals.I != 5 {
		t.Errorf("Default I should be 5, got %d", withoutVals.I)
	}
	if withoutVals.F != 1.5 {
		t.Errorf("Default F should be 1.5, got %f", withoutVals.F)
	}
	if !reflect.DeepEqual(withoutVals.S, []int{1, 2}) {
		t.Errorf("Default S should be [1 2], got %v", withoutVals.S)
	}
}

func TestDecoder_SetMaxSize(t *testing.T) {
	t.Run("default maxsize should be equal to given constant", func(t *testing.T) {
		t.Parallel()
		dec := NewDecoder()
		if !reflect.DeepEqual(dec.maxSize, defaultMaxSize) {
			t.Errorf("unexpected default max size")
		}
	})

	t.Run("configured maxsize should be set properly", func(t *testing.T) {
		t.Parallel()
		configuredMaxSize := 50
		limitedMaxSizeDecoder := NewDecoder()
		limitedMaxSizeDecoder.MaxSize(configuredMaxSize)
		if !reflect.DeepEqual(limitedMaxSizeDecoder.maxSize, configuredMaxSize) {
			t.Errorf("invalid decoder maxsize, expected: %d, got: %d",
				configuredMaxSize, limitedMaxSizeDecoder.maxSize)
		}
	})
}

func TestTimeDurationDecoding(t *testing.T) {
	type DurationStruct struct {
		Timeout time.Duration `schema:"timeout"`
	}

	// Prepare the input data
	input := map[string][]string{
		"timeout": {"2s"},
	}

	// Create a decoder with a converter for time.Duration
	decoder := NewDecoder()
	decoder.RegisterConverter(time.Duration(0), func(s string) reflect.Value {
		d, err := time.ParseDuration(s)
		if err != nil {
			return reflect.Value{}
		}
		return reflect.ValueOf(d)
	})

	var result DurationStruct
	err := decoder.Decode(&result, input)
	if err != nil {
		t.Fatalf("Failed to decode duration: %v", err)
	}

	// Expect 2 seconds
	if result.Timeout != 2*time.Second {
		t.Errorf("Expected 2s, got %v", result.Timeout)
	}
}

func TestTimeDurationDecodingInvalid(t *testing.T) {
	type DurationStruct struct {
		Timeout time.Duration `schema:"timeout"`
	}

	// Prepare the input data
	input := map[string][]string{
		"timeout": {"invalid-duration"},
	}

	// Create a decoder with a converter for time.Duration
	decoder := NewDecoder()
	decoder.RegisterConverter(time.Duration(0), func(s string) reflect.Value {
		// Attempt to parse the duration
		d, err := time.ParseDuration(s)
		if err != nil {
			// Return an invalid reflect.Value to trigger a conversion error
			return reflect.Value{}
		}
		return reflect.ValueOf(d)
	})

	var result DurationStruct
	err := decoder.Decode(&result, input)
	if err == nil {
		t.Error("Expected an error decoding invalid duration, got nil")
	}
}

func TestMultipleConversionErrors(t *testing.T) {
	type Fields struct {
		IntField  int           `schema:"int_field"`
		BoolField bool          `schema:"bool_field"`
		Duration  time.Duration `schema:"duration_field"`
	}

	input := map[string][]string{
		"int_field":      {"invalid-int"},
		"bool_field":     {"invalid-bool"},
		"duration_field": {"invalid-duration"},
	}

	decoder := NewDecoder()
	decoder.RegisterConverter(time.Duration(0), func(s string) reflect.Value {
		d, err := time.ParseDuration(s)
		if err != nil {
			return reflect.Value{}
		}
		return reflect.ValueOf(d)
	})

	var s Fields
	err := decoder.Decode(&s, input)
	if err == nil {
		t.Fatal("Expected multiple conversion errors, got nil")
	}

	// Check that all errors are reported (at least 3).
	mErr, ok := err.(MultiError)
	if !ok {
		t.Fatalf("Expected MultiError, got %T", err)
	}
	if len(mErr) < 3 {
		t.Errorf("Expected at least 3 errors, got %d: %v", len(mErr), mErr)
	}
}

func TestDecoderMultipartFiles(t *testing.T) {
	type S struct {
		A string `schema:"a,required"`
		B int    `schema:"b,required"`
		C bool   `schema:"c,required"`
		D struct {
			E  float64                  `schema:"e,required"`
			F  *multipart.FileHeader    `schema:"f,required"`
			F2 []*multipart.FileHeader  `schema:"f2,required"`
			F3 *[]*multipart.FileHeader `schema:"f3,required"`
			F4 *multipart.FileHeader    `schema:"f4,required"`
		} `schema:"d,required"`
		G *[]*multipart.FileHeader `schema:"g,required"`
		J []struct {
			K *[]*multipart.FileHeader `schema:"k,required"`
		} `schema:"j,required"`
	}
	s := S{}
	data := map[string][]string{
		"a":   {"abc"},
		"b":   {"123"},
		"c":   {"true"},
		"d.e": {"3.14"},
	}

	// Create dummy file headers for testing
	dummyFile := &multipart.FileHeader{
		Filename: "test.txt",
		Size:     4,
	}

	dummyFile2 := &multipart.FileHeader{
		Filename: "test2.txt",
		Size:     4,
	}

	dummyFile3 := &multipart.FileHeader{
		Filename: "test3.txt",
		Size:     4,
	}

	// Create slice for file headers
	fileHeaders := map[string][]*multipart.FileHeader{
		"d.f":   {dummyFile, dummyFile2},
		"d.f2":  {dummyFile2, dummyFile3},
		"d.f3":  {dummyFile, dummyFile2, dummyFile3},
		"d.f4":  {},
		"g":     {dummyFile, dummyFile2},
		"j.0.k": {dummyFile, dummyFile2},
		"j.1.k": {dummyFile2, dummyFile3},
	}

	decoder := NewDecoder()
	err := decoder.Decode(&s, data, fileHeaders)
	if err != nil {
		t.Fatalf("Failed to decode: %v", err)
	}

	if s.A != "abc" {
		t.Errorf("Expected A to be 'abc', got %s", s.A)
	}

	if s.B != 123 {
		t.Errorf("Expected B to be 123, got %d", s.B)
	}

	if s.C != true {
		t.Errorf("Expected C to be true, got %t", s.C)
	}

	if s.D.E != 3.14 {
		t.Errorf("Expected D.E to be 3.14, got %f", s.D.E)
	}

	if s.D.F == nil {
		t.Error("Expected D.F to be a file header, got nil")
	}

	if s.D.F2 == nil {
		t.Error("Expected D.F2 to be a slice of file headers, got nil")
	}

	if s.D.F3 == nil {
		t.Error("Expected D.F3 to be a pointer to a slice of file headers, got nil")
	}

	if s.D.F4 != nil {
		fmt.Print(s.D.F4)
		t.Error("Expected D.F4 to be nil, got a file header")
	}

	if s.G == nil {
		t.Error("Expected G to be a pointer to a slice of file headers, got nil")
	}

	if len(s.D.F2) != 2 {
		t.Errorf("Expected D.F2 to have 2 file headers, got %d", len(s.D.F2))
	}

	if len(*s.D.F3) != 3 {
		t.Errorf("Expected D.F3 to have 3 file headers, got %d", len(*s.D.F3))
	}

	if len(*s.G) != 2 {
		t.Errorf("Expected G to have 2 file headers, got %d", len(*s.G))
	}

	if s.D.F.Filename != "test.txt" {
		t.Errorf("Expected D.F.Filename to be 'test.txt', got %s", s.D.F.Filename)
	}

	if s.D.F2[0].Filename != "test2.txt" {
		t.Errorf("Expected D.F2[0].Filename to be 'test2.txt', got %s", s.D.F2[0].Filename)
	}

	if s.D.F2[1].Filename != "test3.txt" {
		t.Errorf("Expected D.F2[1].Filename to be 'test3.txt', got %s", s.D.F2[1].Filename)
	}

	if (*s.D.F3)[0].Filename != "test.txt" {
		t.Errorf("Expected D.F3[0].Filename to be 'test.txt', got %s", (*s.D.F3)[0].Filename)
	}

	if (*s.D.F3)[1].Filename != "test2.txt" {
		t.Errorf("Expected D.F3[1].Filename to be 'test2.txt', got %s", (*s.D.F3)[1].Filename)
	}

	if (*s.D.F3)[2].Filename != "test3.txt" {
		t.Errorf("Expected D.F3[2].Filename to be 'test3.txt', got %s", (*s.D.F3)[2].Filename)
	}

	if (*s.G)[0].Filename != "test.txt" {
		t.Errorf("Expected G[0].Filename to be 'test.txt', got %s", (*s.G)[0].Filename)
	}

	if (*s.G)[1].Filename != "test2.txt" {
		t.Errorf("Expected G[1].Filename to be 'test2.txt', got %s", (*s.G)[1].Filename)
	}

	if s.J[0].K == nil {
		t.Error("Expected J[0].K to be a pointer to a slice of file headers, got nil")
	}

	if s.J[1].K == nil {
		t.Error("Expected J[1].K to be a pointer to a slice of file headers, got nil")
	}

	if len(*s.J[0].K) != 2 {
		t.Errorf("Expected J[0].K to have 2 file headers, got %d", len(*s.J[0].K))
	}

	if len(*s.J[1].K) != 2 {
		t.Errorf("Expected J[1].K to have 2 file headers, got %d", len(*s.J[1].K))
	}

	if (*s.J[0].K)[0].Filename != "test.txt" {
		t.Errorf("Expected J[0].K[0].Filename to be 'test.txt', got %s", (*s.J[0].K)[0].Filename)
	}

	if (*s.J[0].K)[1].Filename != "test2.txt" {
		t.Errorf("Expected J[0].K[1].Filename to be 'test2.txt', got %s", (*s.J[0].K)[1].Filename)
	}

	if (*s.J[1].K)[0].Filename != "test2.txt" {
		t.Errorf("Expected J[1].K[0].Filename to be 'test2.txt', got %s", (*s.J[1].K)[0].Filename)
	}

	if (*s.J[1].K)[1].Filename != "test3.txt" {
		t.Errorf("Expected J[1].K[1].Filename to be 'test3.txt', got %s", (*s.J[1].K)[1].Filename)
	}
}

func BenchmarkDecoderMultipartFiles(b *testing.B) {
	type S struct {
		A string `schema:"a,required"`
		B int    `schema:"b,required"`
		C bool   `schema:"c,required"`
		D struct {
			E  float64                 `schema:"e,required"`
			F  *multipart.FileHeader   `schema:"f,required"`
			F2 []*multipart.FileHeader `schema:"f2,required"`
		} `schema:"d,required"`
		G *[]*multipart.FileHeader `schema:"g,required"`
	}
	s := S{}
	data := map[string][]string{
		"a":   {"abc"},
		"b":   {"123"},
		"c":   {"true"},
		"d.e": {"3.14"},
	}

	// Create dummy file headers for testing
	dummyFile := &multipart.FileHeader{
		Filename: "test.txt",
		Size:     4,
	}

	dummyFile2 := &multipart.FileHeader{
		Filename: "test2.txt",
		Size:     4,
	}

	dummyFile3 := &multipart.FileHeader{
		Filename: "test3.txt",
		Size:     4,
	}

	// Create slice for file headers
	fileHeaders := map[string][]*multipart.FileHeader{
		"d.f":  {dummyFile, dummyFile2},
		"d.f2": {dummyFile2, dummyFile3},
		"g":    {dummyFile, dummyFile2},
	}

	decoder := NewDecoder()
	var err error
	for b.Loop() {
		err = decoder.Decode(&s, data, fileHeaders)
	}

	if err != nil {
		b.Fatalf("Failed to decode: %v", err)
	}
}

func TestIsMultipartFile(t *testing.T) {
	t.Parallel()

	tc := []struct {
		typ      reflect.Type
		input    map[string][]string
		expected bool
	}{
		{
			typ:      reflect.TypeOf(string("")),
			expected: false,
		},
		{
			typ:      reflect.TypeOf([]string{}),
			expected: false,
		},
		{
			typ:      reflect.TypeOf([]*multipart.FileHeader{}),
			expected: true,
		},
		{
			typ:      reflect.TypeOf(multipart.FileHeader{}),
			expected: false,
		},
		{
			typ:      reflect.TypeOf(&multipart.FileHeader{}),
			expected: true,
		},
		{
			typ:      reflect.TypeOf([]multipart.FileHeader{}),
			expected: false,
		},
		{
			typ:      reflect.TypeOf(&[]*multipart.FileHeader{}),
			expected: true,
		},
	}

	for _, tt := range tc {
		if isMultipartField(tt.typ) != tt.expected {
			t.Errorf("Expected %v, got %v", tt.expected, isMultipartField(tt.typ))
		}
	}
}

func BenchmarkIsMultipartFile(b *testing.B) {
	cases := []struct {
		typ reflect.Type
	}{
		{
			typ: reflect.TypeOf(string("")),
		},
		{
			typ: reflect.TypeOf([]string{}),
		},
		{
			typ: reflect.TypeOf([]*multipart.FileHeader{}),
		},
		{
			typ: reflect.TypeOf(multipart.FileHeader{}),
		},
		{
			typ: reflect.TypeOf(&multipart.FileHeader{}),
		},
		{
			typ: reflect.TypeOf([]multipart.FileHeader{}),
		},
		{
			typ: reflect.TypeOf(&[]*multipart.FileHeader{}),
		},
	}

	for i, bc := range cases {
		b.Run(fmt.Sprintf("IsMultipartFile-%d", i), func(b *testing.B) {
			for b.Loop() {
				isMultipartField(bc.typ)
			}
		})
	}
}

func TestHandleMultipartField(t *testing.T) {
	t.Parallel()

	// Create dummy file headers for testing
	dummyFile := &multipart.FileHeader{
		Filename: "test.txt",
		Size:     4,
	}

	files := map[string][]*multipart.FileHeader{
		"f": {dummyFile},
	}

	type S struct {
		F  *multipart.FileHeader    `schema:"f,required"`
		F2 []*multipart.FileHeader  `schema:"f2,required"`
		F3 *[]*multipart.FileHeader `schema:"f3,required"`
		F4 string                   `schema:"f4,required"`
	}

	s := S{}
	rv := reflect.ValueOf(&s).Elem()

	ok := handleMultipartField(rv.FieldByName("F"), files["f"])
	if !ok {
		t.Error("Expected handleMultipartField to return true")
	}

	ok = handleMultipartField(rv.FieldByName("F2"), files["f"])
	if !ok {
		t.Error("Expected handleMultipartField to return true")
	}

	ok = handleMultipartField(rv.FieldByName("F3"), files["f"])
	if !ok {
		t.Error("Expected handleMultipartField to return true")
	}

	ok = handleMultipartField(rv.FieldByName("F4"), files["f"])
	if ok {
		t.Error("Expected handleMultipartField to return false")
	}

	if s.F == nil {
		t.Error("Expected F to be a file header, got nil")
	}

	if s.F2 == nil {
		t.Error("Expected F2 to be a slice of file headers, got nil")
	}

	if s.F3 == nil {
		t.Error("Expected F3 to be a pointer to a slice of file headers, got nil")
	}

	if len(s.F2) != 1 {
		t.Errorf("Expected F2 to have 1 file header, got %d", len(s.F2))
	}

	if len(*s.F3) != 1 {
		t.Errorf("Expected F3 to have 1 file header, got %d", len(*s.F3))
	}

	if s.F.Filename != "test.txt" {
		t.Errorf("Expected F.Filename to be 'test.txt', got %s", s.F.Filename)
	}

	if s.F2[0].Filename != "test.txt" {
		t.Errorf("Expected F2[0].Filename to be 'test.txt', got %s", s.F2[0].Filename)
	}

	if (*s.F3)[0].Filename != "test.txt" {
		t.Errorf("Expected F3[0].Filename to be 'test.txt', got %s", (*s.F3)[0].Filename)
	}
}

func TestDecodePanicIsCaughtAndReturnedAsError(t *testing.T) {
	type R struct {
		N1 []*struct {
			Value string
		}
	}
	// Simulate a path that uses an invalid (e.g. negative) slice index,
	// which can trigger a panic (e.g. reflect: slice index out of range).
	data := map[string][]string{
		"n1.-1.value": {"Foo"},
	}

	s := new(R)
	decoder := NewDecoder()
	err := decoder.Decode(s, data)
	if err == nil {
		t.Fatal("Expected an error when a panic occurs")
	}

	expected := "schema: panic while decoding: reflect: slice index out of range"
	if err.Error() != expected {
		t.Fatalf("Expected panic error message %q, got: %v", expected, err)
	}
}

func TestDecodeIndexExceedsParserLimit(t *testing.T) {
	type R struct {
		N1 []*struct {
			Value string
		}
	}
	data := map[string][]string{
		"n1.1001.value": {"Foo"},
	}

	s := new(R)
	decoder := NewDecoder()
	err := decoder.Decode(s, data)
	if err == nil {
		t.Fatal("Expected an error when index exceeds parser limit")
	}

	expected := MultiError{"n1.1001.value": errIndexTooLarge}
	if !reflect.DeepEqual(err, expected) {
		t.Fatalf("Expected %v, got: %v", expected, err)
	}
}

func BenchmarkHandleMultipartField(b *testing.B) {
	// Create dummy file headers for testing
	dummyFile := &multipart.FileHeader{
		Filename: "test.txt",
		Size:     4,
	}

	files := map[string][]*multipart.FileHeader{
		"f": {dummyFile},
	}

	type S struct {
		F  *multipart.FileHeader    `schema:"f,required"`
		F2 []*multipart.FileHeader  `schema:"f2,required"`
		F3 *[]*multipart.FileHeader `schema:"f3,required"`
		F4 string                   `schema:"f4,required"`
	}

	s := S{}
	rv := reflect.ValueOf(&s).Elem()

	f := rv.FieldByName("F")
	f2 := rv.FieldByName("F2")
	f3 := rv.FieldByName("F3")
	f4 := rv.FieldByName("F4")

	for b.Loop() {
		handleMultipartField(f, files["f"])
		handleMultipartField(f2, files["f"])
		handleMultipartField(f3, files["f"])
		handleMultipartField(f4, files["f"])
	}
}

func BenchmarkLargeStructDecode(b *testing.B) {
	data := map[string][]string{
		"f1":    {"Lorem"},
		"f2":    {"Ipsum"},
		"f3":    {"123"},
		"f4":    {"456"},
		"f5":    {"A", "B", "C", "D"},
		"f6":    {"10", "20", "30", "40"},
		"f7":    {"3.14159"},
		"f8":    {"true"},
		"f9.n2": {"NestedStringValue"},
	}

	decoder := NewDecoder()
	s := &LargeStructForBenchmark{}
	for b.Loop() {
		_ = decoder.Decode(s, data)
	}
}

func BenchmarkLargeStructDecodeParallel(b *testing.B) {
	data := map[string][]string{
		"f1":    {"Lorem"},
		"f2":    {"Ipsum"},
		"f3":    {"123"},
		"f4":    {"456"},
		"f5":    {"A", "B", "C", "D"},
		"f6":    {"10", "20", "30", "40"},
		"f7":    {"3.14159"},
		"f8":    {"true"},
		"f9.n2": {"NestedStringValue"},
	}

	decoder := NewDecoder()
	s := &LargeStructForBenchmark{}
	b.ResetTimer()

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			_ = decoder.Decode(s, data)
		}
	})
}

func BenchmarkSimpleStructDecode(b *testing.B) {
	type S struct {
		A string  `schema:"a"`
		B int     `schema:"b"`
		C bool    `schema:"c"`
		D float64 `schema:"d"`
		E struct {
			F float64 `schema:"f"`
		} `schema:"e"`
	}
	s := S{}
	data := map[string][]string{
		"a":   {"abc"},
		"b":   {"123"},
		"c":   {"true"},
		"d":   {"3.14"},
		"e.f": {"3.14"},
	}
	decoder := NewDecoder()
	for b.Loop() {
		_ = decoder.Decode(&s, data)
	}
}

func BenchmarkCheckRequiredFields(b *testing.B) {
	type S struct {
		A string `schema:"a,required"`
		B int    `schema:"b,required"`
		C bool   `schema:"c,required"`
		D struct {
			E float64 `schema:"e,required"`
		} `schema:"d,required"`
	}
	s := S{}
	data := map[string][]string{
		"a":   {"abc"},
		"b":   {"123"},
		"c":   {"true"},
		"d.e": {"3.14"},
	}
	decoder := NewDecoder()
	v := reflect.ValueOf(s)
	// v = v.Elem()
	t := v.Type()

	for b.Loop() {
		_ = decoder.checkRequired(t, data)
	}
}

func BenchmarkTimeDurationDecoding(b *testing.B) {
	type DurationStruct struct {
		Timeout time.Duration `schema:"timeout"`
	}

	// Sample input for decoding
	input := map[string][]string{
		"timeout": {"2s"},
	}

	decoder := NewDecoder()
	decoder.RegisterConverter(time.Duration(0), func(s string) reflect.Value {
		d, _ := time.ParseDuration(s)
		return reflect.ValueOf(d)
	})

	var ds DurationStruct
	for b.Loop() {
		_ = decoder.Decode(&ds, input)
	}
}

func TestConversionErrorError(t *testing.T) {
	t.Parallel()
	e := ConversionError{Key: "f", Index: -1}
	if got := e.Error(); got != "schema: error converting value for \"f\"" {
		t.Errorf("unexpected message %q", got)
	}
	e = ConversionError{Key: "f", Index: 2, Err: errors.New("boom")}
	msg := e.Error()
	if !strings.Contains(msg, "index 2 of \"f\"") || !strings.Contains(msg, "boom") {
		t.Errorf("unexpected message %q", msg)
	}
}

type sliceValue []byte

func (sliceValue) UnmarshalText([]byte) error { return nil }

type valueUM string

func (valueUM) UnmarshalText([]byte) error { return nil }

type ptrUM string

func (*ptrUM) UnmarshalText([]byte) error { return nil }

type elemUM struct{}

func (*elemUM) UnmarshalText([]byte) error { return nil }

func TestIsTextUnmarshaler(t *testing.T) {
	t.Parallel()
	cases := []struct {
		name  string
		val   interface{}
		check func(t *testing.T, u unmarshaler)
	}{
		{"value", valueUM(""), func(t *testing.T, u unmarshaler) {
			if !u.IsValid || u.IsPtr {
				t.Fatalf("wrong flags: %+v", u)
			}
		}},
		{"ptr", ptrUM(""), func(t *testing.T, u unmarshaler) {
			if !u.IsValid || !u.IsPtr {
				t.Fatalf("wrong flags: %+v", u)
			}
		}},
		{"sliceValue", sliceValue{}, func(t *testing.T, u unmarshaler) {
			if !u.IsValid {
				t.Fatalf("not valid")
			}
		}},
		{"sliceElemPtr", []*elemUM{}, func(t *testing.T, u unmarshaler) {
			if !u.IsValid || !u.IsSliceElement || !u.IsSliceElementPtr {
				t.Fatalf("wrong flags: %+v", u)
			}
		}},
		{"invalid", 42, func(t *testing.T, u unmarshaler) {
			if u.IsValid {
				t.Fatalf("expected invalid")
			}
		}},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			t.Parallel()
			c.check(t, isTextUnmarshaler(reflect.ValueOf(c.val)))
		})
	}
}

func TestHandleMultipartFieldAdditional(t *testing.T) {
	t.Parallel()
	fh1 := &multipart.FileHeader{Filename: "f1"}
	fh2 := &multipart.FileHeader{Filename: "f2"}

	var a *multipart.FileHeader
	if !handleMultipartField(reflect.ValueOf(&a).Elem(), []*multipart.FileHeader{fh1}) || a != fh1 {
		t.Errorf("single header not set")
	}

	var b []*multipart.FileHeader
	if !handleMultipartField(reflect.ValueOf(&b).Elem(), []*multipart.FileHeader{fh1, fh2}) || len(b) != 2 || b[1] != fh2 {
		t.Errorf("slice headers not set")
	}

	var c *[]*multipart.FileHeader
	if !handleMultipartField(reflect.ValueOf(&c).Elem(), []*multipart.FileHeader{fh1}) || c == nil || len(*c) != 1 || (*c)[0] != fh1 {
		t.Errorf("pointer slice not set")
	}

	var d *multipart.FileHeader
	if !handleMultipartField(reflect.ValueOf(&d).Elem(), nil) || d != nil {
		t.Errorf("empty files not handled")
	}

	x := 0
	if handleMultipartField(reflect.ValueOf(&x).Elem(), []*multipart.FileHeader{fh1}) {
		t.Errorf("non multipart field handled")
	}
}

type unsupported struct {
	C complex64 `schema:"c"`
}

type textErr struct{}

func (*textErr) UnmarshalText([]byte) error { return errors.New("bad") }

type withSlice struct {
	A []struct {
		B int `schema:"b"`
	} `schema:"a"`
}

type withText struct {
	T textErr `schema:"t"`
}

type valueErrUM string

func (valueErrUM) UnmarshalText([]byte) error { return errors.New("bad") }

type sliceUM struct{}

func (*sliceUM) UnmarshalText([]byte) error { return errors.New("bad") }

type panicType int

func TestDecodeErrors(t *testing.T) {
	t.Parallel()
	t.Run("invalid pointer", func(t *testing.T) {
		t.Parallel()
		var s unsupported
		if err := NewDecoder().Decode(s, nil); err == nil {
			t.Fatalf("expected error")
		}
	})

	t.Run("panic converter", func(t *testing.T) {
		t.Parallel()
		dec := NewDecoder()
		dec.RegisterConverter(panicType(0), func(string) reflect.Value { panic("boom") })
		var target struct {
			P panicType `schema:"p"`
		}
		if err := dec.Decode(&target, map[string][]string{"p": {"x"}}); err == nil {
			t.Fatalf("expected panic error")
		}
	})

	t.Run("panic error converter", func(t *testing.T) {
		t.Parallel()
		dec := NewDecoder()
		dec.RegisterConverter(panicType(0), func(string) reflect.Value { panic(errors.New("x")) })
		var target struct {
			P panicType `schema:"p"`
		}
		if err := dec.Decode(&target, map[string][]string{"p": {"x"}}); err == nil {
			t.Fatalf("expected panic error")
		}
	})

	t.Run("unsupported type", func(t *testing.T) {
		t.Parallel()
		var u unsupported
		if err := NewDecoder().Decode(&u, map[string][]string{"c": {"1"}}); err == nil {
			t.Fatalf("expected error")
		}
	})

	t.Run("text unmarshaler error", func(t *testing.T) {
		t.Parallel()
		var w withText
		err := NewDecoder().Decode(&w, map[string][]string{"t": {"x"}})
		if err == nil {
			t.Fatalf("expected error")
		}
		if _, ok := err.(MultiError)["t"].(ConversionError); !ok {
			t.Fatalf("wrong error type: %v", err)
		}
	})

	t.Run("index larger", func(t *testing.T) {
		t.Parallel()
		dec := NewDecoder()
		dec.MaxSize(0)
		var s withSlice
		err := dec.Decode(&s, map[string][]string{"a.1.b": {"5"}})
		if err == nil || !strings.Contains(err.(MultiError)["a.1.b"].Error(), "maxSize") {
			t.Fatalf("unexpected error: %v", err)
		}
	})

	t.Run("slice converter missing", func(t *testing.T) {
		t.Parallel()
		var s struct {
			C []complex64 `schema:"c"`
		}
		if err := NewDecoder().Decode(&s, map[string][]string{"c": {"1"}}); err == nil {
			t.Fatalf("expected error")
		}
	})

	t.Run("slice textunmarshal error", func(t *testing.T) {
		t.Parallel()
		var s struct {
			S []sliceUM `schema:"s"`
		}
		if err := NewDecoder().Decode(&s, map[string][]string{"s": {"a"}}); err == nil {
			t.Fatalf("expected error")
		}
	})

	t.Run("value unmarshal error", func(t *testing.T) {
		t.Parallel()
		var s struct {
			V valueErrUM `schema:"v"`
		}
		if err := NewDecoder().Decode(&s, map[string][]string{"v": {"a"}}); err == nil {
			t.Fatalf("expected error")
		}
	})
}

func TestDecodeMultipartFiles(t *testing.T) {
	type payload struct {
		Single   *multipart.FileHeader    `schema:"single"`
		Multiple []*multipart.FileHeader  `schema:"multi"`
		PtrSlice *[]*multipart.FileHeader `schema:"ptr"`
	}

	fh1 := &multipart.FileHeader{Filename: "a"}
	fh2 := &multipart.FileHeader{Filename: "b"}

	src := map[string][]string{}
	files := map[string][]*multipart.FileHeader{
		"single": {fh1},
		"multi":  {fh1, fh2},
		"ptr":    {fh2},
	}

	var p payload
	if err := NewDecoder().Decode(&p, src, files); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if p.Single != fh1 {
		t.Fatalf("single not set")
	}
	if len(p.Multiple) != 2 || p.Multiple[0] != fh1 || p.Multiple[1] != fh2 {
		t.Fatalf("multi not set")
	}
	if p.PtrSlice == nil || len(*p.PtrSlice) != 1 || (*p.PtrSlice)[0] != fh2 {
		t.Fatalf("ptr slice not set")
	}
}

func TestDecodeSliceTextUnmarshalerError(t *testing.T) {
	type target struct {
		B []rudeBool `schema:"b"`
	}

	var s target
	if err := NewDecoder().Decode(&s, map[string][]string{"b": {"maybe"}}); err == nil {
		t.Fatalf("expected error")
	}
}

func TestDecodeCommaSeparatedZeroEmpty(t *testing.T) {
	type target struct {
		N []int `schema:"n"`
	}
	dec := NewDecoder()
	dec.ZeroEmpty(true)
	var s target
	if err := dec.Decode(&s, map[string][]string{"n": {"1,,2"}}); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !reflect.DeepEqual(s.N, []int{1, 0, 2}) {
		t.Fatalf("unexpected slice: %v", s.N)
	}
}

func TestDecodeCommaSeparatedPointerSlice(t *testing.T) {
	type target struct {
		N []*int `schema:"n"`
	}
	var s target
	if err := NewDecoder().Decode(&s, map[string][]string{"n": {"1,2"}}); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(s.N) != 2 || *s.N[0] != 1 || *s.N[1] != 2 {
		t.Fatalf("unexpected values: %v %v", s.N[0], s.N[1])
	}
}

func TestDecodeCommaSeparatedAliasSliceError(t *testing.T) {
	type target struct {
		A []IntAlias `schema:"a"`
	}

	var s target
	if err := NewDecoder().Decode(&s, map[string][]string{"a": {"1,a"}}); err == nil {
		t.Fatalf("expected error")
	}
}
