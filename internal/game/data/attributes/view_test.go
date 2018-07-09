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

func TestCreateView(t *testing.T) {

	b := Create(7.0)
	v := b.CreateView()

	assert.Equal(t, b.ValueCollection, v.Base, "View is not using requested base.")

	for _, id := range Ids {
		assert.Equal(t, float32(7.0), v.Value(id), "Empty view modified value for %s", id)
	}
}

func TestView_Modify(t *testing.T) {
	b := Create(8.0)
	v := b.CreateView()

	m1 := CreateModifier()

	m1.Set(Finesse, 1.50)
	m1.Set(Allure, 0.5)

	v.Modify("A", &m1)

	assert.Equal(t, float32(b.Value(Finesse)*1.5), v.Value(Finesse), "Modification doesn't seem to be applied.")
	assert.Equal(t, float32(b.Value(Allure)*0.5), v.Value(Allure), "Modification doesn't seem to be applied.")

	m2 := CreateModifier()
	m2.Set(Brawn, 2.00)
	m2.Set(Allure, 2.00)

	v.Modify("B", &m2)

	assert.Equal(t, float32(b.Value(Finesse)*1.5), v.Value(Finesse), "Secondary modification has modified previous value.")
	assert.Equal(t, float32(b.Value(Brawn)*2.0), v.Value(Brawn), "Secondary modification doesn't seem to be applied.")
	assert.Equal(t, float32(b.Value(Allure)*0.5*2.0), v.Value(Allure), "Secondary modification doesn't seem to be compounding.")
}

func TestView_ModifyReplace(t *testing.T) {
	b := Create(8.0)
	v := b.CreateView()

	m1 := CreateModifier()

	m1.Set(Finesse, 1.50)
	m1.Set(Allure, 0.5)

	v.Modify("A", &m1)

	assert.Equal(t, float32(b.Value(Finesse)*1.5), v.Value(Finesse), "Modification doesn't seem to be applied.")
	assert.Equal(t, float32(b.Value(Allure)*0.5), v.Value(Allure), "Modification doesn't seem to be applied.")

	m2 := CreateModifier()
	m2.Set(Brawn, 2.00)
	m2.Set(Allure, 2.00)

	v.Modify("A", &m2)

	assert.Equal(t, float32(b.Value(Finesse)), v.Value(Finesse), "Replacement modification did not clear previous modification.")
	assert.Equal(t, float32(b.Value(Brawn)*2.0), v.Value(Brawn), "Replacement modification did not replace previous modifications")
	assert.Equal(t, float32(b.Value(Allure)*2.0), v.Value(Allure), "Replacement modification did not get applied.")
}
