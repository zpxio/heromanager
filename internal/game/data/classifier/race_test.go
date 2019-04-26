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
	"github.com/ghodss/yaml"
	"github.com/stretchr/testify/suite"
	"github.com/zpxio/heromanager/internal/game/data/attributes"
	"github.com/zpxio/heromanager/internal/game/util"
	"testing"
)

type RaceTestSuite struct {
	suite.Suite
}

func TestRaceSuite(t *testing.T) {
	suite.Run(t, new(RaceTestSuite))
}

func (t *RaceTestSuite) TestYamlLoad_Single() {
	r := BlankRace()

	data, dataErr := util.GameFileData("testdata/game/data/race", "test_race_single.yml")
	t.Require().Nil(dataErr)

	err := yaml.Unmarshal(data, &r)

	t.Require().Nil(err)
	t.Equal(1.1, r.Attributes.Factor(attributes.Brawn))
	t.Equal(1.5, r.Attributes.Factor(attributes.Insight))
	t.Equal(1.0, r.Attributes.Factor(attributes.Allure))
}

func (t *RaceTestSuite) TestYamlLoadAll() {
	manifest := NewManifest()
	LoadRaces("testdata/game/data/race", "test_race_all_simple.yml", manifest)

	t.Len(manifest.AllRaces(), 2)
	t.Contains(manifest.AllRaces(), "Dwarf")
	t.Contains(manifest.AllRaces(), "Elf")
}

func (t *RaceTestSuite) TestLoadRaces_FNF() {
	manifest := NewManifest()
	err := LoadRaces("testdata/game/data/race", "redundant-raccoon.yml", manifest)

	t.NotNil(err)
	t.Len(manifest.AllRaces(), 0)
}

func (t *RaceTestSuite) TestLoadRaces_BadFormat() {
	manifest := NewManifest()
	err := LoadRaces("testdata/game/data/race", "test_race_all_bad.yml", manifest)

	t.NotNil(err)
	t.Len(manifest.AllRaces(), 0)
}
