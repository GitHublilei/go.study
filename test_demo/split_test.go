package test_demo

import (
	"reflect"
	"testing"
)

// func TestSplit(t *testing.T) {
// 	type testCase struct {
// 		str  string
// 		sep  string
// 		want []string
// 	}

// 	testGroup := []testCase{
// 		testCase{"babcbef", "b", []string{"", "a", "c", "ef"}},
// 		testCase{"a:b:c", ":", []string{"a", "b", "c"}},
// 		testCase{"abcef", "bc", []string{"a", "ef"}},
// 	}

// 	for _, tc := range testGroup {
// 		got := Split(tc.str, tc.sep)
// 		if !reflect.DeepEqual(got, tc.want) {
// 			t.Fatalf("want:%#v got:%#v\n", tc.want, got)
// 		}
// 	}
// }

// 子测试
// go test -run=TestSplit/case_1
func TestSplit(t *testing.T) {
	type testCase struct {
		str  string
		sep  string
		want []string
	}

	testGroup := map[string]testCase{
		"case_1": testCase{"babcbef", "b", []string{"", "a", "c", "ef"}},
		"case_2": testCase{"a:b:c", ":", []string{"a", "b", "c"}},
		"case_3": testCase{"abcef", "bc", []string{"a", "ef"}},
	}

	for name, tc := range testGroup {
		t.Run(name, func(t *testing.T) {
			got := Split(tc.str, tc.sep)
			if !reflect.DeepEqual(got, tc.want) {
				t.Fatalf("want:%#v got:%#v\n", tc.want, got)
			}
		})
	}
}

// 基准测试
// go test -bench=Split -benchmem
func BenchmarkSplit(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Split("a:b:c:d:e", ":")
	}
}

// 性能比较测试
func benchmarkFib(b *testing.B, n int) {
	for i := 0; i < b.N; i++ {
		Fib(n)
	}
}

func BenchmarkFib1(b *testing.B) {
	benchmarkFib(b, 1)
}
func BenchmarkFib2(b *testing.B) {
	benchmarkFib(b, 2)
}
func BenchmarkFib3(b *testing.B) {
	benchmarkFib(b, 3)
}
func BenchmarkFib10(b *testing.B) {
	benchmarkFib(b, 10)
}
func BenchmarkFib20(b *testing.B) {
	benchmarkFib(b, 20)
}
func BenchmarkFib40(b *testing.B) {
	benchmarkFib(b, 40)
}
