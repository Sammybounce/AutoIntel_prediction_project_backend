package array

import (
	"sync"
)

func Filter[T any](data *[]T, condition func(T, int) bool) *[]T {
	var (
		filtered []T
		wg       sync.WaitGroup
	)

	ch := make(chan T)

	for i, d := range *data {
		wg.Add(1)

		go func(i int, d T) {
			defer wg.Done()

			if condition(d, i) {
				ch <- d
			}
		}(i, d)
	}
	go func() {
		wg.Wait()
		close(ch)
	}()

	for d := range ch {
		filtered = append(filtered, d)
	}

	return &filtered
}

func Find[T any](data *[]T, condition func(T) bool) (*T, bool, int) {

	var r T

	for i, d := range *data {
		if condition(d) {
			return &d, true, i
		}
	}
	return &r, false, 0
}
