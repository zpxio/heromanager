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

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAttributeValidator_KeyIsValid(t *testing.T) {

	for attrKey := range Keys {
		assert.True(t, Validator.KeyIsValid(attrKey))
	}
}

func TestAttributeValidator_MinValue(t *testing.T) {
	assert.Equal(t, 0.00, Validator.min)
}

func TestAttributeValidator_MaxValue(t *testing.T) {
	assert.Equal(t, float64(AttributeMax), Validator.max)
}
