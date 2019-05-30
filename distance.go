package main

import (
	"log"
	"math"
)

/* distance
 * Prototype for function computing distance between two histograms
 */
type distance func(h1, h2 *hist) uint

func getDistanceFunction(distanceType string) distance {
	switch distanceType {
	case "manhattan":
		return manhattanDistance
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

func manhattanDistance(h1, h2 *hist) uint {
	var sum uint
	for i := 0; i < dimm; i++ {
		for j := 0; j < dimm; j++ {
			for k := 0; k < dimm; k++ {
				sum += uint(math.Abs(float64(h1[i][j][k]) - float64(h2[i][j][k])))
			}
		}
	}
	return sum
}

func chisquareDistance(h1, h2 *hist) uint {
	var sum uint

	for i := 0; i < dimm; i++ {
		for j := 0; j < dimm; j++ {
			for k := 0; k < dimm; k++ {
				if h1[i][j][k]+h2[i][j][k] > 0 {
					sum += uint(math.Pow(float64(h1[i][j][k]-h2[i][j][k]), 2.0) / float64(h1[i][j][k]+h2[i][j][k]))
				}
			}
		}
	}

	return sum / 2
}

func intersectionDistance(h1, h2 *hist) uint {
	panic("intersection not implemented")
	return 0
}

func correlationDistance(h1, h2 *hist) uint {
	panic("correlation not implemented")
	return 0
}

func bhattacharyyaDistance(h1, h2 *hist) uint {
	panic("bhattacharyya not implemented")
	return 0
}
