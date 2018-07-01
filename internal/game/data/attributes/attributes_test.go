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

import "testing"

func TestAttributesCreation(t *testing.T) {
	var base float32 = 3.1337
	attrs := Create(base)

	for id := range Keys {
		if attrs.Value(id) != base {
			t.Errorf("Unexpected base value. Expected %.3f, Saw %.3f", base, attrs.Value(id))
		}
	}
}

func TestAttributesManipulation(t *testing.T) {
	attrs := Create(0.0)

	if attrs.Value("BRN") != 0.0 {
		t.Errorf("Initial BRN value was not zero.")
	}

	var newValue float32 = 3.14156
	attrs.set("BRN", newValue)

	if attrs.Value("BRN") != newValue {
		t.Errorf("BRN value was not updated.")
	}
}

func TestAttributesClamping(t *testing.T) {
	attrs := Create(0.0)

	if attrs.Value("BRN") != 0.0 {
		t.Errorf("Initial BRN factor was not zero.")
	}

	var newValue float32 = -3.113
	attrs.set("BRN", newValue)

	if attrs.Value("BRN") != 0.0 {
		t.Errorf("BRN factor was not clamped to zero")
	}

	var hiValue = AttributeMax + 4.21
	attrs.set("BRN", hiValue)

	if attrs.Value("BRN") != AttributeMax {
		t.Errorf("BRN factor was not clamped to the maximum")
	}
}
