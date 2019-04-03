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

import log "github.com/sirupsen/logrus"

type ClassifierManifest struct {
	races       map[string]Race
	castes      map[string]Caste
	professions map[string]Profession

	raceKeys       []string
	casteKeys      []string
	professionKeys []string
}

func NewManifest() *ClassifierManifest {
	return &ClassifierManifest{
		races:       make(map[string]Race),
		castes:      make(map[string]Caste),
		professions: make(map[string]Profession),

		raceKeys:       make([]string, 0),
		casteKeys:      make([]string, 0),
		professionKeys: make([]string, 0),
	}
}

func (m *ClassifierManifest) RegisterRace(id string, r Race) {
	log.Infof("Registering Race: %s", id)
	m.races[id] = r
	m.raceKeys = nil
}

func (m *ClassifierManifest) RegisterCaste(id string, c Caste) {
	log.Infof("Registering Caste: %s", id)
	m.castes[id] = c
	m.casteKeys = nil
}

func (m *ClassifierManifest) RegisterProfession(id string, p Profession) {
	log.Infof("Registering Profession: %s", id)
	m.professions[id] = p
	m.professionKeys = nil
}

func (m *ClassifierManifest) ResolveRace(id string) (*Race, bool) {
	r, found := m.races[id]

	if found {
		return &r, true
	} else {
		return nil, false
	}
}

func (m *ClassifierManifest) ResolveCaste(id string) (*Caste, bool) {
	c, found := m.castes[id]

	if found {
		return &c, true
	} else {
		return nil, false
	}
}

func (m *ClassifierManifest) ResolveProfession(id string) (*Profession, bool) {
	p, found := m.professions[id]

	if found {
		return &p, true
	} else {
		return nil, false
	}
}

func (m *ClassifierManifest) AllRaces() []string {
	if len(m.raceKeys) < 1 {
		m.raceKeys = make([]string, len(m.races), len(m.races))
		i := 0
		for id := range m.races {
			m.raceKeys[i] = id
			i++
		}
	}

	return m.raceKeys
}

func (m *ClassifierManifest) AllCastes() []string {
	if len(m.casteKeys) < 1 {
		m.casteKeys = make([]string, len(m.castes), len(m.castes))
		i := 0
		for id := range m.castes {
			m.casteKeys[i] = id
			i++
		}
	}

	return m.casteKeys
}

func (m *ClassifierManifest) AllProfessions() []string {
	if len(m.professionKeys) < 1 {
		m.professionKeys = make([]string, len(m.professions), len(m.professions))
		i := 0
		for id := range m.professions {
			m.professionKeys[i] = id
			i++
		}
	}

	return m.professionKeys
}
