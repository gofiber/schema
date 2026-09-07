package schema

import (
	"math"
	"math/rand/v2"
	"strconv"
	"testing"
)

// formatFloatFixed must be indistinguishable from the strconv call it
// replaces, so every case here checks against it.
func checkFloatFixed(t *testing.T, f float64, bitSize int) {
	t.Helper()
	got := formatFloatFixed(f, bitSize)
	want := strconv.FormatFloat(f, 'f', encFloatPrec, bitSize)
	if got != want {
		t.Fatalf("formatFloatFixed(%v, %d) = %q, want %q", f, bitSize, got, want)
	}
}

// The formatter hard-codes the scale and buffer sizes that go with the
// precision; pin them to encFloatPrec.
func TestFormatFloatFixedConstants(t *testing.T) {
	t.Parallel()

	scale := 1
	for i := 0; i < encFloatPrec; i++ {
		scale *= 10
	}
	if scale != scaleFixed {
		t.Fatalf("scaleFixed = %d, want 10^%d = %d", scaleFixed, encFloatPrec, scale)
	}
	// The fraction is zero-padded by formatting scaleFixed+frac and dropping
	// the leading '1', so the scale must be exactly one digit longer.
	if got := len(strconv.Itoa(scaleFixed)); got != encFloatPrec+1 {
		t.Fatalf("scaleFixed has %d digits, want %d", got, encFloatPrec+1)
	}
	// The widest fast-path output must fit the stack buffer the formatter uses.
	widest := len(formatFloatFixed(-fastFixedLimit, 64))
	if widest > 24 {
		t.Fatalf("widest fast-path output is %d bytes, buffer holds 24", widest)
	}
}

func TestFormatFloatFixedEdgeCases(t *testing.T) {
	values := []float64{
		0, math.Copysign(0, -1),
		1, -1, 0.5, -0.5, 3.14, 3.14159, 6.28, -6.28,
		0.1, 0.2, 0.3, 1.0 / 3.0, 2.0 / 3.0,
		1e-7, 1e-6, 1e-9, 0.0000005, -0.0000005, 0.0000015, 0.0000025,
		0.1234565, 0.1234575, 999999.9999994, 999999.9999995,
		123456789.987654321, 1e12, 1e13, 1e15, 1e16,
		4503599627370495.5, 4503599627370496,
		math.SmallestNonzeroFloat64, -math.SmallestNonzeroFloat64,
		math.MaxFloat64, -math.MaxFloat64,
		math.MaxFloat32, -math.MaxFloat32,
		math.Inf(1), math.Inf(-1), math.NaN(),
	}
	for _, f := range values {
		checkFloatFixed(t, f, 64)
		checkFloatFixed(t, f, 32)
	}
}

// A float's exact expansion terminates at its last significand bit, so every
// tie at six decimals is a dyadic rational m/2^k with small k — the only
// inputs where round-half-to-even can disagree with round-half-up.
func TestFormatFloatFixedTies(t *testing.T) {
	for k := 1; k <= 12; k++ {
		unit := math.Ldexp(1, -k)
		for m := int64(-200_000); m <= 200_000; m++ {
			checkFloatFixed(t, float64(m)*unit, 64)
		}
	}
}

// One ULP either side of a x.xxxxxx5 boundary is where a rounding bug shows
// up without being an exact tie.
func TestFormatFloatFixedBoundaries(t *testing.T) {
	for m := int64(0); m < 200_000; m++ {
		f := float64(m)/1e6 + 5e-7
		for _, v := range [...]float64{
			math.Nextafter(f, math.Inf(-1)), f, math.Nextafter(f, math.Inf(1)),
		} {
			checkFloatFixed(t, v, 64)
			checkFloatFixed(t, -v, 64)
		}
	}
}

func TestFormatFloatFixedRandom(t *testing.T) {
	r := rand.New(rand.NewPCG(0x5eed, 0xf10a7))
	for i := 0; i < 200_000; i++ {
		var f float64
		switch i % 5 {
		case 0:
			f = math.Float64frombits(r.Uint64())
		case 1:
			f = (r.Float64() - 0.5) * math.Pow(10, float64(r.IntN(26)-13))
		case 2:
			f = float64(r.Int64N(2_000_000_000)) / 1e6
		case 3:
			f = float64(r.Int64N(4_000_000)-2_000_000) / 2e6
		case 4:
			f = float64(r.Int64N(1_000_000_000_000)) + r.Float64()
		}
		if math.IsNaN(f) {
			continue
		}
		checkFloatFixed(t, f, 64)
		if a := math.Abs(f); a == 0 || (a <= math.MaxFloat32 && a > 1e-40) {
			checkFloatFixed(t, f, 32)
		}
	}
}

// Sweeping float32 bit patterns covers subnormals and the whole exponent
// range the 32-bit encoder path can see.
func TestFormatFloatFixedFloat32Sweep(t *testing.T) {
	for u := uint64(0); u < 1<<32; u += 4127 {
		f := float64(math.Float32frombits(uint32(u)))
		if math.IsNaN(f) || math.IsInf(f, 0) {
			continue
		}
		checkFloatFixed(t, f, 32)
		checkFloatFixed(t, f, 64)
	}
}

func FuzzFormatFloatFixed(f *testing.F) {
	f.Add(uint64(0), false)
	f.Add(math.Float64bits(3.14159), false)
	f.Add(math.Float64bits(0.0000005), true)
	f.Fuzz(func(t *testing.T, u uint64, as32 bool) {
		v := math.Float64frombits(u)
		bitSize := 64
		if as32 {
			bitSize = 32
		}
		got := formatFloatFixed(v, bitSize)
		want := strconv.FormatFloat(v, 'f', encFloatPrec, bitSize)
		if got != want {
			t.Fatalf("formatFloatFixed(%v, %d) = %q, want %q", v, bitSize, got, want)
		}
	})
}

func BenchmarkFormatFloatFixed(b *testing.B) {
	b.ReportAllocs()
	var s string
	for b.Loop() {
		s = formatFloatFixed(3.14159, 64)
	}
	_ = s
}
