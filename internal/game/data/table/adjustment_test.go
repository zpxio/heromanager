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

package table

import (
	"github.com/stretchr/testify/suite"
	"testing"
)

type AdjustmentTestSuite struct {
	suite.Suite
}

func (s AdjustmentTestSuite) TestZero() {
	a := ZeroAdjustment()

	s.Require().Equal(a.Factor, 0.0)
}

func (s AdjustmentTestSuite) TestParse_Normal() {
	s.Equal(0.5, ParseAdjustment("0.5").Factor)
	s.Equal(0.0, ParseAdjustment("0").Factor)
	s.Equal(1.0, ParseAdjustment("1.0").Factor)
	s.Equal(-0.2, ParseAdjustment("-0.2").Factor)
}

func (s AdjustmentTestSuite) TestParse_Fail() {
	s.Equal(0.0, ParseAdjustment("0.5r").Factor)
	s.Equal(0.0, ParseAdjustment("anemone").Factor)
	s.Equal(0.0, ParseAdjustment("").Factor)
}

func (s AdjustmentTestSuite) TestString() {
	examples := []string{"0.5000", "-2.1000", "0.0000"}

	for _, example := range examples {
		a := ParseAdjustment(example)
		s.Equal(example, a.String())
	}
}

func (s AdjustmentTestSuite) TestCombine_Simple() {
	aFactor := 0.5
	a := Adjustment{Factor: aFactor}

	bFactor := 0.2
	b := Adjustment{Factor: bFactor}

	s.Equal(aFactor+bFactor, a.Combine(b).Factor)
}

func (s AdjustmentTestSuite) TestCombine_MixedSign() {
	aFactor := 0.5
	a := Adjustment{Factor: aFactor}

	bFactor := -0.6
	b := Adjustment{Factor: bFactor}

	s.Equal(aFactor+bFactor, a.Combine(b).Factor)
}

func (s AdjustmentTestSuite) TestApply_Simple() {
	factor := 0.5
	value := 6.2
	a := Adjustment{Factor: factor}

	s.InDelta(value*1.5, a.ApplyTo(value), 0.0001)
}

func (s AdjustmentTestSuite) TestApply_Zero() {
	factor := 0.0
	value := 8.1
	a := Adjustment{Factor: factor}

	s.InDelta(value, a.ApplyTo(value), 0.0001)
}

func (s AdjustmentTestSuite) TestApply_Negative() {
	factor := -0.8
	value := 12.42
	a := Adjustment{Factor: factor}

	s.InDelta(value*0.2, a.ApplyTo(value), 0.0001)
}

func TestAdjustmentSuite(t *testing.T) {
	suite.Run(t, new(AdjustmentTestSuite))
}
