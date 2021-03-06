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

type ProfessionTestSuite struct {
	suite.Suite
}

func TestProfessionSuite(t *testing.T) {
	suite.Run(t, new(ProfessionTestSuite))
}

func (t *ProfessionTestSuite) TestYamlLoad_Single() {
	r := BlankProfession()

	data, dataErr := util.GameFileData("testdata/game/data/profession", "test_profession_single.yml")
	t.Require().Nil(dataErr)

	err := yaml.Unmarshal(data, &r)

	t.Require().Nil(err)
	t.Equal(1.1, r.Attributes.Factor(attributes.Brawn))
	t.Equal(1.5, r.Attributes.Factor(attributes.Insight))
	t.Equal(1.0, r.Attributes.Factor(attributes.Allure))
}

func (t *ProfessionTestSuite) TestYamlLoadAll() {
	manifest := NewManifest()
	LoadProfessions("testdata/game/data/profession", "test_profession_all_simple.yml", manifest)

	t.Len(manifest.AllProfessions(), 2)
	t.Contains(manifest.AllProfessions(), "Trader")
	t.Contains(manifest.AllProfessions(), "Hunter")
}
