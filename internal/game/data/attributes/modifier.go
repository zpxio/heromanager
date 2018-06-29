//------------------------------------------------------------------------------
//    Copyright 2018 Jeff Sharpe (zeropointx.io)
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

package attributes

import "math"

type Modifier struct {
	values map[string]float32
}

func CreateModifier() Modifier {
	attrValues := make(map[string]float32)

	attrs := Modifier{values: attrValues}

	return attrs
}

func (attr *Modifier) factor(name string) float32 {
	val, exists := attr.values[name]

	if exists {
		return val
	} else {
		return 1.0
	}
}

func (attr *Modifier) set(name string, value float32) {

	if isValid(name) {

		value = float32(math.Max(float64(value), 0.0))

		attr.values[name] = value
	} else {

	}
}
