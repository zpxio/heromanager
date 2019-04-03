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
	log "github.com/sirupsen/logrus"
	"github.com/zpxio/heromanager/internal/game/util"
	"gopkg.in/yaml.v2"
)

type Caste struct {
	Classifier
}

func BlankCaste() Caste {
	r := Caste{}

	r.Classifier = Initialize()

	return r
}

func LoadCastes(gameDir string, casteFile string, manifest *ClassifierManifest) {

	casteYaml, err := util.GameFileData(gameDir, casteFile)
	if err != nil {
		log.Errorf("yamlFile.Get err %v ", err)
	}

	castes := make(map[string]Caste)
	err = yaml.Unmarshal(casteYaml, &castes)
	if err != nil {
		log.Errorf("failed to parse caste data: %s", err)
	}

	// Register the castes
	for id, caste := range castes {
		manifest.RegisterCaste(id, caste)
	}
}
