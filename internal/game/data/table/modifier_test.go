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
	"encoding/json"
	"github.com/stretchr/testify/suite"
	"testing"
)

type AdjustmentTestSuite struct {
	suite.Suite

	keys   []string
	policy *Policy
}

func (s AdjustmentTestSuite) TestNewModifier() {
	m := NewModifier(s.policy)

	for _, k := range s.keys {
		s.Zero(m.adjustments[k])
	}

	s.Len(m.adjustments, len(s.keys))
}

func (s AdjustmentTestSuite) TestSet() {
	m := NewModifier(s.policy)

	testKey := s.keys[1]
	s.Zero(m.adjustments[testKey])

	m.set(testKey, 0.5)
	s.Equal(0.5, m.adjustments[testKey])
}

func (s AdjustmentTestSuite) TestFactor() {
	m := NewModifier(s.policy)

	testZero := s.keys[0]
	s.Zero(m.adjustments[testZero])
	s.Equal(1.0, m.Factor(testZero))

	testKey := s.keys[1]
	s.Zero(m.adjustments[testKey])

	m.set(testKey, 0.5)
	s.Equal(1.5, m.Factor(testKey))
}

func (s AdjustmentTestSuite) TestFactor_Invalid() {
	m := NewModifier(s.policy)

	testMissing := "MISSING"
	s.Require().NotContains(s.keys, testMissing)
	s.Zero(m.adjustments[testMissing])
	s.Equal(1.0, m.Factor(testMissing))
}

func (s AdjustmentTestSuite) TestLoad() {
	m := NewModifier(s.policy)

	data := map[string]float64{
		s.keys[0]: 0.1,
		s.keys[1]: 0.2,
		s.keys[2]: 0.3,
	}

	// Ensure everything starts at zero
	for _, k := range s.keys {
		s.Zero(m.adjustments[k])
	}

	// Set some initial values
	m.set(s.keys[0], 0.7)
	m.set(s.keys[3], 0.8)

	// Load values over the existing set
	m.Load(data)

	// Check the values
	s.Equal(0.1, m.adjustments[s.keys[0]])
	s.Equal(0.2, m.adjustments[s.keys[1]])
	s.Equal(0.3, m.adjustments[s.keys[2]])
	s.Equal(0.8, m.adjustments[s.keys[3]])
}

func (s AdjustmentTestSuite) TestAdd() {
	m := NewModifier(s.policy)

	// Ensure everything starts at zero
	for _, k := range s.keys {
		s.Zero(m.adjustments[k])
	}

	// Set initial values
	m.set(s.keys[0], 0.5)
	m.set(s.keys[1], 0.5)

	// Create a second modifier
	m2 := NewModifier(s.policy)
	m2.set(s.keys[0], 0.5)

	// Add the second modifier to the existing
	m.Add(m2)

	// Check the values
	s.Equal(1.0, m.adjustments[s.keys[0]])
	s.Equal(0.5, m.adjustments[s.keys[1]])
}

func (s AdjustmentTestSuite) TestAdd_Negative() {
	m := NewModifier(s.policy)

	// Ensure everything starts at zero
	for _, k := range s.keys {
		s.Zero(m.adjustments[k])
	}

	// Set initial values
	m.set(s.keys[0], 0.5)
	m.set(s.keys[1], 0.5)

	// Create a second modifier
	m2 := NewModifier(s.policy)
	m2.set(s.keys[0], -0.5)

	// Add the second modifier to the existing
	m.Add(m2)

	// Check the values
	s.Equal(0.0, m.adjustments[s.keys[0]])
	s.Equal(0.5, m.adjustments[s.keys[1]])
}

func (s AdjustmentTestSuite) TestApply() {
	m := NewModifier(s.policy)

	// Ensure everything starts at zero
	for _, k := range s.keys {
		s.Zero(m.adjustments[k])
	}

	// Set initial values
	m.set(s.keys[0], 0.5)

	// Pick a test value
	testValue := 15.0

	// Check the applied modifications
	s.InDelta(testValue*1.5, m.Apply(s.keys[0], testValue), 0.0001)
	s.Equal(testValue, m.Apply(s.keys[1], testValue))
}

func (s *AdjustmentTestSuite) TestMarshalJSON_Simple() {
	v := NewModifier(testPolicy)

	jsonData := []byte(`{ "A": 0.3, "B": 1.1 }`)

	err := json.Unmarshal(jsonData, &v)

	s.Require().Nil(err)
	s.Equal(1.3, v.Factor("A"))
	s.Equal(2.1, v.Factor("B"))
	s.Equal(1.0, v.Factor("C"))
}

func (s *AdjustmentTestSuite) TestMarshalJSON_BadForm() {
	v := NewModifier(testPolicy)

	jsonData := []byte(`{ "A": "foo", "B": 1.1, "BLAH": 6.0}`)

	err := json.Unmarshal(jsonData, &v)

	s.Require().Nil(err)
	s.Equal(1.0, v.Factor("A"))
	s.Equal(2.1, v.Factor("B"))
	s.Equal(1.0, v.Factor("C"))
	_, ok := v.adjustments["BLAH"]
	s.False(ok)
}

func TestAdjustmentSuite(t *testing.T) {
	s := new(AdjustmentTestSuite)
	s.keys = []string{"A", "B", "C", "D"}
	s.policy = NewPolicy(0, 200, 0, s.keys)
	suite.Run(t, s)
}
