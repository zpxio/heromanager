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

package hero

import (
	"github.com/stretchr/testify/suite"
	"github.com/zpxio/heromanager/internal/game/data/classifier"
	"testing"
)

type GenerateTestSuite struct {
	suite.Suite
	manifest *classifier.ClassifierManifest
}

func TestGenerateSuite(t *testing.T) {

	s := new(GenerateTestSuite)
	s.manifest = classifier.NewManifest()

	suite.Run(t, s)
}

func (s *GenerateTestSuite) TestGenerate_Basic() {
	x := NewSelector()
	h := Generate(s.manifest, x)

	s.Require().NotNil(h)
}
