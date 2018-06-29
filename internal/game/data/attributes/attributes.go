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

var Keys = make(map[string]string)

const AttributeMax float32 = 500

func init() {
	Keys["BRN"] = "Brawn"
	Keys["INS"] = "Insight"
	Keys["FIN"] = "Finesse"
	Keys["VIG"] = "Vigor"
	Keys["ALL"] = "Allure"
}

type Attributes struct {
	values map[string]float32
}

func Create(baseValue float32) Attributes {
	attrValues := make(map[string]float32)

	for key := range Keys {
		attrValues[key] = baseValue
	}

	attrs := Attributes{values: attrValues}

	return attrs
}

func (attr *Attributes) value(name string) float32 {
	val, exists := attr.values[name]

	if exists {
		return val
	} else {
		return 0.0
	}
}

func (attr *Attributes) set(name string, value float32) {

	if isValid(name) {

		value = float32(math.Max(float64(value), 0.0))
		value = float32(math.Min(float64(value), float64(AttributeMax)))

		attr.values[name] = value
	} else {

	}
}

func (attr *Attributes) modify(mod Modifier) {

	for key, factor := range mod.values {
		attr.set(key, attr.value(key)*factor)
	}
}

func (attr *Attributes) CreateView() AttributesView {
	return CreateAttributesView(attr)
}

func isValid(name string) bool {
	_, exists := Keys[name]

	return exists
}
