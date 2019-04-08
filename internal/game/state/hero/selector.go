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

type Selector struct {
	raceOptions       map[string]bool
	casteOptions      map[string]bool
	professionOptions map[string]bool
}

func NewSelector() *Selector {
	return &Selector{
		raceOptions:       make(map[string]bool, 0),
		casteOptions:      make(map[string]bool, 0),
		professionOptions: make(map[string]bool, 0),
	}
}

func (s *Selector) AddRaceOption(raceId string) {
	s.raceOptions[raceId] = true
}

func (s *Selector) AddCasteOption(casteId string) {
	s.casteOptions[casteId] = true
}

func (s *Selector) AddProfessionOption(professionId string) {
	s.professionOptions[professionId] = true
}

func (s *Selector) GetRaceOptions() []string {
	return keys(s.raceOptions)
}

func (s *Selector) GetCasteOptions() []string {
	return keys(s.casteOptions)
}

func (s *Selector) GetProfessionOptions() []string {
	return keys(s.professionOptions)
}

func keys(set map[string]bool) []string {
	opts := make([]string, len(set), len(set))

	i := 0
	for k := range set {
		opts[i] = k
		i++
	}

	return opts
}
