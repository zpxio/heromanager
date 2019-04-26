//------------------------------------------------------------------------------
//    Copyright 2019 Jeff Sharpe (zeropointx.io)
//
//    Licensed under the Apache License, Version 2.0 (the "License");
//    you may not use this file except in compliance with the License.
//    You may obtain a copy of the License at
//
//        http://www.apache.org/licenses/LICENSE-2.0
//
//    Unless required by applicable law or agreed to in writing, software
//    distributed under the License is distributed on an "AS IS" BASIS,
//    WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//    See the License for the specific language governing permissions and
//    limitations under the License.
//------------------------------------------------------------------------------

package util

import (
	"math/rand"
)

type Weighted interface {
	Rarity() float32
}

func Pick(options []interface{}, randSource ...float32) int {
	weights := make([]float32, len(options), len(options))
	for i, item := range options {
		var weight float32
		switch item.(type) {
		case Weighted:
			weight = item.(Weighted).Rarity()
		default:
			weight = float32(1)
		}

		weights[i] = weight
	}

	return PickWeightedIndex(weights, randSource...)
}

func PickWeightedIndex(weights []float32, randSource ...float32) int {
	weightBoundaries := make([]float32, len(weights), len(weights))

	// Fill in weight boundaries
	cumulativeWeight := float32(0.0)
	for i, weight := range weights {
		cumulativeWeight += weight
		weightBoundaries[i] = cumulativeWeight
	}

	// Pull a random value @ [0.0, 1.0)
	var rv float32
	if len(randSource) > 0 {
		rv = randSource[0]
	} else {
		rv = rand.Float32()
	}

	// Set the weighted target value
	wtv := cumulativeWeight * rv

	// Find the weighted target value
	for ti, upperBound := range weightBoundaries {
		if upperBound > wtv {
			return ti
		}
	}

	// Otherwise, pick the last option
	return len(weights) - 1
}
