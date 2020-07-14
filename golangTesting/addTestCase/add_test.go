package main

import "testing"

// Test Driven Test Cases in Golang

type AddResult struct{
	x int
	y int
	expected int
}

var addResults = []AddResult{
	{1, 1,2},
	{2,3,5},
}

func TestAdd(t *testing.T){
	for _, test:= range addResults{
		result := Add(test.x, test.y)
		if result != test.expected {
			t.Fatal("Expected result not given")
		}
	}
}