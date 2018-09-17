package print

import (
	"fmt"

	"github.com/AsynkronIT/protoactor-go/actor"
)

// Actor
type Printer struct{}

func (state *Printer) Receive(context actor.Context) {
	switch msg := context.Message().(type) {
	case string:
		fmt.Printf("%v\n", msg)
	}
}
