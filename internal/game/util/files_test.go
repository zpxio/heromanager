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

package util

import (
	"github.com/stretchr/testify/suite"
	"path"
	"testing"
)

type FileTestSuite struct {
	suite.Suite
}

func TestFileSuite(t *testing.T) {
	suite.Run(t, new(FileTestSuite))
}

func (t *FileTestSuite) TestFindAncestor() {

	dirs := []string{"data", "dir1", "dir2", "dir3", "sub", "data"}
	t.Equal(path.Join(dirs[0:3]...), FindAncestor(path.Join(dirs...), "dir2"))
}

func (t *FileTestSuite) TestFindAncestor_NoMatch() {

	dirs := []string{"data", "dir1", "dir2", "dir3", "sub", "data"}
	t.Equal(".", FindAncestor(path.Join(dirs...), "baz"))
}
