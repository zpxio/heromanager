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

package hero

import (
	"github.com/stretchr/testify/suite"
	"github.com/zpxio/heromanager/internal/game/data/classifier"
	"testing"
)

type SelectorTestSuite struct {
	suite.Suite
	manifest *classifier.ClassifierManifest
}

func SetupSuite() *SelectorTestSuite {
	s := new(SelectorTestSuite)
	s.manifest = setupManifest()

	return s
}

func conflicts(ids ...string) []string {
	return ids
}

func testRace(name string, conflictCastes []string, conflictProfessions []string) classifier.Race {
	r := classifier.BlankRace()
	r.Name = name

	for _, id := range conflictCastes {
		r.Conflicts.Add(classifier.ConflictCastes, id)
	}

	for _, id := range conflictProfessions {
		r.Conflicts.Add(classifier.ConflictProfessions, id)
	}

	return r
}

func testCaste(name string, conflictRaces []string, conflictProfessions []string) classifier.Caste {
	c := classifier.BlankCaste()
	c.Name = name

	for _, id := range conflictRaces {
		c.Conflicts.Add(classifier.ConflictRaces, id)
	}

	for _, id := range conflictProfessions {
		c.Conflicts.Add(classifier.ConflictProfessions, id)
	}

	return c
}

func testProfession(name string, conflictRaces []string, conflictCastes []string) classifier.Profession {
	p := classifier.BlankProfession()
	p.Name = name

	for _, id := range conflictRaces {
		p.Conflicts.Add(classifier.ConflictRaces, id)
	}

	for _, id := range conflictCastes {
		p.Conflicts.Add(classifier.ConflictCastes, id)
	}

	return p
}

func setupManifest() *classifier.ClassifierManifest {
	manifest := classifier.NewManifest()

	manifest.RegisterRace("Dwarf", testRace("Dwarf", conflicts("Serf", "Noble"), conflicts("Farmer", "Pirate")))
	manifest.RegisterRace("Elf", testRace("Elf", conflicts("Outcast", "Serf"), conflicts("Miner", "Mason")))
	manifest.RegisterRace("Human", testRace("Human", conflicts("Elder"), conflicts("Mystic")))

	manifest.RegisterCaste("Serf", testCaste("Serf", conflicts("Dwarf", "Elf"), conflicts()))
	manifest.RegisterCaste("Noble", testCaste("Noble", conflicts("Dwarf"), conflicts()))
	manifest.RegisterCaste("Outcast", testCaste("Outcast", conflicts("Elf"), conflicts()))
	manifest.RegisterCaste("Elder", testCaste("Elder", conflicts("Human"), conflicts("Pirate")))

	manifest.RegisterProfession("Farmer", testProfession("Farmer", conflicts("Dwarf"), conflicts()))
	manifest.RegisterProfession("Pirate", testProfession("Pirate", conflicts("Dwarf"), conflicts("Elder")))
	manifest.RegisterProfession("Miner", testProfession("Miner", conflicts("Elf"), conflicts()))
	manifest.RegisterProfession("Mason", testProfession("Mason", conflicts("Elf"), conflicts()))
	manifest.RegisterProfession("Mystic", testProfession("Mystic", conflicts("Human"), conflicts()))

	return manifest
}

func TestSelectorSuite(t *testing.T) {
	suite.Run(t, SetupSuite())
}

func (s *SelectorTestSuite) TestNewSelector() {
	x := NewSelector(s.manifest)

	s.Empty(x.raceOptions)
	s.Empty(x.casteOptions)
	s.Empty(x.professionOptions)
}

func (s *SelectorTestSuite) TestRaceOptions() {
	x := NewSelector(s.manifest)

	s.Empty(x.raceOptions)

	x.AddRaceOption("A")
	s.Len(x.raceOptions, 1)
	s.Len(x.GetRaceOptions(), 1)
	s.Contains(x.GetRaceOptions(), "A")

	x.AddRaceOption("B")
	x.AddRaceOption("C")
	s.Len(x.GetRaceOptions(), 3)

	x.AddRaceOption("A")
	s.Len(x.GetRaceOptions(), 3)
}

func (s *SelectorTestSuite) TestCasteOptions() {
	x := NewSelector(s.manifest)

	s.Empty(x.casteOptions)

	x.AddCasteOption("A")
	s.Len(x.casteOptions, 1)
	s.Len(x.GetCasteOptions(), 1)
	s.Contains(x.GetCasteOptions(), "A")

	x.AddCasteOption("B")
	x.AddCasteOption("C")
	s.Len(x.GetCasteOptions(), 3)

	x.AddCasteOption("A")
	s.Len(x.GetCasteOptions(), 3)
}

func (s *SelectorTestSuite) TestProfessionOptions() {
	x := NewSelector(s.manifest)

	s.Empty(x.professionOptions)

	x.AddProfessionOption("A")
	s.Len(x.professionOptions, 1)
	s.Len(x.GetProfessionOptions(), 1)
	s.Contains(x.GetProfessionOptions(), "A")

	x.AddProfessionOption("B")
	x.AddProfessionOption("C")
	s.Len(x.GetProfessionOptions(), 3)

	x.AddProfessionOption("A")
	s.Len(x.GetProfessionOptions(), 3)
}

func (s *SelectorTestSuite) TestPickRace_Single() {
	x := NewSelector(s.manifest)

	x.AddRaceOption("Dwarf")

	r := x.PickRace()

	s.Equal(r.Name, "Dwarf")
}

func (s *SelectorTestSuite) TestPickRace_Random() {
	x := NewSelector(s.manifest)

	x.AddRaceOption("Dwarf")

	r := x.PickRace()

	s.NotNil(r)
}

func (s *SelectorTestSuite) TestGetSelectableRaces_All() {
	x := NewSelector(s.manifest)

	targets := x.GetSelectableRaces()

	s.Len(targets, len(s.manifest.AllRaces()))
}

func (s *SelectorTestSuite) TestGetSelectableRaces_Single() {
	x := NewSelector(s.manifest)

	x.AddRaceOption("Dwarf")
	targets := x.GetSelectableRaces()

	s.Len(targets, 1)
}

func (s *SelectorTestSuite) TestGetSelectableRaces_NoneByCaste() {
	x := NewSelector(s.manifest)

	x.AddRaceOption("Dwarf")
	x.AddCasteOption("Noble")
	targets := x.GetSelectableRaces()

	s.Len(targets, 0)
}

func (s *SelectorTestSuite) TestGetSelectableRaces_NoneByProfession() {
	x := NewSelector(s.manifest)

	x.AddRaceOption("Dwarf")
	x.AddProfessionOption("Pirate")
	targets := x.GetSelectableRaces()

	s.Len(targets, 0)
}

func (s *SelectorTestSuite) TestGetSelectableCastes_All() {
	x := NewSelector(s.manifest)

	targets := x.GetSelectableCastes()

	s.Len(targets, len(s.manifest.AllCastes()))
}

func (s *SelectorTestSuite) TestGetSelectableCastes_Single() {
	x := NewSelector(s.manifest)

	x.AddCasteOption("Noble")
	targets := x.GetSelectableCastes()

	s.Len(targets, 1)
}

func (s *SelectorTestSuite) TestGetSelectableCastes_NoneByRace() {
	x := NewSelector(s.manifest)

	x.AddRaceOption("Dwarf")
	x.AddCasteOption("Noble")
	targets := x.GetSelectableCastes()

	s.Len(targets, 0)
}

func (s *SelectorTestSuite) TestGetSelectableCastes_NoneByProfession() {
	x := NewSelector(s.manifest)

	x.AddCasteOption("Elder")
	x.AddProfessionOption("Pirate")
	targets := x.GetSelectableCastes()

	s.Len(targets, 0)
}

func (s *SelectorTestSuite) TestGetSelectableProfessions_All() {
	x := NewSelector(s.manifest)

	targets := x.GetSelectableProfessions()

	s.Len(targets, len(s.manifest.AllProfessions()))
}

func (s *SelectorTestSuite) TestGetSelectableProfessions_Single() {
	x := NewSelector(s.manifest)

	x.AddProfessionOption("Mystic")
	targets := x.GetSelectableProfessions()

	s.Len(targets, 1)
}

func (s *SelectorTestSuite) TestGetSelectableProfessions_NoneByRace() {
	x := NewSelector(s.manifest)

	x.AddRaceOption("Dwarf")
	x.AddProfessionOption("Pirate")
	targets := x.GetSelectableProfessions()

	s.Len(targets, 0)
}

func (s *SelectorTestSuite) TestGetSelectableProfessions_NoneByCaste() {
	x := NewSelector(s.manifest)

	x.AddCasteOption("Elder")
	x.AddProfessionOption("Pirate")
	targets := x.GetSelectableProfessions()

	s.Len(targets, 0)
}
