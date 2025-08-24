package main

import (
	"fmt"
	"sync"
	"idgen.local/idgenerator2/idgenerator"
)

func main() {
	gen, err := idgenerator.New(idgenerator.Settings{
		MachineID: func() (uint16, error) { return 42, nil },
	})
	if err != nil {
		panic(err)
	}

	var wg sync.WaitGroup

	const IDs = 20

	wg.Add(1)
	for i := 0; i < 1; i++ {
		go func() {
			wg.Done()
			for j := 0; j < IDs; j++ {
				id, err := gen.NextID()
				if err != nil {
					panic(err)
				}
				fmt.Printf("%v\n", id)
			}
		}()
	}
	wg.Wait()

	fmt.Println("Generated IDs:", IDs)
}
