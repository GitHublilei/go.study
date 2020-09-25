package leetcode

import (
	"reflect"
	"testing"
)

func TestTwoSum(t *testing.T) {
	type testCase struct {
		nums   []int
		target int
		want   []int
	}

	testGroup := map[string]testCase{
		"case_1": testCase{[]int{2, 7, 11, 15}, 9, []int{0, 1}},
	}

	for name, tc := range testGroup {
		t.Run(name, func(t *testing.T) {
			got := twoSum(tc.nums, tc.target)
			if !reflect.DeepEqual(got, tc.want) {
				t.Fatalf("want:%#v but get:%#v\n", tc.want, got)
			}
		})
	}
}

func TestTwoSum2(t *testing.T) {
	type testCase struct {
		nums   []int
		target int
		want   []int
	}

	testGroup := map[string]testCase{
		"case_1": testCase{[]int{2, 7, 11, 15}, 9, []int{0, 1}},
	}

	for name, tc := range testGroup {
		t.Run(name, func(t *testing.T) {
			got := twoSum(tc.nums, tc.target)
			if !reflect.DeepEqual(got, tc.want) {
				t.Fatalf("want:%#v but get:%#v\n", tc.want, got)
			}
		})
	}
}

func TestTwoSum3(t *testing.T) {
	type testCase struct {
		nums   []int
		target int
		want   []int
	}

	testGroup := map[string]testCase{
		"case_1": testCase{[]int{2, 7, 11, 15}, 9, []int{0, 1}},
	}

	for name, tc := range testGroup {
		t.Run(name, func(t *testing.T) {
			got := twoSum(tc.nums, tc.target)
			if !reflect.DeepEqual(got, tc.want) {
				t.Fatalf("want:%#v but get:%#v\n", tc.want, got)
			}
		})
	}
}
