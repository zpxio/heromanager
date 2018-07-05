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

type AttributeValidator struct {
	min  float64
	max  float64
	keys map[string]bool
}

func keyNames(m map[string]string) map[string]bool {
	keys := make(map[string]bool)

	for k := range m {
		keys[k] = true
	}

	return keys
}

var Validator AttributeValidator

func init() {
	Validator = AttributeValidator{min: 0, max: float64(AttributeMax), keys: keyNames(Keys)}
}

func (av *AttributeValidator) KeyIsValid(name string) bool {
	_, keyExists := Keys[name]

	return keyExists
}

func (av *AttributeValidator) MaxValue() float64 {
	return av.max
}

func (av *AttributeValidator) MinValue() float64 {
	return av.min
}

func (av *AttributeValidator) Keys() map[string]bool {
	return av.keys
}
