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

package values

type KeySet struct {
	keys     map[string]bool
	allCache []string
}

func NewKeySet(keys ...string) KeySet {
	s := KeySet{keys: make(map[string]bool)}

	for _, key := range keys {
		s.keys[key] = true
	}

	return s
}

func (s *KeySet) Add(key string) {
	s.keys[key] = true
	s.allCache = nil
}

func (s *KeySet) Contains(key string) bool {
	_, ok := s.keys[key]
	return ok
}

func (s *KeySet) All() []string {
	if s.allCache == nil {
		s.allCache = make([]string, len(s.keys), len(s.keys))

		// Populate the index
		index := 0
		for k := range s.keys {
			s.allCache[index] = k
			index++
		}
	}

	return s.allCache
}
