package race_condition

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func MyRate() {

	sharedData := SharedData{1000, sync.Mutex{}}

	for i := 1; i <= 10; i++ {
		wg.Add(1)

		go func() {
			defer wg.Done()
			sharedData.addMoney(20)
		}()

	}

	wg.Wait()

	fmt.Println("Done - so tien la: ", sharedData.money)

	// output: Done - so tien la:  1200
}
