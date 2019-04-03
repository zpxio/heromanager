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

package data

import (
	"github.com/stretchr/testify/suite"
	"github.com/zpxio/heromanager/internal/game/data/caste"
	"github.com/zpxio/heromanager/internal/game/data/profession"
	"github.com/zpxio/heromanager/internal/game/data/race"
	"testing"
)

type ManifestTestSuite struct {
	suite.Suite
}

func TestManifestSuite(t *testing.T) {
	suite.Run(t, new(ManifestTestSuite))
}

func (s *ManifestTestSuite) TestNewManifest() {
	m := NewManifest()

	s.Empty(m.races)
	s.Empty(m.castes)
	s.Empty(m.professions)
	s.Empty(m.raceKeys)
	s.Empty(m.casteKeys)
	s.Empty(m.professionKeys)
}

func (s *ManifestTestSuite) TestRaceUsage() {
	m := NewManifest()

	s.Empty(m.races)
	testRace1 := race.Blank()
	testRace1.Name = "TEST1"
	m.RegisterRace(testRace1.Name, testRace1)
	s.Len(m.races, 1)

	// Add a non-colliding key
	testRace2 := race.Blank()
	testRace2.Name = "TEST2"
	m.RegisterRace(testRace2.Name, testRace2)
	s.Len(m.races, 2)

	// Add a colliding key
	testRace3 := race.Blank()
	testRace3.Name = "TEST3"
	m.RegisterRace(testRace2.Name, testRace3)
	s.Len(m.races, 2)

	// Check the keys
	s.Len(m.AllRaces(), 2)

	// Resolve a key
	resolve1, found1 := m.ResolveRace(testRace1.Name)
	s.True(found1)
	s.Equal(testRace1.Name, resolve1.Name)

	// Check that resolving the colliding key returns the second value
	resolve2, found2 := m.ResolveRace(testRace2.Name)
	s.True(found2)
	s.Equal(testRace3.Name, resolve2.Name)

	// Try to resolve a key that doesn't exist
	resolve3, found3 := m.ResolveRace(testRace1.Name + "-Not-Exist")
	s.False(found3)
	s.Nil(resolve3)
}

func (s *ManifestTestSuite) TestCasteUsage() {
	m := NewManifest()

	s.Empty(m.castes)
	testCaste1 := caste.Blank()
	testCaste1.Name = "TEST1"
	m.RegisterCaste(testCaste1.Name, testCaste1)
	s.Len(m.castes, 1)

	// Add a non-colliding key
	testCaste2 := caste.Blank()
	testCaste2.Name = "TEST2"
	m.RegisterCaste(testCaste2.Name, testCaste2)
	s.Len(m.castes, 2)

	// Add a colliding key
	testCaste3 := caste.Blank()
	testCaste3.Name = "TEST3"
	m.RegisterCaste(testCaste2.Name, testCaste3)
	s.Len(m.castes, 2)

	// Check the keys
	s.Len(m.AllCastes(), 2)

	// Resolve a key
	resolve1, found1 := m.ResolveCaste(testCaste1.Name)
	s.True(found1)
	s.Equal(testCaste1.Name, resolve1.Name)

	// Check that resolving the colliding key returns the second value
	resolve2, found2 := m.ResolveCaste(testCaste2.Name)
	s.True(found2)
	s.Equal(testCaste3.Name, resolve2.Name)

	// Try to resolve a key that doesn't exist
	resolve3, found3 := m.ResolveCaste(testCaste1.Name + "-Not-Exist")
	s.False(found3)
	s.Nil(resolve3)
}

func (s *ManifestTestSuite) TestProfessionUsage() {
	m := NewManifest()

	s.Empty(m.professions)
	testJob1 := profession.Blank()
	testJob1.Name = "TEST1"
	m.RegisterProfession(testJob1.Name, testJob1)
	s.Len(m.professions, 1)

	// Add a non-colliding key
	testJob2 := profession.Blank()
	testJob2.Name = "TEST2"
	m.RegisterProfession(testJob2.Name, testJob2)
	s.Len(m.professions, 2)

	// Add a colliding key
	testJob3 := profession.Blank()
	testJob3.Name = "TEST3"
	m.RegisterProfession(testJob2.Name, testJob3)
	s.Len(m.professions, 2)

	// Check the keys
	s.Len(m.AllProfessions(), 2)

	// Resolve a key
	resolve1, found1 := m.ResolveProfession(testJob1.Name)
	s.True(found1)
	s.Equal(testJob1.Name, resolve1.Name)

	// Check that resolving the colliding key returns the second value
	resolve2, found2 := m.ResolveProfession(testJob2.Name)
	s.True(found2)
	s.Equal(testJob3.Name, resolve2.Name)

	// Try to resolve a key that doesn't exist
	resolve3, found3 := m.ResolveProfession(testJob1.Name + "-Not-Exist")
	s.False(found3)
	s.Nil(resolve3)
}
