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
	log "github.com/sirupsen/logrus"
	"strings"
)

const (
	ConflictRaces       = "races"
	ConflictCastes      = "castes"
	ConflictProfessions = "professions"
)

type ConflictGroup struct {
	races       map[string]bool
	castes      map[string]bool
	professions map[string]bool
}

func EmptyConflicts() ConflictGroup {
	return ConflictGroup{
		races:       make(map[string]bool),
		castes:      make(map[string]bool),
		professions: make(map[string]bool),
	}
}

func (c *ConflictGroup) Load(conflictData map[string][]string) {
	// Load races
	for _, id := range conflictData[ConflictRaces] {
		c.Add(ConflictRaces, id)
	}

	// Load castes
	for _, id := range conflictData[ConflictCastes] {
		c.Add(ConflictCastes, id)
	}

	// Load professions
	for _, id := range conflictData[ConflictProfessions] {
		c.Add(ConflictProfessions, id)
	}
}

func (c *ConflictGroup) Add(target string, id string) {
	switch strings.ToLower(target) {
	case ConflictRaces:
		c.races[id] = true
	case ConflictCastes:
		c.castes[id] = true
	case ConflictProfessions:
		c.professions[id] = true
	default:
		log.Warnf("abnormal conflict target: %s:%s", target, id)
	}
}

func (c *ConflictGroup) AllowRace(id string) bool {
	_, ok := c.races[id]

	return !ok
}

func (c *ConflictGroup) AllowCaste(id string) bool {
	_, ok := c.castes[id]

	return !ok
}

func (c *ConflictGroup) AllowProfession(id string) bool {
	_, ok := c.professions[id]

	return !ok
}

func (c *ConflictGroup) UnmarshalJSON(data []byte) error {
	conflicts := map[string][]string{}
	err := json.Unmarshal(data, &conflicts)
	if err != nil {
		return err
	}

	c.Load(conflicts)

	return nil
}

func (c *ConflictGroup) UnmarshalYAML(unmarshal func(interface{}) error) error {
	conflicts := map[string][]string{}
	err := unmarshal(&conflicts)
	if err != nil {
		return err
	}

	c.Load(conflicts)

	return nil
}
