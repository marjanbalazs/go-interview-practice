package main

import (
	"sync"
)

// ConcurrentBFSQueries concurrently processes BFS queries on the provided graph.
// - graph: adjacency list, e.g., graph[u] = []int{v1, v2, ...}
// - queries: a list of starting nodes for BFS.
// - numWorkers: how many goroutines can process BFS queries simultaneously.
//
// Return a map from the query (starting node) to the BFS order as a slice of nodes.
// YOU MUST use concurrency (goroutines + channels) to pass the performance tests.

func queryGraph(graph map[int][]int, query int) []int {
	ret := make([]int, 0)
	queue := []int{query}
	visited := make(map[int]bool)

	for len(queue) > 0 {
		elem := queue[0]
		queue = queue[1:]
		visited[elem] = true

		ret = append(ret, elem)

		for _, val := range graph[elem] {
			if !visited[val] {
				visited[val] = true
				queue = append(queue, val)
			}
		}
	}

	return ret
}

func ConcurrentBFSQueries(graph map[int][]int, queries []int, numWorkers int) map[int][]int {
	jobs := make(chan int, len(queries))

	var result sync.Map
	var wg sync.WaitGroup

	for range numWorkers {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := range jobs {
				res := queryGraph(graph, j)
				result.Store(j, res)
			}
		}()
	}

	for _, q := range queries {
		jobs <- q
	}

	close(jobs)
	wg.Wait()

	finalRes := make(map[int][]int, len(queries))

	result.Range(func(key, value any) bool {
		finalRes[key.(int)] = value.([]int)
		return true
	})

	return finalRes
}

func main() {
	// You can insert optional local tests here if desired.
}
