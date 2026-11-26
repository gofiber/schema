package schema

import (
	"errors"
	"mime/multipart"
	"reflect"
	"strings"
	"testing"
)

// splitPairs splits "a=1&b=2" into its keys and values.
func splitPairs(s string) (keys, values []string) { //nolint:nonamedreturns // the two results read best named
	for pair := range strings.SplitSeq(s, "&") {
		key, value, _ := strings.Cut(pair, "=")
		keys, values = append(keys, key), append(values, value)
	}
	return keys, values
}

// With ExactNames a tagged field matches its tags as written, joined by dots,
// with plain-digit indexes; a promoted field matches only its own name; an
// untagged field still matches its Go name in any case.
func TestExactNames(t *testing.T) {
	t.Parallel()

	type Inner struct {
		Deep string `schema:"deep"`
	}
	type Host struct {
		Value string `schema:"__Host-token"`
	}
	// Mid's Auth is shadowed by session's field tagged auth.
	type Auth struct {
		Other string `schema:"other"`
	}
	type Mid struct {
		Inner
		Auth
	}
	type hidden struct {
		Hidden string `schema:"hidden"`
	}
	type session struct {
		Meta map[string]string `schema:"meta"`
		Host
		Mid
		hidden
		Auth  Inner    `schema:"auth"`
		Token string   `schema:"access_token"`
		Name  string   // untagged
		IDs   []string `schema:"ids"`
		Items []Inner  `schema:"items"`
	}
	for _, tc := range []struct {
		name, pairs string
		want        session
	}{
		{name: "exact first", pairs: "access_token=a&Access_Token=b", want: session{Token: "a"}},
		{name: "folded first", pairs: "Access_Token=b&access_token=a", want: session{Token: "a"}},
		{name: "folded alone", pairs: "ACCESS_TOKEN=b"},
		{name: "path forms", pairs: ".access_token=b&Access_Token.0=b&access_token[0]=b"},
		{name: "promoted", pairs: "__host-token=b&__Host-token=a", want: session{Host: Host{Value: "a"}}},
		{name: "through its embedded struct", pairs: "Host=b&Host.__Host-token=b&host.__Host-token=b"},
		{name: "through an unexported embedded struct", pairs: "Hidden=b&hidden=a&HIDDEN=b", want: session{hidden: hidden{"a"}}},
		{name: "two levels down", pairs: "deep=a&Inner.deep=b&Mid.Inner.deep=b", want: session{Mid: Mid{Inner: Inner{"a"}}}},
		{name: "shadowed embedded struct", pairs: "auth.other=b&Mid.Auth.other=b&other=a", want: session{Mid: Mid{Auth: Auth{"a"}}}},
		{name: "nested", pairs: "Auth.deep=b&auth.DEEP=b&auth.deep=a", want: session{Auth: Inner{"a"}}},
		{name: "nested in a slice", pairs: "items.0.deep=a&Items.0.deep=b&items.0.DEEP=b", want: session{Items: []Inner{{"a"}}}},
		{name: "index spellings", pairs: "items.0.deep=a&items.00.deep=b&items.+0.deep=b", want: session{Items: []Inner{{"a"}}}},
		{name: "map field", pairs: "meta.k=a&META.k=b&Meta.k=b", want: session{Meta: map[string]string{"k": "a"}}},
		{name: "slice", pairs: "IDS=1&ids=2", want: session{IDs: []string{"2"}}},
		{name: "untagged field folds", pairs: "NAME=john", want: session{Name: "john"}},
	} {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			d := NewDecoder()
			d.ExactNames(true)
			d.IgnoreUnknownKeys(true)
			keys, values := splitPairs(tc.pairs)
			src := map[string][]string{}
			for i, key := range keys {
				src[key] = append(src[key], values[i])
			}
			var viaValues, viaMap session
			if err := d.DecodeValues(&viaValues, keys, values); err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if err := d.Decode(&viaMap, src); err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if !reflect.DeepEqual(viaValues, tc.want) || !reflect.DeepEqual(viaMap, tc.want) {
				t.Fatalf("DecodeValues decoded %+v, Decode %+v, want %+v", viaValues, viaMap, tc.want)
			}
		})
	}
}

// A key in another spelling names no field: it is an unknown key, fills no
// required field, and leaves a default in place.
func TestExactNamesOtherSpellingsAreUnknown(t *testing.T) {
	t.Parallel()

	type Host struct {
		Value string `schema:"__Host-token,required"`
	}
	type form struct {
		Host
		Avatar *multipart.FileHeader `schema:"avatar,required"`
		Doc    *multipart.FileHeader `schema:"doc"`
		Token  string                `schema:"token,required"`
		Lang   string                `schema:"lang,default:en"`
		Nick   string                `schema:",required"`
	}
	doc := &multipart.FileHeader{Filename: "doc"}
	files := map[string][]*multipart.FileHeader{"AVATAR": {doc}, "doc": {doc}}
	keys, values := splitPairs("host.__Host-token=x&TOKEN=x&Lang=fr&NICK=n&nope=x")
	d := NewDecoder()
	d.ExactNames(true)
	d.IgnoreUnknownKeys(true)

	var got form
	requireMissing(t, d.DecodeValues(&got, keys, values, files), "Host.__Host-token", "avatar", "token")
	if want := (form{Doc: doc, Lang: "en", Nick: "n"}); got != want {
		t.Fatalf("got %+v, want %+v", got, want)
	}

	d.IgnoreUnknownKeys(false)
	var multi MultiError
	err := d.DecodeValues(&form{}, keys, values, files)
	for _, key := range []string{"host.__Host-token", "TOKEN", "Lang", "AVATAR", "nope"} {
		if !errors.As(err, &multi) || !errors.As(multi[key], new(UnknownKeyError)) {
			t.Fatalf("expected %q unknown, got %v", key, err)
		}
	}
}

// Both modes on one type and one decoder: a path one mode resolved and cached
// must not answer the other.
func TestExactNamesSharesNoPaths(t *testing.T) {
	t.Parallel()

	type Auth struct {
		Other string `schema:"other"`
	}
	type Inner struct {
		Deep string `schema:"deep"`
	}
	type target struct {
		Auth
		Field Inner  `schema:"auth"`
		Tok   string `schema:"Token"`
	}
	keys, values := splitPairs("auth.other=o&auth.deep=d&token=t")
	folded, exact := target{Auth: Auth{"o"}, Tok: "t"}, target{Field: Inner{"d"}}
	d := NewDecoder()
	d.IgnoreUnknownKeys(true)
	for _, mode := range []bool{false, true, false, true} {
		d.ExactNames(mode)
		want := folded
		if mode {
			want = exact
		}
		for range 2 { // the second decode reads the paths the first one cached
			var got target
			if err := d.DecodeValues(&got, keys, values); err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if got != want {
				t.Fatalf("ExactNames %v: got %+v, want %+v", mode, got, want)
			}
		}
	}
}

func BenchmarkDecodeValuesExactNames(b *testing.B) {
	type session struct {
		ID    string `schema:"session_id"`
		Token string `schema:"__Host-token"`
		Theme string `schema:"theme"`
		Lang  string `schema:"lang"`
	}
	const cookies = "session_id=abc&__Host-token=t&theme=dark&lang=en&_ga=GA1.2"
	for _, bc := range []struct {
		name, query string
		exact       bool
	}{
		{"folded", cookies, false},
		{"exact", cookies, true},
		{"exact_other_spelling", cookies + "&Theme=light", true},
	} {
		keys, values := splitPairs(bc.query)
		d := NewDecoder()
		d.IgnoreUnknownKeys(true)
		d.ExactNames(bc.exact)
		var s session
		b.Run(bc.name, func(b *testing.B) {
			b.ReportAllocs()
			for b.Loop() {
				_ = d.DecodeValues(&s, keys, values) //nolint:errcheck // benchmark
			}
		})
	}
}
