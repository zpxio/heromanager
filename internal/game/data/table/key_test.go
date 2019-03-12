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
	"testing"
)

type KeyTestSuite struct {
	suite.Suite
}

func (s *KeyTestSuite) TestCreateSet_Empty() {
	set := NewKeySet()

	s.Require().Empty(set.keys)
}

func (s *KeyTestSuite) TestCreateSet() {
	testKeys := []string{"A", "B", "C"}

	set := NewKeySet(testKeys...)

	s.Require().Len(set.keys, len(testKeys))
}

func (s *KeyTestSuite) TestAll() {
	testKeys := []string{"A", "B", "C"}

	set := NewKeySet(testKeys...)

	for _, k := range testKeys {
		s.Require().Contains(set.All(), k)
	}

	s.Require().Len(set.All(), len(testKeys))
}

func (s *KeyTestSuite) TestAll_Mutability() {
	testKeys := []string{"A", "B", "C"}
	addKey := "D"

	set := NewKeySet(testKeys...)

	s.Require().Len(set.All(), len(testKeys))
	set.Add(addKey)

	s.Require().Len(set.All(), len(testKeys)+1)
	s.Require().Contains(set.All(), addKey)
}

func (s *KeyTestSuite) TestContains_Positive() {
	testKeys := []string{"A", "B", "C"}

	set := NewKeySet(testKeys...)

	for _, k := range testKeys {
		s.Require().True(set.Contains(k))
	}
}

func (s *KeyTestSuite) TestContains_Negative() {
	testKeys := []string{"A", "B", "C"}
	missingKeys := []string{"D", ""}

	set := NewKeySet(testKeys...)

	for _, k := range missingKeys {
		s.Require().False(set.Contains(k))
	}
}

func TestKeySuite(t *testing.T) {
	suite.Run(t, new(KeyTestSuite))
}
