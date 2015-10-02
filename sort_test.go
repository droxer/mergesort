package mergesort_test

import (
    "github.com/droxer/mergesort"
    "math/rand"
    "reflect"
    "testing"
)

func TestMSort(t *testing.T) {
    expected := []int{3, 9, 10, 27, 38, 43, 82}
    orig := []int{38, 27, 43, 3, 9, 82, 10}

    mergesort.Sort(orig)

    if !reflect.DeepEqual(orig, expected) {
        t.Fatalf("expected %v, actual is %v", expected, orig)
    }
}

func BenchmarkMSort100(b *testing.B) {
    benchmarkMSort(100, b)
}

func BenchmarkMSort1000(b *testing.B) {
    benchmarkMSort(1000, b)
}

func BenchmarkMSort10000(b *testing.B) {
    benchmarkMSort(10000, b)
}

func benchmarkMSort(i int, b *testing.B) {
    for j := 0; j < b.N; j++ {
        values := rand.Perm(i)
        mergesort.Sort(values)
    }
}
