package schema

import (
	"errors"
	"fmt"
	"reflect"
	"strconv"
	"sync"

	utils "github.com/gofiber/utils/v2"
)

type encoderFunc func(reflect.Value) string

// Encoder encodes values from a struct into url.Values.
type Encoder struct {
	cache  *cache
	regenc map[reflect.Type]encoderFunc
	// encCache memoizes the per-struct-type encoding plan
	// (map[reflect.Type][]encField) so tags are parsed and encoder
	// functions resolved once per type instead of on every Encode call.
	encCache sync.Map
}

// encField is the precomputed encoding plan for one struct field.
type encField struct {
	name      string
	enc       encoderFunc // immediate encoder; nil for structs and slices
	elemEnc   encoderFunc // slice element encoder, when the field is a slice
	idx       int
	omitEmpty bool
	// recurseStructPtr marks pointer-to-struct fields without a custom
	// encoder: non-nil values are encoded by recursing into the element.
	recurseStructPtr bool
	isStruct         bool
}

// NewEncoder returns a new Encoder with defaults.
func NewEncoder() *Encoder {
	return &Encoder{cache: newCache(), regenc: make(map[reflect.Type]encoderFunc)}
}

// Encode encodes a struct into map[string][]string.
//
// Intended for use with url.Values.
func (e *Encoder) Encode(src interface{}, dst map[string][]string) error {
	v := reflect.ValueOf(src)

	return e.encode(v, dst)
}

// RegisterEncoder registers a converter for encoding a custom type.
func (e *Encoder) RegisterEncoder(value interface{}, encoder func(reflect.Value) string) {
	e.regenc[reflect.TypeOf(value)] = encoder
	e.encCache.Clear()
}

// SetAliasTag changes the tag used to locate custom field aliases.
// The default tag is "schema".
func (e *Encoder) SetAliasTag(tag string) {
	e.cache.tag = tag
	e.encCache.Clear()
}

// structInfo returns the cached encoding plan for struct type t, building it
// on first use.
func (e *Encoder) structInfo(t reflect.Type) []encField {
	if cached, ok := e.encCache.Load(t); ok {
		return cached.([]encField)
	}
	fields := make([]encField, 0, t.NumField())
	for i := 0; i < t.NumField(); i++ {
		sf := t.Field(i)
		name, opts := fieldAlias(sf, e.cache.tag)
		if name == "-" {
			continue
		}
		ft := sf.Type
		f := encField{
			idx:       i,
			name:      name,
			omitEmpty: opts.Contains("omitempty"),
			recurseStructPtr: ft.Kind() == reflect.Ptr &&
				ft.Elem().Kind() == reflect.Struct &&
				!e.hasCustomEncoder(ft),
			enc: typeEncoder(ft, e.regenc),
		}
		if f.enc == nil {
			if ft.Kind() == reflect.Struct {
				f.isStruct = true
			} else if ft.Kind() == reflect.Slice {
				f.elemEnc = typeEncoder(ft.Elem(), e.regenc)
			}
		}
		fields = append(fields, f)
	}
	e.encCache.Store(t, fields)
	return fields
}

// isValidStructPointer test if input value is a valid struct pointer.
func isValidStructPointer(v reflect.Value) bool {
	return v.Type().Kind() == reflect.Ptr && v.Elem().IsValid() && v.Elem().Type().Kind() == reflect.Struct
}

func isZero(v reflect.Value) bool {
	switch v.Kind() {
	case reflect.Func:
	case reflect.Map, reflect.Slice:
		return v.IsNil() || v.Len() == 0
	case reflect.Array:
		for i := 0; i < v.Len(); i++ {
			if !isZero(v.Index(i)) {
				return false
			}
		}
		return true
	case reflect.Struct:
		if v.CanInterface() {
			if iz, ok := v.Interface().(interface{ IsZero() bool }); ok {
				return iz.IsZero()
			}
		}
		for i := 0; i < v.NumField(); i++ {
			if !isZero(v.Field(i)) {
				return false
			}
		}
		return true
	}
	// Compare other types directly:
	return v.IsZero()
}

func (e *Encoder) encode(v reflect.Value, dst map[string][]string) error {
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}
	if v.Kind() != reflect.Struct {
		return errors.New("schema: interface must be a struct")
	}

	var errs MultiError

	fields := e.structInfo(v.Type())
	for i := range fields {
		f := &fields[i]
		fieldValue := v.Field(f.idx)

		// Encode struct pointer types if the field is a valid pointer and a struct.
		if f.recurseStructPtr && !fieldValue.IsNil() {
			if err := e.encode(fieldValue.Elem(), dst); err != nil {
				errs = setError(errs, fieldValue.Elem().Type().String(), err)
			}
			continue
		}

		// Encode non-slice types and custom implementations immediately.
		if f.enc != nil {
			if f.omitEmpty && isZero(fieldValue) {
				continue
			}
			dst[f.name] = append(dst[f.name], f.enc(fieldValue))
			continue
		}

		if f.isStruct {
			if err := e.encode(fieldValue, dst); err != nil {
				errs = setError(errs, fieldValue.Type().String(), err)
			}
			continue
		}

		if f.elemEnc == nil {
			errs = setError(errs, fieldValue.Type().String(), fmt.Errorf("schema: encoder not found for %v", fieldValue))
			continue
		}

		// Encode a slice.
		n := fieldValue.Len()
		if n == 0 && f.omitEmpty {
			continue
		}

		values := make([]string, n)
		for j := 0; j < n; j++ {
			values[j] = f.elemEnc(fieldValue.Index(j))
		}
		dst[f.name] = values
	}

	if len(errs) > 0 {
		return errs
	}
	return nil
}

// setError lazily allocates m and stores err under key, overwriting any
// previous entry (matching the historical encoder error semantics).
func setError(m MultiError, key string, err error) MultiError {
	if m == nil {
		m = make(MultiError)
	}
	m[key] = err
	return m
}

func (e *Encoder) hasCustomEncoder(t reflect.Type) bool {
	_, exists := e.regenc[t]
	return exists
}

func typeEncoder(t reflect.Type, reg map[reflect.Type]encoderFunc) encoderFunc {
	if f, ok := reg[t]; ok {
		return f
	}

	switch t.Kind() {
	case reflect.Bool:
		return encodeBool
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return encodeInt
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return encodeUint
	case reflect.Float32:
		return encodeFloat32
	case reflect.Float64:
		return encodeFloat64
	case reflect.Ptr:
		f := typeEncoder(t.Elem(), reg)
		return func(v reflect.Value) string {
			if v.IsNil() {
				return "null"
			}
			return f(v.Elem())
		}
	case reflect.String:
		return encodeString
	default:
		return nil
	}
}

func encodeBool(v reflect.Value) string {
	return strconv.FormatBool(v.Bool())
}

func encodeInt(v reflect.Value) string {
	return utils.FormatInt(v.Int())
}

func encodeUint(v reflect.Value) string {
	return utils.FormatUint(v.Uint())
}

func encodeFloat(v reflect.Value, bits int) string {
	return strconv.FormatFloat(v.Float(), 'f', 6, bits)
}

func encodeFloat32(v reflect.Value) string {
	return encodeFloat(v, 32)
}

func encodeFloat64(v reflect.Value) string {
	return encodeFloat(v, 64)
}

func encodeString(v reflect.Value) string {
	return v.String()
}
