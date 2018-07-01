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

type View struct {
	base *Base
	mods map[string]*Modifier

	renderedValues map[string]float32
}

func CreateView(attrs *Base) View {
	view := View{base: attrs, mods: make(map[string]*Modifier)}

	return view
}

func (view *View) Value(name string) float32 {

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

func (view *View) Modify(name string, mod *Modifier) {

	view.mods[name] = mod

	// Don't force the values to be recalculated.
	// Just clear the storage to trigger a recalculation the next time anyone actually needs them.
	view.renderedValues = nil
}

func (view *View) render() {

	rendered := make(map[string]float32)
	for key := range Keys {
		value := view.base.Value(key)

		for _, mod := range view.mods {
			value *= mod.Factor(key)
		}

		rendered[key] = value
	}

	view.renderedValues = rendered
}
