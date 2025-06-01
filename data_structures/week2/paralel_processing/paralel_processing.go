package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Thread struct {
	idx           int
	executionTime int
}

type PriorityQueue []Thread

func GetLeftChildIdx(i int) int {
	return 2*i + 1
}

func GetRightChildIdx(i int) int {
	return 2*i + 2
}

func GetParentIdx(i int) int {
	return (i - 1) / 2
}

func (queue PriorityQueue) isLess(idx1, idx2 int) bool {
	return queue[idx1].executionTime < queue[idx2].executionTime
}

func (queue *PriorityQueue) Swap(i, j int) {
	(*queue)[i], (*queue)[j] = (*queue)[j], (*queue)[i]
}

func (queue *PriorityQueue) SiftDown(idx int) {
	leftChildidx, rightChildIdx, minIdx := GetLeftChildIdx(idx), GetRightChildIdx(idx), idx
	if leftChildidx < len(*queue) && queue.isLess(leftChildidx, minIdx) {
		minIdx = leftChildidx
	}
	if rightChildIdx < len(*queue) && queue.isLess(rightChildIdx, minIdx) {
		minIdx = rightChildIdx
	}
	if minIdx != idx {
		queue.Swap(idx, minIdx)
		queue.SiftDown(minIdx)
	}
}

func (queue *PriorityQueue) SiftUp(idx int) {
	if idx > 0 {
		parentIdx := GetParentIdx(idx)
		if queue.isLess(idx, parentIdx) {
			queue.Swap(idx, parentIdx)
			queue.SiftUp(parentIdx)
		}
	}

}

func (queue *PriorityQueue) insert(thread Thread) {
	*queue = append(*queue, thread)
	queue.SiftUp(len(*queue) - 1)
}

func (queue *PriorityQueue) GetMin() Thread {
	minThread, size := (*queue)[0], len(*queue)
	queue.Swap(0, size-1)
	(*queue) = (*queue)[:size-1]
	queue.SiftDown(0)
	return minThread
}

func ProcessJobs(numJobs, numThreads int, jobs []int) []Thread {
	queue := make(PriorityQueue, 0, numThreads)
	result := make([]Thread, 0, numJobs)
	for i := 0; i < numThreads; i++ {
		if i >= numJobs {
			return result
		}
		queue = append(queue, Thread{i, jobs[i]})
		result = append(result, Thread{i, 0})
	}

	for pos := len(queue); pos < numJobs; pos++ {
		minTread := queue.GetMin()
		result = append(result, minTread)
		minTread.executionTime += jobs[pos]
		queue.insert(minTread)
	}
	return result
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	line1, _ := reader.ReadString('\n')
	line2, _ := reader.ReadString('\n')
	parts := strings.Fields(line1)
	numThreads, _ := strconv.Atoi(parts[0])
	numJobs, _ := strconv.Atoi(parts[1])

	jobParts := strings.Fields(line2)
	jobs := make([]int, numJobs)
	for i := 0; i < numJobs; i++ {
		jobs[i], _ = strconv.Atoi(jobParts[i])
	}

	threads := ProcessJobs(numJobs, numThreads, jobs)

	for i := 0; i < len(threads); i++ {
		fmt.Println(threads[i].idx, threads[i].executionTime)
	}
}
