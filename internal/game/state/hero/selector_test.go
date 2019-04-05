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
	"testing"
)

type SelectorTestSuite struct {
	suite.Suite
}

func TestSelectorSuite(t *testing.T) {
	suite.Run(t, new(SelectorTestSuite))
}

func (s *SelectorTestSuite) TestNewSelector() {
	x := NewSelector()

	s.Empty(x.raceOptions)
	s.Empty(x.casteOptions)
	s.Empty(x.professionOptions)
}

func (s *SelectorTestSuite) TestRaceOptions() {
	x := NewSelector()

	s.Empty(x.raceOptions)

	x.AddRaceOption("A")
	s.Len(x.raceOptions, 1)
	s.Len(x.GetRaceOptions(), 1)
	s.Contains(x.GetRaceOptions(), "A")

	x.AddRaceOption("B")
	x.AddRaceOption("C")
	s.Len(x.GetRaceOptions(), 3)

	x.AddRaceOption("A")
	s.Len(x.GetRaceOptions(), 3)
}

func (s *SelectorTestSuite) TestCasteOptions() {
	x := NewSelector()

	s.Empty(x.casteOptions)

	x.AddCasteOption("A")
	s.Len(x.casteOptions, 1)
	s.Len(x.GetCasteOptions(), 1)
	s.Contains(x.GetCasteOptions(), "A")

	x.AddCasteOption("B")
	x.AddCasteOption("C")
	s.Len(x.GetCasteOptions(), 3)

	x.AddCasteOption("A")
	s.Len(x.GetCasteOptions(), 3)
}

func (s *SelectorTestSuite) TestProfessionOptions() {
	x := NewSelector()

	s.Empty(x.professionOptions)

	x.AddProfessionOption("A")
	s.Len(x.professionOptions, 1)
	s.Len(x.GetProfessionOptions(), 1)
	s.Contains(x.GetProfessionOptions(), "A")

	x.AddProfessionOption("B")
	x.AddProfessionOption("C")
	s.Len(x.GetProfessionOptions(), 3)

	x.AddProfessionOption("A")
	s.Len(x.GetProfessionOptions(), 3)
}
