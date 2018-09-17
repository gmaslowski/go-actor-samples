package print

import (
	"log"

	"github.com/AsynkronIT/protoactor-go/actor"
)

// Actor
type Printer struct{}

func (state *Printer) Receive(context actor.Context) {
	switch msg := context.Message().(type) {
	case string:
		log.Printf("%v\n", msg)
	}
}
