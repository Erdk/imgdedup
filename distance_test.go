package main

import (
	"fmt"
	"testing"
)

var result int

// Test helper, loads two histograms for testing purposes

func testHelper() (*histrec, *histrec, error) {
	histogram1, err := histogram("test/img1.jpg") // [15][0][0] = 16384, rest is 0
	if err != nil {
		return nil, nil, fmt.Errorf("Cannot load file: %s", err.Error())
	}

	histogram2, err := histogram("test/img2.jpg") // [0][15][0] = 16384, rest is 0
	if err != nil {
		return nil, nil, fmt.Errorf("Cannot load file: %s", err.Error())
	}

	return histogram1, histogram2, err
}

// Manhattan

func TestManhattan(t *testing.T) {
	histogram1, histogram2, err := testHelper()
	if err != nil {
		t.Error(err)
	}

	result := manhattanDistance(histogram1.h, histogram2.h) // 32768
	if result != 32768 {
		t.Errorf("%s: should be %d, instead of %d\n", t.Name(), 32786, result)
	}
}

func TestManhattan2(t *testing.T) {
	histogram1, histogram2, err := testHelper()
	if err != nil {
		t.Error(err)
	}

	result := manhattanDistance2(histogram1.h, histogram2.h) // 32768
	if result != 32768 {
		t.Errorf("%s: should be %d, instead of %d\n", t.Name(), 32786, result)
	}
}

func TestManhattan3(t *testing.T) {
	histogram1, histogram2, err := testHelper()
	if err != nil {
		t.Error(err)
	}

	result := manhattanDistance3(histogram1.h, histogram2.h) // 32768
	if result != 32768 {
		t.Errorf("%s: should be %d, instead of %d\n", t.Name(), 32786, result)
	}
}

func TestManhattan4(t *testing.T) {
	histogram1, histogram2, err := testHelper()
	if err != nil {
		t.Error(err)
	}

	result := manhattanDistance4(histogram1.h, histogram2.h) // 32768
	if result != 32768 {
		t.Errorf("%s: should be %d, instead of %d\n", t.Name(), 32786, result)
	}
}

func BenchmarkManhattan(b *testing.B) {
	histogram1, histogram2, err := testHelper()
	if err != nil {
		b.Error(err)
	}

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		result = manhattanDistance(histogram1.h, histogram2.h)
	}
}

func BenchmarkManhattan2(b *testing.B) {
	histogram1, histogram2, err := testHelper()
	if err != nil {
		b.Error(err)
	}

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		result = manhattanDistance2(histogram1.h, histogram2.h)
	}
}

func BenchmarkManhattan3(b *testing.B) {
	histogram1, histogram2, err := testHelper()
	if err != nil {
		b.Error(err)
	}

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		result = manhattanDistance3(histogram1.h, histogram2.h)
	}
}

func BenchmarkManhattan4(b *testing.B) {
	histogram1, histogram2, err := testHelper()
	if err != nil {
		b.Error(err)
	}

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		result = manhattanDistance4(histogram1.h, histogram2.h)
	}
}

// Chi Square

func TestChiSquare(t *testing.T) {
	histogram1, histogram2, err := testHelper()
	if err != nil {
		t.Error(err)
	}

	result := chisquareDistance(histogram1.h, histogram2.h) // 32768 / 2 = 16384
	if result != 16384 {
		t.Errorf("%s: should be %d, instead of %d\n", t.Name(), 16384, result)
	}
}

func TestChiSquare2(t *testing.T) {
	histogram1, histogram2, err := testHelper()
	if err != nil {
		t.Error(err)
	}

	result := chisquareDistance2(histogram1.h, histogram2.h) // 32768 / 2 = 16384
	if result != 16384 {
		t.Errorf("%s: should be %d, instead of %d\n", t.Name(), 16384, result)
	}
}

func BenchmarkChiSquare(b *testing.B) {
	histogram1, histogram2, err := testHelper()
	if err != nil {
		b.Error(err)
	}

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		result = chisquareDistance(histogram1.h, histogram2.h)
	}
}

func BenchmarkChiSquare2(b *testing.B) {
	histogram1, histogram2, err := testHelper()
	if err != nil {
		b.Error(err)
	}

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		result = chisquareDistance2(histogram1.h, histogram2.h)
	}
}
