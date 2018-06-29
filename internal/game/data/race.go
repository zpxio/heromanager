//------------------------------------------------------------------------------
//    Copyright 2018 Jeff Sharpe (zeropointx.io)
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
	"io/ioutil"
	"log"
	"gopkg.in/yaml.v2"
)

type Race struct {
	Name string `json:"name"`

}

type RaceManifest struct {
	lookup map[string]Race
}

func Load(raceFile string) RaceManifest {
	manifest := RaceManifest{}
	manifest.lookup = make(map[string]Race)

	raceYaml, err := ioutil.ReadFile(raceFile)
	if err != nil {
		log.Fatalf("yamlFile.Get err   #%v ", err)
	}

	yaml.Unmarshal(raceYaml, &manifest.lookup)

	return manifest
}
