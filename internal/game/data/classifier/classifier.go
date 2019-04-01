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

package classifier

import (
	"github.com/zpxio/heromanager/internal/game/data/attributes"
	"github.com/zpxio/heromanager/internal/game/data/table"
)

type Classifier struct {
	Name       string         `yaml:"name"`
	Attributes table.Modifier `yaml:"attributes"`

	Conflicts ConflictGroup `yaml:"conflicts"`
}

func Initialize() Classifier {
	c := Classifier{
		Name:       "",
		Attributes: attributes.NewAttributeModifier(),

		Conflicts: EmptyConflicts(),
	}

	return c
}
