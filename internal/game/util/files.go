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
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"log"
	"path"
	"path/filepath"
	"runtime"
)

var GameDirBasePath string

func init() {
	_, filename, _, _ := runtime.Caller(0)

	GameDirBasePath = FindAncestor(filename, "heromanager")

	log.Printf("Using Game Directory: %s", GameDirBasePath)
}

func FindAncestor(dir string, targetDir string) string {
	ancestor := dir

	for path.Base(ancestor) != targetDir {
		ancestor = path.Dir(ancestor)
		if ancestor == "." {
			return ancestor
		}
	}

	return ancestor
}

func GameFileData(gameDir string, dataPath string) ([]byte, error) {
	relpath := path.Join(GameDirBasePath, gameDir, dataPath)

	abspath, err := filepath.Abs(relpath)
	if err != nil {
		log.Printf("Failed to resolve absolute path for: %s (%s)", relpath, err)
	}

	data, err := ioutil.ReadFile(abspath)
	if err != nil {
		logrus.Errorf("file data read: %v ", err)
		return make([]byte, 0), err
	}

	return data, err
}
