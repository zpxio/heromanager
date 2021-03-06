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
	"sort"
	"testing"
)

type PolicyTestSuite struct {
	suite.Suite
}

func (t *PolicyTestSuite) TestNewPolicy() {
	testMin := 4.0
	testMax := 10.0
	testDefault := 5.0
	testKeys := []string{"A", "B", "C", "D"}

	policy := NewPolicy(testMin, testMax, testDefault, testKeys)

	t.Require().Equal(testMin, policy.MinValue())
	t.Require().Equal(testMax, policy.MaxValue())
	t.Require().Equal(testDefault, policy.DefaultValue())

	sortedResult := policy.keys.All()
	sort.Strings(sortedResult)
	t.Require().EqualValues(sortedResult, testKeys)
}

func (t *PolicyTestSuite) TestClamp() {
	testMin := 4.0
	testMax := 10.0
	testDefault := 5.0
	testKeys := []string{"A", "B", "C", "D"}

	policy := NewPolicy(testMin, testMax, testDefault, testKeys)

	// No clamping necessary
	t.Require().Equal(8.0, policy.Clamp(8.0))
	// Clamp high
	t.Require().Equal(testMax, policy.Clamp(testMax+2.0))
	// Clamp low
	t.Require().Equal(testMin, policy.Clamp(testMin-2.0))
}

func (t *PolicyTestSuite) TestValidKey_Positive() {
	testMin := 4.0
	testMax := 10.0
	testDefault := 5.0
	testKeys := []string{"A", "B", "C", "D"}

	policy := NewPolicy(testMin, testMax, testDefault, testKeys)

	for _, k := range testKeys {
		t.Require().True(policy.ValidKey(k))
	}
}

func (t *PolicyTestSuite) TestValidKey_Negative() {
	testMin := 4.0
	testMax := 10.0
	testDefault := 5.0
	testKeys := []string{"A", "B", "C", "D"}
	missingKeys := []string{"E", "F", "G", "H"}

	policy := NewPolicy(testMin, testMax, testDefault, testKeys)

	for _, k := range missingKeys {
		t.Require().False(policy.ValidKey(k))
	}
}

func (t *PolicyTestSuite) TestValidKeys() {
	testMin := 4.0
	testMax := 10.0
	testDefault := 5.0
	testKeys := []string{"A", "B", "C", "D"}

	policy := NewPolicy(testMin, testMax, testDefault, testKeys)

	sortedResult := policy.ValidKeys()
	sort.Strings(sortedResult)
	t.Require().EqualValues(sortedResult, testKeys)
}

func TestPolicySuite(t *testing.T) {
	suite.Run(t, new(PolicyTestSuite))
}
