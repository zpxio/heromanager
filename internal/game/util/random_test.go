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
	"github.com/stretchr/testify/suite"
	"testing"
)

type RandomTestSuite struct {
	suite.Suite
}

func TestRandomSuite(t *testing.T) {
	suite.Run(t, new(RandomTestSuite))
}

func (s *RandomTestSuite) TestPick_Simple() {
	testArray := []interface{}{
		"A",
		"B",
		"C",
	}

	s.Equal(0, Pick(testArray, 0.02))
	s.Equal(1, Pick(testArray, 0.6))
	s.Equal(2, Pick(testArray, 0.999))
}

type SimpleWeighted struct {
	value  string
	weight float32
}

func (w SimpleWeighted) Rarity() float32 {
	return w.weight
}

func (s *RandomTestSuite) TestPick_Weighted() {
	testArray := []interface{}{
		SimpleWeighted{value: "A", weight: 2.2},
		SimpleWeighted{value: "B", weight: 0.3},
		SimpleWeighted{value: "C", weight: 3.0},
	}

	s.Equal(0, Pick(testArray, 0.02))
	s.Equal(0, Pick(testArray, 2.1/5.3))
	s.Equal(1, Pick(testArray, 2.4/5.3))
	s.Equal(2, Pick(testArray, 5.2))
}

func (s *RandomTestSuite) TestPick_Mixed() {
	testArray := []interface{}{
		SimpleWeighted{value: "A", weight: 2.2},
		"B",
		SimpleWeighted{value: "C", weight: 3.0},
	}

	s.Equal(0, Pick(testArray, 0.02))
	s.Equal(0, Pick(testArray, 2.1/6.2))
	s.Equal(1, Pick(testArray, 2.4/6.2))
	s.Equal(2, Pick(testArray, 6.1))
}

func (s *RandomTestSuite) TestPick_ActualRandom() {
	testArray := []interface{}{
		SimpleWeighted{value: "A", weight: 2.2},
		"B",
		SimpleWeighted{value: "C", weight: 3.0},
	}

	for i := 0; i < 1000; i++ {
		si := Pick(testArray)

		s.True(si >= 0)
		s.True(si < len(testArray))
	}
}
