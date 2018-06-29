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
	"log"
	"sync"
	"time"
)

type TickListener interface {
	OnTick(id uint64)
}

type Tick struct {
	id      uint64
	enabled bool
	delay   time.Duration
	barrier *sync.Cond

	subscribers []TickListener
}

func Create(initialId uint64, delay time.Duration) *Tick {
	log.Printf("Initializing game tick counter at %d", initialId)
	ticker := Tick{id: initialId, enabled: false, delay: delay, barrier: sync.NewCond(&sync.Mutex{}), subscribers: []TickListener{}}

	return &ticker
}

func (t *Tick) Start() {

	t.enabled = true
	go func() {
		lastTick := time.Now()
		nextTick := lastTick.Add(t.delay)

		for t.enabled {
			nextTick = lastTick.Add(t.delay)
			time.Sleep(time.Until(nextTick))
			lastTick = nextTick

			for _, s := range t.subscribers {
				s.OnTick(t.id)
			}

			t.Next()
		}
	}()
}

func (t *Tick) Stop() {
	t.enabled = false
}

func (t *Tick) Subscribe(subscriber TickListener) {
	t.subscribers = append(t.subscribers, subscriber)
}

func (t *Tick) Next() {
	lastTick := t.barrier
	defer lastTick.Broadcast()

	t.barrier = sync.NewCond(&sync.Mutex{})
	t.id++
}

func (t *Tick) Wait() {
	t.barrier.L.Lock()
	defer t.barrier.L.Unlock()

	t.barrier.Wait()
}

func (t *Tick) WaitFor(tickId uint64) {
	//log.Printf("tickId(%d) <= tick(%d)", tickId, t.id)
	for tickId > t.id {
		t.Wait()
	}
}
