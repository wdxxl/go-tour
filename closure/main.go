package main

import "sync"

func factory(vals []int, multiplier int) []int {
	result := make([]int, len(vals))
	var wg sync.WaitGroup
	for i, v := range vals {
		wg.Add(1)
		go func(i, v int) {
			defer wg.Done()
			result[i] = v * multiplier
		}(i, v)
	}
	wg.Wait()
	return result

}
