package main

import "testing"

var result int

func TestManhattan(t *testing.T) {
	histogram1, err := histogram("test/img1.jpg") // [15][0][0] = 16384, rest is 0
	if err != nil {
		t.Errorf("Cannot load test/img1.jpg")
	}

	histogram2, err := histogram("test/img2.jpg") // [0][15][0] = 16384, rest is 0
	if err != nil {
		t.Errorf("Cannot load test/img2.jpg")
	}

	result := manhattanDistance(histogram1.h, histogram2.h) // 32768
	if result != 32768 {
		t.Errorf("%s: should be %d, instead of %d\n", t.Name(), 32786, result)
	}
}

func TestManhattan2(t *testing.T) {
	histogram1, err := histogram("test/img1.jpg") // [15][0][0] = 16384, rest is 0
	if err != nil {
		t.Errorf("Cannot load test/img1.jpg")
	}

	histogram2, err := histogram("test/img2.jpg") // [0][15][0] = 16384, rest is 0
	if err != nil {
		t.Errorf("Cannot load test/img2.jpg")
	}

	result := manhattanDistance2(histogram1.h, histogram2.h) // 32768
	if result != 32768 {
		t.Errorf("%s: should be %d, instead of %d\n", t.Name(), 32786, result)
	}
}

func TestManhattan3(t *testing.T) {
	histogram1, err := histogram("test/img1.jpg") // [15][0][0] = 16384, rest is 0
	if err != nil {
		t.Errorf("Cannot load test/img1.jpg")
	}

	histogram2, err := histogram("test/img2.jpg") // [0][15][0] = 16384, rest is 0
	if err != nil {
		t.Errorf("Cannot load test/img2.jpg")
	}

	result := manhattanDistance3(histogram1.h, histogram2.h) // 32768
	if result != 32768 {
		t.Errorf("%s: should be %d, instead of %d\n", t.Name(), 32786, result)
	}
}

func TestManhattan4(t *testing.T) {
	histogram1, err := histogram("test/img1.jpg") // [15][0][0] = 16384, rest is 0
	if err != nil {
		t.Errorf("Cannot load test/img1.jpg")
	}

	histogram2, err := histogram("test/img2.jpg") // [0][15][0] = 16384, rest is 0
	if err != nil {
		t.Errorf("Cannot load test/img2.jpg")
	}

	result := manhattanDistance4(histogram1.h, histogram2.h) // 32768
	if result != 32768 {
		t.Errorf("%s: should be %d, instead of %d\n", t.Name(), 32786, result)
	}
}

func BenchmarkManhattan(b *testing.B) {
	histogram1, err := histogram("test/img1.jpg")
	if err != nil {
		b.Errorf("Cannot load test/img1.jpg")
	}

	histogram2, err := histogram("test/img2.jpg")
	if err != nil {
		b.Errorf("Cannot load test/img2.jpg")
	}

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		result = manhattanDistance(histogram1.h, histogram2.h)
	}
}

func BenchmarkManhattan2(b *testing.B) {
	histogram1, err := histogram("test/img1.jpg")
	if err != nil {
		b.Errorf("Cannot load test/img1.jpg")
	}

	histogram2, err := histogram("test/img2.jpg")
	if err != nil {
		b.Errorf("Cannot load test/img2.jpg")
	}

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		result = manhattanDistance2(histogram1.h, histogram2.h)
	}
}

func BenchmarkManhattan3(b *testing.B) {
	histogram1, err := histogram("test/img1.jpg")
	if err != nil {
		b.Errorf("Cannot load test/img1.jpg")
	}

	histogram2, err := histogram("test/img2.jpg")
	if err != nil {
		b.Errorf("Cannot load test/img2.jpg")
	}

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		result = manhattanDistance3(histogram1.h, histogram2.h)
	}
}

func BenchmarkManhattan4(b *testing.B) {
	histogram1, err := histogram("test/img1.jpg")
	if err != nil {
		b.Errorf("Cannot load test/img1.jpg")
	}

	histogram2, err := histogram("test/img2.jpg")
	if err != nil {
		b.Errorf("Cannot load test/img2.jpg")
	}

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		result = manhattanDistance4(histogram1.h, histogram2.h)
	}
}
