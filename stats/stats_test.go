package ghproject

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var testData = []IssueData{
	{
		Assignee: "bob",
		Id:       101,
		Labels:   []string{"size/s", "triaged", "area/cfg"},
	},
	{
		Assignee: "sally",
		Id:       202,
		Labels:   []string{"size/l", "area/live", "area/cfg"},
	},
	{
		Assignee: "sally",
		Id:       203,
		Labels:   []string{"size/xl", "area/live", "documentation"},
	},
	{
		Assignee: "sally",
		Id:       212,
		Labels:   []string{"size/s", "area/live", "triaged"},
	},
}

func TestSortByUser(t *testing.T) {
	result := SortByUser(testData)
	assert.Contains(t, result, "bob")
	assert.Contains(t, result, "sally")
	assert.Len(t, result["bob"], 1)
	assert.Len(t, result["sally"], 3)
}

func TestWorkload(t *testing.T) {
	expected := []int{
		1, 4, 8, 1,
	}

	for i := 0; i < 4; i++ {
		result := Workload(testData[i : i+1])
		assert.Equal(t, expected[i], result)
	}
}

func TestWorkloadByUser(t *testing.T) {
	expected := map[string]int{
		"bob":   1,
		"sally": 13,
	}

	results := WorkloadByUser(testData)
	assert.Equal(t, expected, results)
}
