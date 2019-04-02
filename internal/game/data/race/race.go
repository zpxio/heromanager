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

package race

import (
	log "github.com/sirupsen/logrus"
	"github.com/zpxio/heromanager/internal/game/data/classifier"
	"github.com/zpxio/heromanager/internal/game/util"
	"gopkg.in/yaml.v2"
)

type Race struct {
	classifier.Classifier
}

type Manifest struct {
	lookup map[string]Race
}

func Blank() Race {
	r := Race{}

	r.Classifier = classifier.Initialize()

	return r
}

func LoadAll(gameDir string, raceFile string) Manifest {
	manifest := Manifest{}
	manifest.lookup = make(map[string]Race)

	raceYaml, err := util.GameFileData(gameDir, raceFile)
	if err != nil {
		log.Errorf("yamlFile.Get err %v ", err)
	}

	err = yaml.Unmarshal(raceYaml, &manifest.lookup)
	if err != nil {
		log.Errorf("failed to parse race data: %s", err)
	}

	return manifest
}
