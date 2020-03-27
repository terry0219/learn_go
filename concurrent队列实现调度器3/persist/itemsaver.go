//新增itemsaver模块

package persist

import "fmt"

func ItemSaver() chan interface{} {
	out := make(chan interface{})

	go func() {
		itemCount := 0
		for {
			item := <- out
			fmt.Printf("ItemSaver Got item #%d, %v\n",itemCount,item)
			itemCount++
		}

	}()
	return out
}
