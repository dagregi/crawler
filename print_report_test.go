package main

import (
	"reflect"
	"testing"
)

func TestSortPages(t *testing.T) {
	testCases := []struct {
		name       string
		inputPages []Page
		expected   []Page
	}{
		{
			name: "sort by page count",
			inputPages: []Page{
				{URL: "test.com/page2", Count: 8},
				{URL: "test.com", Count: 1},
				{URL: "test.com/page5", Count: 5},
				{URL: "test.com/page1", Count: 3},
				{URL: "test.com/page4", Count: 7},
				{URL: "test.com/page3", Count: 9},
			},
			expected: []Page{
				{URL: "test.com/page3", Count: 9},
				{URL: "test.com/page2", Count: 8},
				{URL: "test.com/page4", Count: 7},
				{URL: "test.com/page5", Count: 5},
				{URL: "test.com/page1", Count: 3},
				{URL: "test.com", Count: 1},
			},
		},
		{
			name: "sort by page count and alphabetically",
			inputPages: []Page{
				{URL: "test.com/other/2", Count: 3},
				{URL: "test.com/after", Count: 7},
				{URL: "test.com", Count: 19},
				{URL: "test.com/other/3", Count: 3},
				{URL: "test.com/another/page", Count: 7},
				{URL: "test.com/page", Count: 7},
				{URL: "test.com/before", Count: 9},
				{URL: "test.com/bifrost", Count: 9},
				{URL: "test.com/other/1", Count: 3},
				{URL: "test.com/author", Count: 3},
			},
			expected: []Page{
				{URL: "test.com", Count: 19},
				{URL: "test.com/before", Count: 9},
				{URL: "test.com/bifrost", Count: 9},
				{URL: "test.com/after", Count: 7},
				{URL: "test.com/another/page", Count: 7},
				{URL: "test.com/page", Count: 7},
				{URL: "test.com/author", Count: 3},
				{URL: "test.com/other/1", Count: 3},
				{URL: "test.com/other/2", Count: 3},
				{URL: "test.com/other/3", Count: 3},
			},
		},
	}
	for i, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actual := sortPages(tc.inputPages)
			if !reflect.DeepEqual(actual, tc.expected) {
				t.Errorf("Test %v - %s FAIL:\nexpected: %v\nactual: %v", i, tc.name, tc.expected, actual)
			}
		})
	}
}
