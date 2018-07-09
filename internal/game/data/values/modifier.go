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

package values

import (
	"log"
	"math"
)

type Modifier struct {
	values    map[string]float32
	validator ValueValidator
}

const DefaultFactor = float32(1.0)

func CreateModifier(validator ValueValidator) Modifier {
	attrValues := make(map[string]float32)

	attrs := Modifier{values: attrValues, validator: validator}

	return attrs
}

func (attr *Modifier) Factor(name string) float32 {
	val, exists := attr.values[name]

	if exists {
		return val
	} else {
		return DefaultFactor
	}
}

func (attr *Modifier) Set(name string, value float64) {

	if attr.validator.KeyIsValid(name) {

		value = math.Max(value, 0.0)

		attr.values[name] = float32(value)
	} else {
		log.Printf("Ignoring modifier set attempt of unrecognized attribute ID: %s", name)
	}
}
