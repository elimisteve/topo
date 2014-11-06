package main

import (
	"fmt"
	"sync"

	"github.com/mdmarek/topo"
	"github.com/mdmarek/topo/topoutil"
)

const seed = 92882
const nworkers = 2

func main() {
	wg := new(sync.WaitGroup)
	wg.Add(nworkers)

	t := topo.New(seed)
	source, err := topoutil.NewUsaGovSource(t)

	if err != nil {
		fmt.Printf("Failed to open source: %v\n", err)
		return
	}

	// Randomly send messages read from the source
	// to each output channel.
	outputs := t.Shuffle(nworkers, source)

	// Each output channel is read by one Sink, which
	// prints to stdout the messages it receives.
	for i := 0; i < nworkers; i++ {
		go topoutil.Sink(i, wg, outputs[i])
	}

	// Wait for the sinks to finish, if ever.
	wg.Wait()
}
