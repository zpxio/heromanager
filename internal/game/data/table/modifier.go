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

package table

import (
	"encoding/json"
)

type Modifier struct {
	adjustments map[string]float64
	policy      *Policy
}

func NewModifier(policy *Policy) Modifier {
	m := Modifier{adjustments: make(map[string]float64, len(policy.ValidKeys()))}

	m.policy = policy
	for _, k := range policy.ValidKeys() {
		m.set(k, 0.0)
	}

	return m
}

func (m *Modifier) Load(adjustments map[string]float64) {
	for k, v := range adjustments {
		m.set(k, v)
	}
}

func (m *Modifier) set(key string, factor float64) {
	if m.policy.ValidKey(key) {
		m.adjustments[key] = factor
	}
}

func (m *Modifier) Add(a Modifier) {
	for _, k := range m.policy.ValidKeys() {
		m.adjustments[k] = m.adjustments[k] + a.adjustments[k]
	}
}

func (m *Modifier) Factor(key string) float64 {
	factor, exists := m.adjustments[key]
	if exists {
		return 1.0 + factor
	} else {
		return 1.0
	}
}

func (m *Modifier) Apply(key string, value float64) float64 {
	return value * m.Factor(key)
}

func (m *Modifier) UnmarshalJSON(data []byte) error {
	valueTable := map[string]float64{}
	err := json.Unmarshal(data, &valueTable)
	if err != nil {
		return err
	}

	m.Load(valueTable)

	return nil
}

func (m *Modifier) UnmarshalYAML(unmarshal func(interface{}) error) error {
	modTable := map[string]float64{}
	err := unmarshal(&modTable)
	if err != nil {
		return err
	}

	m.Load(modTable)

	return nil
}
