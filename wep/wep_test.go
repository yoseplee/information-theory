package wep

import (
	"math"
	"testing"
)

const EPS float64 = 0.0000000001

var c1 = Codeword{
	Idx:     1,
	Symbol:  "00",
	Encoded: "000",
	Length:  4,
}
var c2 = Codeword{
	Idx:     2,
	Symbol:  "01",
	Encoded: "101",
	Length:  4,
}
var c3 = Codeword{
	Idx:     3,
	Symbol:  "10",
	Encoded: "110",
	Length:  4,
}
var c4 = Codeword{
	Idx:     4,
	Symbol:  "11",
	Encoded: "111",
	Length:  4,
}

func TestCube(t *testing.T) {
	want := 27.0
	if got := Cube(3); got != want {
		t.Errorf("Cube(3) = %f, want %f", got, want)
	}
}

func TestCompareCode(t *testing.T) {
	toCompare := "001"
	wantMiss, wantMatch := 1, 2

	if miss, match := c1.CompareCode(toCompare); miss != wantMiss && match != wantMatch {
		t.Errorf("CompareCode(\"001\") = miss-%d, match-%d, want miss-%d, match-%d", miss, match, wantMiss, wantMatch)
	}

	toCompare = "011"
	wantMiss, wantMatch = 2, 1

	if miss, match := c1.CompareCode(toCompare); miss != wantMiss && match != wantMatch {
		t.Errorf("CompareCode(\"001\") = miss-%d, match-%d, want miss-%d, match-%d", miss, match, wantMiss, wantMatch)
	}
}

func TestCalculateLikelihoodRate(t *testing.T) {
	Y := "000"
	e := 0.1

	want := 81.0
	got := c1.CalculateLikelihoodRate(c2, Y, e)
	diff := math.Abs(want - got)
	if diff > EPS {
		t.Errorf("CalculateLikelihoodRate(c2, %s, %f) = %f, want = %f", Y, e, got, want)
	}

	got = c1.CalculateLikelihoodRate(c3, Y, e)
	diff = math.Abs(want - got)
	if diff > EPS {
		t.Errorf("CalculateLikelihoodRate(c2, %s, %f) = %f, want = %f", Y, e, got, want)
	}
}

func TestGetId(t *testing.T) {
	want := "C1"
	if got := c1.GetId(); want != got {
		t.Errorf("GetId() = %q, want = %q", got, want)
	}

	want = "C2"
	if got := c2.GetId(); want != got {
		t.Errorf("GetId() = %q, want = %q", got, want)
	}

	want = "C3"
	if got := c3.GetId(); want != got {
		t.Errorf("GetId() = %q, want = %q", got, want)
	}

	want = "C4"
	if got := c4.GetId(); want != got {
		t.Errorf("GetId() = %q, want = %q", got, want)
	}
}
