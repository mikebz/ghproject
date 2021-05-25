package ghproject

import "strings"

// this is the relevant issue data
// which takes away the necessity to create
// GitHub issues for testing and mocking
type IssueData struct {
	Assignee string
	Id       int
	Labels   []string
}

var sizeMap = map[string]int{
	"size/s":   1,
	"size/m":   2,
	"size/l":   4,
	"size/xl":  8,
	"size/xxl": 16,
}

// a sorting function that takes a bunch of issues
// and sorts it by the assignee
func SortByUser(issues []IssueData) map[string][]IssueData {
	result := make(map[string][]IssueData)
	for _, issue := range issues {
		dataSlice, ok := result[issue.Assignee]

		if !ok {
			dataSlice = make([]IssueData, 1, 10)
			dataSlice[0] = issue
		} else {
			dataSlice = append(dataSlice, issue)
		}

		result[issue.Assignee] = dataSlice
	}
	return result
}

// Get the sum of the workload
func Workload(issues []IssueData) int {
	result := 0
	for _, issue := range issues {
		for _, label := range issue.Labels {
			size := LabelToSize(label)
			result += size
			if size > 0 {
				break
			}
		}
	}
	return result
}

// Given the array of issue data return an array of workload by user.
func WorkloadByUser(issues []IssueData) map[string]int {
	userData := SortByUser(issues)
	result := make(map[string]int)
	for user, issues := range userData {
		workload := Workload(issues)
		result[user] = workload
	}
	return result
}

// get the size based on the label
// created because it handles upper and lower case
func LabelToSize(label string) int {
	label = strings.ToLower(label)
	result, ok := sizeMap[label]
	if ok {
		return result
	} else {
		return 0
	}
}
