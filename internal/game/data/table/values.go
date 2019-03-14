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

type Values struct {
	values map[string]float64
	policy *Policy
}

func NewValues(policy *Policy) Values {
	at := Values{values: make(map[string]float64), policy: policy}

	for _, attr := range policy.ValidKeys() {
		at.values[attr] = policy.defaultValue
	}

	return at
}

func (t *Values) Copy() Values {
	at := NewValues(t.policy)
	for attr, val := range t.values {
		at.values[attr] = t.policy.Clamp(val)
	}

	return at
}

func (t *Values) Adjust(modifier Modifier) Values {

	adjusted := t.Copy()

	for _, key := range t.policy.ValidKeys() {
		adjusted.Set(key, modifier.Apply(key, t.Get(key)))
	}

	return adjusted
}

func (t *Values) Load(values map[string]float64) {
	for id, value := range values {
		if t.policy.ValidKey(id) {
			t.Set(id, value)
		}
	}
}

func (t *Values) Set(key string, value float64) {
	if t.policy.ValidKey(key) {
		t.values[key] = t.policy.Clamp(value)
	}
}

func (t *Values) Get(key string) float64 {
	if value, ok := t.values[key]; ok {
		return value
	} else {
		return t.policy.DefaultValue()
	}
}
