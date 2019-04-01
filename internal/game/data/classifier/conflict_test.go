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

package classifier

import (
	"encoding/json"
	"github.com/stretchr/testify/suite"
	"github.com/zpxio/heromanager/internal/game/util"
	"gopkg.in/yaml.v2"
	"testing"
)

type ConflictTestSuite struct {
	suite.Suite
}

func TestConflictSuite(t *testing.T) {
	suite.Run(t, new(ConflictTestSuite))
}

func (s *ConflictTestSuite) TestEmptyConflicts() {
	c := EmptyConflicts()

	s.Empty(c.professions)
	s.Empty(c.castes)
	s.Empty(c.races)
}

func (s *ConflictTestSuite) TestAdd_Simple() {
	c := EmptyConflicts()

	c.add("races", "TEST1")
	s.Len(c.races, 1)
	s.Contains(c.races, "TEST1")
}

func (s *ConflictTestSuite) TestAdd_Invalid() {
	c := EmptyConflicts()

	s.Empty(c.professions)
	s.Empty(c.castes)
	s.Empty(c.races)

	c.add("Invalid", "TEST4")

	s.Empty(c.professions)
	s.Empty(c.castes)
	s.Empty(c.races)
}

func (s *ConflictTestSuite) TestAdd_Idempotence() {
	c := EmptyConflicts()

	c.add("races", "TEST1")
	s.Len(c.races, 1)
	s.Contains(c.races, "TEST1")

	c.add("races", "TEST1")
	s.Len(c.races, 1)
	s.Contains(c.races, "TEST1")
}

func (s *ConflictTestSuite) TestAdd_Multiple() {
	c := EmptyConflicts()

	c.add("races", "TEST1")
	c.add("castes", "TEST2")
	c.add("professions", "TEST3")

	s.Len(c.races, 1)
	s.Contains(c.races, "TEST1")

	s.Len(c.castes, 1)
	s.Contains(c.castes, "TEST2")

	s.Len(c.professions, 1)
	s.Contains(c.professions, "TEST3")
}

func (s *ConflictTestSuite) TestAllowRace() {
	c := EmptyConflicts()

	testKey := "TEST1"

	s.True(c.AllowRace(testKey))
	c.add(ConflictRaces, testKey)
	s.False(c.AllowRace(testKey))
}

func (s *ConflictTestSuite) TestAllowCaste() {
	c := EmptyConflicts()

	testKey := "TEST1"

	s.True(c.AllowCaste(testKey))
	c.add(ConflictCastes, testKey)
	s.False(c.AllowCaste(testKey))
}

func (s *ConflictTestSuite) TestAllowProfession() {
	c := EmptyConflicts()

	testKey := "TEST1"

	s.True(c.AllowProfession(testKey))
	c.add(ConflictProfessions, testKey)
	s.False(c.AllowProfession(testKey))
}

func (s *ConflictTestSuite) TestUnmarshallYAML() {
	c := EmptyConflicts()

	data, dataErr := util.GameFileData("testdata/game/data/classifier", "test_conflict_simple.yml")
	s.Require().Nil(dataErr)

	err := yaml.Unmarshal(data, &c)
	s.Require().Nil(err)

	s.Len(c.races, 2)
	s.Len(c.castes, 3)
	s.Len(c.professions, 4)
}

func (s *ConflictTestSuite) TestUnmarshallJSON() {
	c := EmptyConflicts()

	data, dataErr := util.GameFileData("testdata/game/data/classifier", "test_conflict_simple.json")
	s.Require().Nil(dataErr)

	err := json.Unmarshal(data, &c)
	s.Require().Nil(err)

	s.Len(c.races, 2)
	s.Len(c.castes, 3)
	s.Len(c.professions, 4)
}

func (s *ConflictTestSuite) TestUnmarshallYAML_Malformed() {
	c := EmptyConflicts()

	data, dataErr := util.GameFileData("testdata/game/data/classifier", "test_conflict_malformed.yml")
	s.Require().Nil(dataErr)

	err := yaml.Unmarshal(data, &c)
	s.NotNil(err)
}

func (s *ConflictTestSuite) TestUnmarshallJSON_Malformed() {
	c := EmptyConflicts()

	data, dataErr := util.GameFileData("testdata/game/data/classifier", "test_conflict_malformed.json")
	s.Require().Nil(dataErr)

	err := json.Unmarshal(data, &c)
	s.NotNil(err)
}
