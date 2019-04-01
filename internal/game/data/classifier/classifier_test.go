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
	"github.com/stretchr/testify/suite"
	"github.com/zpxio/heromanager/internal/game/data/attributes"
	"github.com/zpxio/heromanager/internal/game/util"
	"gopkg.in/yaml.v2"
	"testing"
)

type ClassifierTestSuite struct {
	suite.Suite
}

func TestClassifierSuite(t *testing.T) {
	suite.Run(t, new(ClassifierTestSuite))
}

func (s *ClassifierTestSuite) TestInitialize() {
	c := Initialize()

	s.Equal("", c.Name)
}

func (s *ClassifierTestSuite) TestUnmarshallYAML() {
	c := Initialize()

	data, dataErr := util.GameFileData("testdata/game/data/classifier", "test_classifier_simple.yml")
	s.Require().Nil(dataErr)

	err := yaml.Unmarshal(data, &c)
	s.Require().Nil(err)

	s.Equal("Tester", c.Name)
	s.False(c.Conflicts.AllowRace("Dwarf"))
	s.True(c.Conflicts.AllowRace("Elf"))
	s.False(c.Conflicts.AllowCaste("Noble"))
	s.True(c.Conflicts.AllowCaste("Peasant"))
	s.False(c.Conflicts.AllowProfession("Pirate"))
	s.True(c.Conflicts.AllowProfession("Architect"))

	s.Equal(1.2, c.Attributes.Factor(attributes.Allure))
	s.Equal(0.88, c.Attributes.Factor(attributes.Brawn))
}
