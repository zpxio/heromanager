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

import "strconv"

type AdjustmentMethod int

type Adjustment struct {
	Factor float64
}

func ZeroAdjustment() Adjustment {
	return Adjustment{Factor: 0.0}
}

func ParseAdjustment(text string) Adjustment {

	factor, err := strconv.ParseFloat(text, 64)

	if err != nil {
		factor = 0.0
	}
	return Adjustment{Factor: factor}
}

func (a Adjustment) Combine(adjustment Adjustment) Adjustment {
	return Adjustment{Factor: a.Factor + adjustment.Factor}
}

func (a Adjustment) ApplyTo(value float64) float64 {
	return value * (1.0 + a.Factor)
}

func (a Adjustment) String() string {
	return strconv.FormatFloat(a.Factor, 'f', 4, 64)
}
