package hello

import (
	"fmt"

	"github.com/AsynkronIT/protoactor-go/actor"
)

// Actor API
type Greeting struct{ Who string }

// Actor
type Greeter struct{}

func (state *Greeter) Receive(context actor.Context) {
	switch msg := context.Message().(type) {
	case Greeting:
		fmt.Printf("Hello %v\n", msg.Who)
	}
}
