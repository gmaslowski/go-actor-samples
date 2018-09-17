package main

import (
	"github.com/AsynkronIT/goconsole"
	"github.com/AsynkronIT/protoactor-go/actor"

	"github.com/gmaslowski/go-actor-samples/internal/app"
)

func main() {
	props := actor.FromInstance(&hello.Greeter{})
	pid := actor.Spawn(props)
	pid.Tell(hello.Greeting{Who: "Roger"})
	pid.Tell(hello.Greeting{Who: "John"})
	pid.Tell(hello.Greeting{Who: "Jimmy"})
	console.ReadLine()
}
