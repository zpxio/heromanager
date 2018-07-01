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

func TestCreateModifier(t *testing.T) {
	mod := CreateModifier()

	for _, id := range Ids {
		assert.Equal(t, defaultFactor, mod.Factor(id), "Expected the default factor.")
	}
}

func TestModifier_Set(t *testing.T) {
	mod := CreateModifier()

	adjust := 1.4

	mod.Set(Brawn, adjust)

	assert.Equal(t, float32(adjust), mod.Factor(Brawn), "Modification value not set: Expected %f, saw %f", adjust, mod.Factor(Brawn))
}
