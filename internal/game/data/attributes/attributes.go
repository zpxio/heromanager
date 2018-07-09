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

import (
	"github.com/zpxio/heromanager/internal/game/data/values"
)

var Keys = make(map[string]string)

const AttributeMax float32 = 500

const Brawn = "BRN"
const Insight = "INS"
const Finesse = "FIN"
const Vigor = "VIG"
const Allure = "ALL"

var Ids = [...]string{Brawn, Insight, Finesse, Vigor, Allure}

func init() {
	Keys["BRN"] = "Brawn"
	Keys["INS"] = "Insight"
	Keys["FIN"] = "Finesse"
	Keys["VIG"] = "Vigor"
	Keys["ALL"] = "Allure"
}

type Base struct {
	values.ValueCollection
}

func Create(baseValue float32) Base {
	attrs := Base{ValueCollection: values.NewCollection(&Validator, baseValue)}

	return attrs
}

func (b *Base) CreateView() values.View {
	return values.CreateView(b.ValueCollection)
}

func CreateModifier() values.Modifier {
	return values.CreateModifier(&Validator)
}
