package race_condition

import "sync"

type SharedData struct {
	money int
	mutex sync.Mutex
}

func (x *SharedData) addMoney(amount int) {
	x.mutex.Lock()
	defer x.mutex.Unlock()

	x.money += amount

}
