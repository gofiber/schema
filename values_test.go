package schema

import (
	"errors"
	"math/rand/v2"
	"reflect"
	"slices"
	"strconv"
	"strings"
	"testing"
	"time"
)

// valuesTarget exercises every way a key reaches a field: scalars and their
// pointers, slices built from repeated keys and from commas, nested and
// embedded structs, slices of structs addressed by index, TextUnmarshaler
// and custom-converter types, required keys and defaults.
type valuesTarget struct {
	Embedded
	Time    time.Time      `schema:"time"`
	Nested  valuesNested   `schema:"nested"`
	PNested *valuesNested  `schema:"pnested"`
	Custom  valuesCustom   `schema:"custom"`
	PStr    *string        `schema:"pstr"`
	Rude    rudeBool       `schema:"rude"`
	Str     string         `schema:"str"`
	Strs    []string       `schema:"strs"`
	Ints    []int          `schema:"ints"`
	Items   []valuesNested `schema:"items"`
	Named   []IntAlias     `schema:"named"`
	Float   float64        `schema:"float"`
	Int     int            `schema:"int"`
	Def     int            `schema:"def,default:42"`
	Req     string         `schema:"req,required"`
	Bool    bool           `schema:"bool"`
}

type Embedded struct {
	Emb string `schema:"emb"`
}

// valuesCustom decodes through a registered converter.
type valuesCustom struct {
	S string
}

type valuesNested struct {
	A    string `schema:"a"`
	B    []int  `schema:"b"`
	NReq string `schema:"nreq,required"`
}

// decodeBoth decodes pairs through DecodeValues and through Decode, into
// fresh values of the same type, and returns both results.
func decodeBoth[T any](t *testing.T, d *Decoder, keys, values []string) (viaValues, viaMap T, errValues, errMap error) { //nolint:nonamedreturns,gocritic // the four results read best named
	t.Helper()
	m := make(map[string][]string, len(keys))
	for i, key := range keys {
		m[key] = append(m[key], values[i])
	}
	errValues = d.DecodeValues(&viaValues, keys, values)
	errMap = d.Decode(&viaMap, m)
	return viaValues, viaMap, errValues, errMap
}

// requireSameErrors requires two decode errors to report the same keys with
// the same messages.
func requireSameErrors(t *testing.T, want, got error, context string) {
	t.Helper()
	if (want == nil) != (got == nil) {
		t.Fatalf("%s: Decode returned %v, DecodeValues %v", context, want, got)
	}
	if want == nil {
		return
	}
	var wantMulti, gotMulti MultiError
	if !errors.As(want, &wantMulti) || !errors.As(got, &gotMulti) {
		if want.Error() != got.Error() {
			t.Fatalf("%s: Decode returned %q, DecodeValues %q", context, want, got)
		}
		return
	}
	if len(wantMulti) != len(gotMulti) {
		t.Fatalf("%s: Decode reported %v, DecodeValues %v", context, wantMulti, gotMulti)
	}
	for key, err := range wantMulti {
		if gotErr, ok := gotMulti[key]; !ok || gotErr.Error() != err.Error() {
			t.Fatalf("%s: key %q: Decode reported %v, DecodeValues %v", context, key, err, gotMulti[key])
		}
	}
}

// TestDecodeValuesMatchesDecode is the differential guard for DecodeValues:
// across generated pairs, repeated keys adjacent and apart, empty values,
// unknown and malformed keys, and every decoder option, it must decode what
// Decode decodes from the map the same pairs make, and report the same errors.
func TestDecodeValuesMatchesDecode(t *testing.T) {
	t.Parallel()

	keyPool := []string{
		"str", "Str", "pstr", "strs", "ints", "int", "float", "bool", "rude",
		"time", "custom", "named", "emb", "def", "req", "nested.a", "nested.b",
		"nested.nreq", "pnested.a", "pnested.nreq", "items.0.a", "items.1.b",
		"items.3.nreq", "unknown", "nested.unknown", "items.x.a", "items.1001.a", "",
		"nested", "pnested", ".", "a..b",
	}
	valuePool := []string{
		"", "x", "1", "2", "-3", "4.5", "true", "on", "no", "1,2", "a,b,", "nope",
		"2006-01-02T15:04:05Z", "99999999999999999999",
	}

	rng := rand.New(rand.NewPCG(7, 11)) //nolint:gosec // deterministic test input, not security
	for _, option := range []func(*Decoder){
		func(*Decoder) {},
		func(d *Decoder) { d.IgnoreUnknownKeys(true) },
		func(d *Decoder) { d.ZeroEmpty(true) },
		func(d *Decoder) { d.IgnoreUnknownKeys(true); d.ZeroEmpty(true) },
	} {
		d := NewDecoder()
		option(d)
		d.RegisterConverter(valuesCustom{}, func(s string) reflect.Value {
			if s == "nope" {
				return reflect.Value{}
			}
			return reflect.ValueOf(valuesCustom{S: s})
		})
		for round := range 3000 {
			n := rng.IntN(12)
			keys := make([]string, 0, n)
			values := make([]string, 0, n)
			for range n {
				key := keyPool[rng.IntN(len(keyPool))]
				// The keys that differ from "str" only in case name one field
				// between them, which Decode resolves in map order: keep one.
				if key == "Str" && slices.Contains(keys, "str") || key == "str" && slices.Contains(keys, "Str") {
					continue
				}
				// An empty value under a struct's own key zeroes the whole
				// struct with ZeroEmpty, wiping the fields its dotted keys
				// set before it or not, by map order: give it a value.
				value := func() string {
					for {
						v := valuePool[rng.IntN(len(valuePool))]
						if v != "" || (key != "nested" && key != "pnested") {
							return v
						}
					}
				}
				keys = append(keys, key)
				values = append(values, value())
				// Now and then repeat the key, next to it or later on.
				if rng.IntN(4) == 0 {
					keys = append(keys, key)
					values = append(values, value())
				}
			}

			viaValues, viaMap, errValues, errMap := decodeBoth[valuesTarget](t, d, keys, values)
			context := "round " + strconv.Itoa(round) + " pairs " + pairsString(keys, values)
			requireSameErrors(t, errMap, errValues, context)
			if !reflect.DeepEqual(viaMap, viaValues) {
				t.Fatalf("%s: Decode decoded %+v, DecodeValues %+v", context, viaMap, viaValues)
			}
		}
	}
}

func pairsString(keys, values []string) string {
	var b strings.Builder
	for i := range keys {
		if i > 0 {
			b.WriteByte('&')
		}
		b.WriteString(keys[i] + "=" + values[i])
	}
	return b.String()
}

// TestDecodeValuesGroupsKeys pins how pairs of one key become its values:
// in the order given, whether the pairs are adjacent or not, past the pairs
// grouped on the stack, and with the key found again when required keys and
// defaults look it up.
func TestDecodeValuesGroupsKeys(t *testing.T) {
	t.Parallel()

	type target struct {
		A   []string `schema:"a"`
		B   []int    `schema:"b"`
		C   string   `schema:"c"`
		Req []string `schema:"req,required"`
		Def string   `schema:"def,default:d"`
	}
	d := NewDecoder()

	var got target
	err := d.DecodeValues(&got,
		[]string{"a", "b", "a", "a", "c", "b", "c", "req", "req"},
		[]string{"1", "2", "3", "4", "first", "5", "last", "", "r"})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	// Without ZeroEmpty an empty value adds no element.
	want := target{A: []string{"1", "3", "4"}, B: []int{2, 5}, C: "last", Req: []string{"r"}, Def: "d"}
	if !reflect.DeepEqual(want, got) {
		t.Fatalf("got %+v, want %+v", got, want)
	}

	// More pairs than the stack holds, so the grouping is allocated.
	var keys, values []string
	for i := range maxInlinePairs * 3 {
		keys = append(keys, "ab"[i%2:i%2+1])
		values = append(values, strconv.Itoa(i))
	}
	keys = append(keys, "req")
	values = append(values, "r")
	got = target{}
	if err = d.DecodeValues(&got, keys, values); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(got.A) != maxInlinePairs*3/2 || got.A[0] != "0" || got.A[len(got.A)-1] != strconv.Itoa(maxInlinePairs*3-2) {
		t.Fatalf("grouped values out of order: %v", got.A)
	}
	if len(got.B) != maxInlinePairs*3/2 || got.B[0] != 1 {
		t.Fatalf("grouped values out of order: %v", got.B)
	}

	// A required key given only empty values is still missing, and a key
	// given no pairs at all takes its default.
	got = target{}
	err = d.DecodeValues(&got, []string{"c"}, []string{"x"})
	var multi MultiError
	if !errors.As(err, &multi) || multi["req"] == nil {
		t.Fatalf("expected the missing required key to be reported, got %v", err)
	}
	if got.Def != "d" {
		t.Fatalf("expected the default, got %q", got.Def)
	}
}

func TestDecodeValuesRejectsMismatchedInput(t *testing.T) {
	t.Parallel()

	d := NewDecoder()
	var s struct {
		A string `schema:"a"`
	}
	if err := d.DecodeValues(&s, []string{"a"}, nil); !errors.Is(err, errValuesLength) {
		t.Fatalf("expected errValuesLength, got %v", err)
	}
	if err := d.DecodeValues(s, []string{"a"}, []string{"x"}); !errors.Is(err, errNotPointerToStruct) {
		t.Fatalf("expected errNotPointerToStruct, got %v", err)
	}
	if err := d.DecodeValues(&s, nil, nil); err != nil {
		t.Fatalf("expected no error for no pairs, got %v", err)
	}
}

// TestDecodeValuesDoesNotAllocate pins that grouping the pairs costs no
// allocation: a struct of scalars decodes without any.
func TestDecodeValuesDoesNotAllocate(t *testing.T) { //nolint:paralleltest // testing.AllocsPerRun panics in a parallel test
	type target struct {
		A string  `schema:"a"`
		B int     `schema:"b"`
		C bool    `schema:"c"`
		D float64 `schema:"d"`
	}
	d := NewDecoder()
	keys := []string{"a", "b", "c", "d", "b", "unknown"}
	values := []string{"x", "1", "true", "2.5", "2", "y"}
	d.IgnoreUnknownKeys(true)
	var s target
	if err := d.DecodeValues(&s, keys, values); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if allocs := testing.AllocsPerRun(100, func() {
		_ = d.DecodeValues(&s, keys, values) //nolint:errcheck // checked above
	}); allocs != 0 {
		t.Fatalf("DecodeValues allocated %v times per run", allocs)
	}
}

func BenchmarkDecodeValues(b *testing.B) {
	type target struct {
		Name  string   `schema:"name"`
		Hobby []string `schema:"hobby"`
		ID    int      `schema:"id"`
	}
	keys := []string{"id", "name", "hobby", "hobby"}
	values := []string{"1", "tom", "basketball", "football"}
	m := map[string][]string{"id": {"1"}, "name": {"tom"}, "hobby": {"basketball", "football"}}
	d := NewDecoder()
	var s target

	b.Run("values", func(b *testing.B) {
		b.ReportAllocs()
		for b.Loop() {
			_ = d.DecodeValues(&s, keys, values) //nolint:errcheck // benchmark
		}
	})
	b.Run("map", func(b *testing.B) {
		b.ReportAllocs()
		for b.Loop() {
			_ = d.Decode(&s, m) //nolint:errcheck // benchmark
		}
	})
}

// TestDecodeStackStateDoesNotAllocate pins that what a decode keeps from key
// to key, the required-key bitset and the slice growth tracker, stays on the
// stack for both Decode and DecodeValues: a struct with required keys and an
// indexed slice that needs no growing decodes without allocating.
func TestDecodeStackStateDoesNotAllocate(t *testing.T) { //nolint:paralleltest // testing.AllocsPerRun panics in a parallel test
	type item struct {
		A string `schema:"a"`
	}
	type target struct {
		Req   string `schema:"req,required"`
		Items []item `schema:"items"`
	}
	d := NewDecoder()
	keys := []string{"req", "items.0.a", "items.1.a"}
	values := []string{"x", "y", "z"}
	m := map[string][]string{"req": {"x"}, "items.0.a": {"y"}, "items.1.a": {"z"}}
	s := target{Items: make([]item, 2)}

	if err := d.DecodeValues(&s, keys, values); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if err := d.Decode(&s, m); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if s.Req != "x" || s.Items[0].A != "y" || s.Items[1].A != "z" {
		t.Fatalf("unexpected result %+v", s)
	}
	if allocs := testing.AllocsPerRun(100, func() {
		_ = d.DecodeValues(&s, keys, values) //nolint:errcheck // checked above
	}); allocs != 0 {
		t.Fatalf("DecodeValues allocated %v times per run", allocs)
	}
	if allocs := testing.AllocsPerRun(100, func() {
		_ = d.Decode(&s, m) //nolint:errcheck // checked above
	}); allocs != 0 {
		t.Fatalf("Decode allocated %v times per run", allocs)
	}
}
