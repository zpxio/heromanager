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

package game

import (
	"time"
	"sync"
	"log"
	"github.com/zpxio/heromanager/internal/gamernal/game/state"
	"github.com/ghodss/yaml"
	"os"
	"fmt"
	"github.com/zpxio/heromanager/internal/gamernal/game/data"
	"path"
)

type World struct {
	tick *Tick

	running bool
	runningLatch sync.WaitGroup

	state state.State
	saver StateSaver
}

func CreateWorld() *World {
	w := World{tick: Create(1, time.Millisecond * 1000)}

	w.tick.Subscribe(&w)

	return &w
}

func (world *World) Load(dataDirectory string) {
	log.Printf("Loading world resources from: %s", dataDirectory)
	data.Load(path.Join(dataDirectory, "races.yml"))
}

func (world *World) Start() {
	world.running = true
	world.runningLatch.Add(1)

	world.saver = StateSaver{world: world}

	world.tick.Start()
	world.saver.reschedule()
}

func (world *World) Shutdown() {
	world.running = false
	defer func() {
		world.runningLatch.Done()
	}()
}

func (world *World) SaveState() {
	stateYaml, err := yaml.Marshal(world.state)
	if err != nil {
		log.Printf("ERROR: Could not compile world state.")
		return
	}

	timestamp := time.Now().Format("20060102_150405_000")
	statePath := fmt.Sprintf("store/state-%s.yml", timestamp)
	f, err := os.Create(statePath)
	if err != nil {
		log.Printf("ERROR: Could not create state manifest: %s", statePath)
		return
	}
	defer f.Close()

	log.Printf("Saving state to: %s (%d bytes)", statePath, len(stateYaml))
	f.Write(stateYaml)
}

func (world *World) OnTick(id uint64) {
	log.Printf("Executing world updates: T+%d", id)
}

func (world *World) Tick() uint64 {
	return world.tick.id
}

func (world *World) Wait() {
	world.tick.Wait()
}

func (world *World) WaitFor(tickId uint64) {
	world.tick.WaitFor(tickId)
}

func (world *World) AwaitShutdown() {
	world.runningLatch.Wait()
}

type StateSaver struct {
	world *World
}

func (saver *StateSaver) reschedule() {
	targetTick := saver.world.Tick()

	targetTick += 5

	go func(target uint64) {
		saver.world.WaitFor(target)
		saver.world.SaveState()
		saver.reschedule()
	}(targetTick)
}
