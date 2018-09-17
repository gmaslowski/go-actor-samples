package main

import (
	"github.com/AsynkronIT/goconsole"
	"github.com/AsynkronIT/protoactor-go/actor"

	"github.com/gmaslowski/go-actor-samples/internal/app/hello"
	"github.com/gmaslowski/go-actor-samples/internal/app/print"
)

func main() {
	greeterProps := actor.FromInstance(&hello.Greeter{})
	greeter := actor.Spawn(greeterProps)

	printerProps := actor.FromInstance(&print.Printer{})
	printer := actor.Spawn(printerProps)

	greeter.Request(hello.Greeting{Who: "Roger"}, printer)
	greeter.Request(hello.Greeting{Who: "John"}, printer)
	greeter.Request(hello.Greeting{Who: "Jimmy"}, printer)

	console.ReadLine()
}
