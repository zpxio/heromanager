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

package data

import (
	"log"
	"math"
)

type ValueCollection struct {
	values    map[string]float32
	validator ValueValidator
}

type ValueValidator interface {
	KeyIsValid(name string) bool
	Keys() map[string]bool
	MaxValue() float64
	MinValue() float64
}

func (vc *ValueCollection) Value(name string) float32 {
	val, exists := vc.values[name]

	if exists {
		return val
	} else {
		return 0.0
	}
}

func NewCollection(validator ValueValidator, initialValue float32) ValueCollection {
	coll := ValueCollection{validator: validator, values: make(map[string]float32)}

	coll.SetAll(initialValue)

	return coll
}

func (vc *ValueCollection) Set(name string, value float32) {

	if vc.validator.KeyIsValid(name) {

		value = float32(math.Max(float64(value), vc.validator.MinValue()))
		value = float32(math.Min(float64(value), vc.validator.MaxValue()))

		vc.values[name] = value
	} else {
		log.Printf("Attempt to set invalid value: %s", name)
	}
}

func (vc *ValueCollection) SetAll(value float32) {

	for key := range vc.validator.Keys() {
		vc.Set(key, value)
	}
}

func (vc *ValueCollection) hasValue(name string) bool {
	return vc.validator.KeyIsValid(name)
}
