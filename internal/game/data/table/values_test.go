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

import (
	"github.com/stretchr/testify/suite"
	"log"
	"testing"
)

type ValuesTestSuite struct {
	suite.Suite
}

var testPolicy *Policy

func init() {
	testMin := 4.0
	testMax := 10.0
	testDefault := 5.0
	testKeys := []string{"A", "B", "C", "D"}
	testPolicy = NewPolicy(testMin, testMax, testDefault, testKeys)
}

func (t *ValuesTestSuite) TestNewValues() {
	v := NewValues(testPolicy)

	// All values are default
	for _, k := range testPolicy.ValidKeys() {
		t.Equal(testPolicy.DefaultValue(), v.Get(k))
	}

	// Value count equals key count
	t.Len(v.values, len(testPolicy.ValidKeys()))
}

func (t *ValuesTestSuite) TestLoad_Partial() {
	v := NewValues(testPolicy)

	// All values are default
	for _, k := range testPolicy.ValidKeys() {
		t.Equal(testPolicy.DefaultValue(), v.Get(k))
	}

	// Load a map
	testData := map[string]float64{"A": 5.5, "C": 6.7}
	v.Load(testData)

	// Ensure changes
	for k, j := range testData {
		t.Equal(j, v.values[k])
	}

	// Ensure values not loaded are still default
	for _, k := range testPolicy.ValidKeys() {
		if _, ok := testData[k]; !ok {
			t.Equal(testPolicy.DefaultValue(), v.Get(k))
		}
	}
}

func (t *ValuesTestSuite) TestLoad_Abnormal() {
	v := NewValues(testPolicy)
	log.Printf("Table Created")

	// All values are default
	for _, k := range testPolicy.ValidKeys() {
		t.Equal(testPolicy.DefaultValue(), v.Get(k))
	}
	log.Printf("Defaults checked")

	// Load a map
	testData := map[string]float64{
		"A": testPolicy.MinValue() - 1.2,
		"B": testPolicy.DefaultValue(),
		"C": testPolicy.MaxValue() + 4.4,
		"E": testPolicy.MinValue() + 1.2,
	}
	log.Printf("Test data created")
	v.Load(testData)
	log.Printf("Map data loaded.")

	// Ensure low clamping
	t.Equal(testPolicy.MinValue(), v.Get("A"), "Failed to apply low clamp")
	// Ensure no clamping
	t.Equal(testPolicy.DefaultValue(), v.Get("B"), "Failed to set default value")
	// Ensure high clamping
	t.Equal(testPolicy.MaxValue(), v.Get("C"), "Failed to apply high clamp")
	// Ensure no change
	t.Equal(testPolicy.DefaultValue(), v.Get("D"), "Modified ignored value")
	// Ensure invalid key was not created
	_, vFound := v.values["E"]
	t.False(vFound)
}

func (t *ValuesTestSuite) TestGetSet() {
	v := NewValues(testPolicy)

	// All values are default
	for _, k := range testPolicy.ValidKeys() {
		t.Require().Equal(testPolicy.DefaultValue(), v.Get(k))
	}

	// Set a value
	valA := 6.6
	v.Set("A", valA)
	t.Equal(valA, v.Get("A"))
}

func (t *ValuesTestSuite) TestSet_Invalid() {
	v := NewValues(testPolicy)

	// All values are default
	for _, k := range testPolicy.ValidKeys() {
		t.Require().Equal(testPolicy.DefaultValue(), v.Get(k))
	}
	defaultLen := len(v.values)

	// Set a value for an invalid key
	valA := 6.6
	v.Set("E", valA)

	// All values are default
	for _, k := range testPolicy.ValidKeys() {
		t.Require().Equal(testPolicy.DefaultValue(), v.Get(k))
	}
	t.Require().Len(v.values, defaultLen)
}

func (t *ValuesTestSuite) TestGet_Invalid() {
	v := NewValues(testPolicy)

	// Set all values to non-default
	newVal := 7.7
	for _, k := range testPolicy.ValidKeys() {
		v.Set(k, newVal)
	}

	// All values are set
	for _, k := range testPolicy.ValidKeys() {
		t.Require().Equal(newVal, v.Get(k))
	}

	// Get a value for an invalid key
	invalidVal := v.Get("E")

	t.Require().Equal(testPolicy.DefaultValue(), invalidVal)
}

func (t *ValuesTestSuite) TestCopy() {
	v := NewValues(testPolicy)

	// All values are default
	for _, k := range testPolicy.ValidKeys() {
		t.Equal(testPolicy.DefaultValue(), v.Get(k))
	}

	// Load a map
	testData := map[string]float64{
		"A": testPolicy.MinValue() + 0.8,
		"B": testPolicy.MinValue() + 1.2,
		"C": testPolicy.MinValue() + 2.1,
		"D": testPolicy.MinValue() + 3.6,
	}
	v.Load(testData)

	vCopy := v.Copy()

	// Lengths are equal
	t.Require().Len(vCopy.values, len(v.values))

	// Contents are equal
	for _, k := range testPolicy.ValidKeys() {
		t.Equal(v.Get(k), vCopy.Get(k))
	}
}

func TestValuesSuite(t *testing.T) {
	suite.Run(t, new(ValuesTestSuite))
}
