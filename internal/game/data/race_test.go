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

	data, dataErr := util.GameFileData("testdata/game/data", "test_race_single.yml")
	t.Require().Nil(dataErr)

	err := yaml.Unmarshal(data, &r)

	t.Require().Nil(err)
	t.Equal(1.1, r.BaseAttributes.Factor(attributes.Brawn))
	t.Equal(1.5, r.BaseAttributes.Factor(attributes.Insight))
	t.Equal(1.0, r.BaseAttributes.Factor(attributes.Allure))
}

func (t *RaceTestSuite) TestYamlLoadAll() {
	manifest := LoadAll("testdata/game/data", "test_race_all_simple.yml")

	t.Len(manifest.lookup, 2)
	t.Contains(manifest.lookup, "Dwarf")
	t.Contains(manifest.lookup, "Elf")
}
