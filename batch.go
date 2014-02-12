package batch

import (
//"fmt"
)

func Batcher(slice []interface{}, count int64) (c chan interface{}) {
	c = make(chan interface{})

	go func() {

		batch := []interface{}{}

		var i int64

		for _, item := range slice {
			i++
			if i%count == 0 {

				c <- batch
				batch = nil

			}

			batch = append(batch, item)

		}

		close(c)

	}()

	return
}
