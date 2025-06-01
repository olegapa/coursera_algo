package main

// import (
// 	"bufio"
// 	"fmt"
// 	"os"
// 	"strconv"
// 	"strings"
// )

// // Job represents a job to be processed
// type Job struct {
// 	index    int
// 	duration int
// }

// // Thread represents a thread with its next free time
// type Thread struct {
// 	index    int
// 	finishAt int
// }

// // MinHeap is a min-heap of threads based on finishAt, then index
// type MinHeap []Thread

// func (h MinHeap) Len() int           { return len(h) }
// func (h MinHeap) Less(i, j int) bool {
// 	if h[i].finishAt == h[j].finishAt {
// 		return h[i].index < h[j].index
// 	}
// 	return h[i].finishAt < h[j].finishAt
// }
// func (h MinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
// func (h *MinHeap) Push(x Thread)     { *h = append(*h, x) }
// func (h *MinHeap) Pop() Thread {
// 	old := *h
// 	n := len(old)
// 	x := old[n-1]
// 	*h = old[0 : n-1]
// 	return x
// }

// // HeapifyDown maintains the heap property after Pop
// func (h *MinHeap) HeapifyDown(i int) {
// 	n := len(*h)
// 	for {
// 		left, right, smallest := 2*i+1, 2*i+2, i
// 		if left < n && h.Less(left, smallest) {
// 			smallest = left
// 		}
// 		if right < n && h.Less(right, smallest) {
// 			smallest = right
// 		}
// 		if smallest == i {
// 			break
// 		}
// 		h.Swap(i, smallest)
// 		i = smallest
// 	}
// }

// // HeapifyUp maintains the heap property after Push
// func (h *MinHeap) HeapifyUp(i int) {
// 	for i > 0 {
// 		parent := (i - 1) / 2
// 		if !h.Less(i, parent) {
// 			break
// 		}
// 		h.Swap(i, parent)
// 		i = parent
// 	}
// }

// func main() {
// 	reader := bufio.NewReader(os.Stdin)
// 	line1, _ := reader.ReadString('\n')
// 	line2, _ := reader.ReadString('\n')
// 	parts := strings.Fields(line1)
// 	numThreads, _ := strconv.Atoi(parts[0])
// 	numJobs, _ := strconv.Atoi(parts[1])

// 	jobParts := strings.Fields(line2)
// 	jobs := make([]int, numJobs)
// 	for i := 0; i < numJobs; i++ {
// 		jobs[i], _ = strconv.Atoi(jobParts[i])
// 	}

// 	// Initialize heap
// 	heap := make(MinHeap, numThreads)
// 	for i := 0; i < numThreads; i++ {
// 		heap[i] = Thread{index: i, finishAt: 0}
// 	}
// 	// Heapify initial heap
// 	for i := numThreads/2 - 1; i >= 0; i-- {
// 		heap.HeapifyDown(i)
// 	}

// 	for _, duration := range jobs {
// 		// Pop min
// 		thread := heap[0]
// 		fmt.Printf("%d %d\n", thread.index, thread.finishAt)
// 		thread.finishAt += duration
// 		heap[0] = thread
// 		heap.HeapifyDown(0)
// 	}
// }