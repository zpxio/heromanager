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
	"github.com/zpxio/heromanager/internal/game/data/classifier"
	"github.com/zpxio/heromanager/internal/game/util"
)

type Selector struct {
	manifest          *classifier.ClassifierManifest
	raceOptions       map[string]bool
	casteOptions      map[string]bool
	professionOptions map[string]bool
}

func NewSelector(manifest *classifier.ClassifierManifest) *Selector {
	return &Selector{
		manifest:          manifest,
		raceOptions:       make(map[string]bool, 0),
		casteOptions:      make(map[string]bool, 0),
		professionOptions: make(map[string]bool, 0),
	}
}

func (s *Selector) AddRaceOption(raceId string) {
	s.raceOptions[raceId] = true
}

func (s *Selector) AddCasteOption(casteId string) {
	s.casteOptions[casteId] = true
}

func (s *Selector) AddProfessionOption(professionId string) {
	s.professionOptions[professionId] = true
}

func (s *Selector) GetRaceOptions() []string {
	if len(s.raceOptions) == 0 {
		return s.manifest.AllRaces()
	} else {
		return keys(s.raceOptions)
	}
}

func (s *Selector) GetCasteOptions() []string {
	if len(s.casteOptions) == 0 {
		return s.manifest.AllCastes()
	} else {
		return keys(s.casteOptions)
	}
}

func (s *Selector) GetProfessionOptions() []string {
	if len(s.professionOptions) == 0 {
		return s.manifest.AllProfessions()
	} else {
		return keys(s.professionOptions)
	}
}

func (s *Selector) PickRace() *classifier.Race {
	set := s.GetSelectableRaces()
	races := make([]interface{}, len(set), len(set))

	i := 0
	for k := range set {
		r, found := s.manifest.ResolveRace(k)
		if found {
			races[i] = r
			i++
		}
	}

	selected := util.Pick(races)
	selectedRace := races[selected].(*classifier.Race)

	return selectedRace
}

func keys(set map[string]bool) []string {
	opts := make([]string, len(set), len(set))

	i := 0
	for k := range set {
		opts[i] = k
		i++
	}

	return opts
}

func (s *Selector) GetSelectableRaces() map[string]bool {
	selectable := map[string]bool{}

	// Build the base set of selectable options
	for _, k := range s.GetRaceOptions() {
		selectable[k] = true
	}

	// Remove options that are excluded by all caste options
RaceCasteScan:
	for r := range selectable {
		// Look for a valid caste that doesn't exclude the race
		for _, k := range s.GetCasteOptions() {
			caste, found := s.manifest.ResolveCaste(k)
			if found {
				if caste.Conflicts.AllowRace(r) {
					continue RaceCasteScan
				}
			}
		}

		// None found.
		delete(selectable, r)
	}

	// Remove options that are excluded by all profession options
RaceProfessionScan:
	for r := range selectable {
		// Look for a valid profession that doesn't exclude the race
		for _, k := range s.GetProfessionOptions() {
			profession, found := s.manifest.ResolveProfession(k)
			if found {
				if profession.Conflicts.AllowRace(r) {
					continue RaceProfessionScan
				}
			}
		}

		// None found.
		delete(selectable, r)
	}

	return selectable
}

func (s *Selector) GetSelectableCastes() map[string]bool {
	selectable := map[string]bool{}

	// Build the base set of selectable options
	for _, k := range s.GetCasteOptions() {
		selectable[k] = true
	}

	// Remove options that are excluded by all race options
CasteRaceScan:
	for r := range selectable {
		// Look for a valid caste that doesn't exclude the race
		for _, k := range s.GetRaceOptions() {
			race, found := s.manifest.ResolveRace(k)
			if found {
				if race.Conflicts.AllowCaste(r) {
					continue CasteRaceScan
				}
			}
		}

		// None found.
		delete(selectable, r)
	}

	// Remove options that are excluded by all profession options
CasteProfessionScan:
	for r := range selectable {
		// Look for a valid profession that doesn't exclude the race
		for _, k := range s.GetProfessionOptions() {
			profession, found := s.manifest.ResolveProfession(k)
			if found {
				if profession.Conflicts.AllowCaste(r) {
					continue CasteProfessionScan
				}
			}
		}

		// None found.
		delete(selectable, r)
	}

	return selectable
}

func (s *Selector) GetSelectableProfessions() map[string]bool {
	selectable := map[string]bool{}

	// Build the base set of selectable options
	for _, k := range s.GetProfessionOptions() {
		selectable[k] = true
	}

	// Remove options that are excluded by all race options
ProfessionRaceScan:
	for r := range selectable {
		// Look for a valid caste that doesn't exclude the race
		for _, k := range s.GetRaceOptions() {
			race, found := s.manifest.ResolveRace(k)
			if found {
				if race.Conflicts.AllowProfession(r) {
					continue ProfessionRaceScan
				}
			}
		}

		// None found.
		delete(selectable, r)
	}

	// Remove options that are excluded by all profession options
ProfessionCasteScan:
	for r := range selectable {
		// Look for a valid profession that doesn't exclude the race
		for _, k := range s.GetCasteOptions() {
			profession, found := s.manifest.ResolveCaste(k)
			if found {
				if profession.Conflicts.AllowProfession(r) {
					continue ProfessionCasteScan
				}
			}
		}

		// None found.
		delete(selectable, r)
	}

	return selectable
}
