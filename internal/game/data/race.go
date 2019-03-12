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
	"github.com/zpxio/heromanager/internal/game/data/table"
	"github.com/zpxio/heromanager/internal/game/util"
	"gopkg.in/yaml.v2"
	"log"
)

type Race struct {
	Name           string `json:"name"`
	BaseAttributes table.Values
}

type RaceManifest struct {
	lookup map[string]Race
}

func LoadRaces(gameDir string, raceFile string) RaceManifest {
	manifest := RaceManifest{}
	manifest.lookup = make(map[string]Race)

	raceYaml, err := util.GameFileData(gameDir, raceFile)
	if err != nil {
		log.Fatalf("yamlFile.Get err %v ", err)
	}

	yaml.Unmarshal(raceYaml, &manifest.lookup)

	return manifest
}
