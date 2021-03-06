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

package main

import (
	"flag"
	"github.com/zpxio/heromanager/internal/api"
	"github.com/zpxio/heromanager/internal/game"
	"log"
)

func main() {

	log.Printf("Starting up...")

	dataDirectory := flag.String("data", "/usr/local/share/heromanager", "The directory to load game data from.")
	flag.Parse()

	log.Printf("Data directory: %s", *dataDirectory)

	world := game.CreateWorld()
	world.Load(*dataDirectory)

	log.Printf("Game world created. Turn=%d", world.Tick())

	apiServer := api.CreateServer(world)

	world.Start()
	apiServer.Start()

	world.AwaitShutdown()
}
