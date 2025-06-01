package main

import (
	"reflect"
	"testing"
)

func TestProcessJobs(t *testing.T) {
	numThreads := 2
	numJobs := 5
	jobs := []int{1, 2, 3, 4, 5}
	expected := []Thread{
		{0, 0},
		{1, 0},
		{0, 1},
		{1, 2},
		{0, 4},
	}

	result := ProcessJobs(numJobs, numThreads, jobs)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}

func TestSingleThread(t *testing.T) {
	numThreads := 1
	numJobs := 4
	jobs := []int{2, 3, 1, 4}
	expected := []Thread{
		{0, 0},
		{0, 2},
		{0, 5},
		{0, 6},
	}
	result := ProcessJobs(numJobs, numThreads, jobs)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("SingleThread: Expected %v, got %v", expected, result)
	}
}

func TestMoreThreadsThanJobs(t *testing.T) {
	numThreads := 5
	numJobs := 3
	jobs := []int{7, 8, 9}
	expected := []Thread{
		{0, 0},
		{1, 0},
		{2, 0},
	}
	result := ProcessJobs(numJobs, numThreads, jobs)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("MoreThreadsThanJobs: Expected %v, got %v", expected, result)
	}
}

func TestAllJobsSameDuration(t *testing.T) {
	numThreads := 3
	numJobs := 6
	jobs := []int{5, 5, 5, 5, 5, 5}
	expected := []Thread{
		{0, 0},
		{1, 0},
		{2, 0},
		{0, 5},
		{1, 5},
		{2, 5},
	}
	result := ProcessJobs(numJobs, numThreads, jobs)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("AllJobsSameDuration: Expected %v, got %v", expected, result)
	}
}

func TestNoJobs(t *testing.T) {
	numThreads := 3
	numJobs := 0
	jobs := []int{}
	expected := []Thread{}
	result := ProcessJobs(numJobs, numThreads, jobs)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("NoJobs: Expected %v, got %v", expected, result)
	}
}
