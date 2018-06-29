//------------------------------------------------------------------------------
//    Copyright 2018 Jeff Sharpe (zeropointx.io)
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

package attributes

type AttributesView struct {
	base *Attributes
	mods map[string]AttributeModifier

	renderedValues map[string]float32
}

func CreateAttributesView(attrs *Attributes) AttributesView {
	view := AttributesView{base: attrs, mods: make(map[string]AttributeModifier)}

	return view
}

func (view *AttributesView) value(name string) float32 {

	if view.renderedValues == nil {
		view.render()
	}

	val, exists := view.renderedValues[name]

	if exists {
		return val
	} else {
		return 0.0
	}
}

func (view *AttributesView) modify(name string, mod AttributeModifier) {

	view.mods[name] = mod
	view.renderedValues = nil
}

func (view *AttributesView) render() {

	rendered := make(map[string]float32)
	for key := range AttributeKeys {
		value := view.base.value(key)

		for _, mod := range view.mods {
			value *= mod.factor(key)
		}

		view.renderedValues[key] = value
	}

	view.renderedValues = rendered
}