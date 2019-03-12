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

type Policy struct {
	min          float64
	max          float64
	defaultValue float64
	keys         KeySet
}

func NewPolicy(min float64, max float64, defaultValue float64, keys []string) *Policy {
	return &Policy{min: min, max: max, defaultValue: defaultValue, keys: NewKeySet(keys...)}
}

func (p *Policy) MaxValue() float64 {
	return p.max
}

func (p *Policy) MinValue() float64 {
	return p.min
}

func (p *Policy) ValidKey(k string) bool {
	return p.keys.Contains(k)
}

func (p *Policy) Clamp(value float64) float64 {
	if value < p.min {
		return p.min
	} else if value > p.max {
		return p.max
	} else {
		return value
	}
}

func (p *Policy) DefaultValue() float64 {
	return p.defaultValue
}

func (p *Policy) ValidKeys() []string {
	return p.keys.All()
}
