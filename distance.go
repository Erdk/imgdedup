package main

import (
	"log"
	"math"
)

/* distance
 * Prototype for function computing distance between two histograms
 */
type distance func(h1, h2 *hist) int

func getDistanceFunction(distanceType string) distance {
	switch distanceType {
	case "manhattan":
		return manhattanDistance3
	case "chisquare":
		return chisquareDistance
	case "intersection":
		return intersectionDistance
	case "correlation":
		return correlationDistance
	case "bhattacharyya":
		return bhattacharyyaDistance
	default:
		log.Fatal("unsupported distance")
	}

	return manhattanDistance
}

// Manhattan distance variants

func manhattanDistance(h1, h2 *hist) int {
	var sum int

	for i := 0; i < dimm; i++ {
		for j := 0; j < dimm; j++ {
			for k := 0; k < dimm; k++ {
				sum += int(math.Abs(float64(h1[i][j][k]) - float64(h2[i][j][k])))
			}
		}
	}
	return sum
}

func manhattanDistance2(h1, h2 *hist) int {
	var sum int

	for i := 0; i < dimm; i++ {
		for j := 0; j < dimm; j++ {
			for k := 0; k < dimm; k++ {
				sum += int(math.Abs(float64(h1[i][j][k] - h2[i][j][k])))
			}
		}
	}
	return sum
}

func manhattanDistance3(h1, h2 *hist) int {
	var sum int
	var res int

	for i := 0; i < dimm; i++ {
		for j := 0; j < dimm; j++ {
			for k := 0; k < dimm; k++ {
				res = h1[i][j][k] - h2[i][j][k]
				if res > 0 {
					sum += res
				} else {
					sum -= res
				}
			}
		}
	}
	return sum
}

func manhattanDistance4(h1, h2 *hist) int {
	var sum int
	var res int
	var mask int

	for i := 0; i < dimm; i++ {
		for j := 0; j < dimm; j++ {
			for k := 0; k < dimm; k++ {
				res = h1[i][j][k] - h2[i][j][k]
				mask = res >> 63
				sum += res ^ mask - mask
			}
		}
	}
	return sum
}

// Chi Square

func chisquareDistance(h1, h2 *hist) int {
	var sum int

	for i := 0; i < dimm; i++ {
		for j := 0; j < dimm; j++ {
			for k := 0; k < dimm; k++ {
				if h1[i][j][k]+h2[i][j][k] > 0 {
					sum += int(math.Pow(float64(h1[i][j][k]-h2[i][j][k]), 2.0) / float64(h1[i][j][k]+h2[i][j][k]))
				}
			}
		}
	}

	return sum / 2
}

func chisquareDistance2(h1, h2 *hist) int {
	var sum int
	var p1 float64
	var p2 float64

	for i := 0; i < dimm; i++ {
		for j := 0; j < dimm; j++ {
			for k := 0; k < dimm; k++ {
				p1 = float64(h1[i][j][k] - h2[i][j][k])
				p2 = float64(h1[i][j][k] + h2[i][j][k])
				if p2 > 0.0 {
					sum += int(p1 * p1 / p2)
				}
			}
		}
	}

	return sum / 2
}

// Intersection

func intersectionDistance(h1, h2 *hist) int {
	panic("intersection not implemented")
}

// Correlation

func correlationDistance(h1, h2 *hist) int {
	panic("correlation not implemented")
}

// Bhattacharry

func bhattacharyyaDistance(h1, h2 *hist) int {
	panic("bhattacharyya not implemented")
}
