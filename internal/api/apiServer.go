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

package api

import (
	"log"
	"github.com/gin-gonic/gin"
	"github.com/zpxio/heromanager/internal/game"
	"strconv"
)

type Server struct {
	router *gin.Engine
	world *game.World
}

func CreateServer(world *game.World) *Server {

	server := Server{world: world}
	server.router = gin.Default()

	// Base Middleware for decorating responses
	server.router.Use(server.SetupApiResponse)

	// Basic Logging

	// Standard system API
	server.router.GET("/sys/ping", Ping)

	return &server
}

func (server *Server) SetupApiResponse(c *gin.Context) {
	c.Header("X-HeroManager", "VERSION")
	c.Header("X-Game-Tick", strconv.FormatUint(server.world.Tick(), 10))
}

func (server *Server) Start() {
	go func() {
		log.Printf("Starting API Sever")
		log.Fatal(server.router.Run(":8080"))
	}()

}