package strand

import (
	"fmt"
	"reflect"
	"runtime"
	"strings"
	"testing"
)

var oldFunction = ToRNA_old
var newFunction = ToRNA

// old is a benchmark function that measures the performance of the
// oldFunction function.
//
// It takes a pointer to a testing.B object as a parameter. The testing.B
// object is used to control the benchmarking process. The function doesn't
// return anything.
func old(b *testing.B) {
	for i := 0; i < b.N; i++ {
		oldFunction("ACGT")
	}
}

// new is a benchmark function that measures the performance of the
// newFunction function.
//
// It takes a pointer to a testing.B object as a parameter. The testing.B
// object is used to control the benchmarking process. The function doesn't
// return anything.
func new(b *testing.B) {
	for i := 0; i < b.N; i++ {
		newFunction("ACGT")
	}
}

// TestDifference is a Go function that calculates the time difference between
// two benchmark results.
//
// It takes no parameters and does not return anything.
func TestDifference(t *testing.T) {
	newFunctionBenchmark := testing.Benchmark(new)
	oldFunctionBenchmark := testing.Benchmark(old)

	changePercent := 100 * float64(oldFunctionBenchmark.NsPerOp()) / float64(newFunctionBenchmark.NsPerOp())
	oldFuncionName := runtime.FuncForPC(reflect.ValueOf(oldFunction).Pointer()).Name()
	newFuncionName := runtime.FuncForPC(reflect.ValueOf(newFunction).Pointer()).Name()
	oldFuncionName = strings.Split(oldFuncionName, ".")[1]
	newFuncionName = strings.Split(newFuncionName, ".")[1]

	var coloredChangePercent string
	if changePercent > 100 {
		coloredChangePercent = fmt.Sprintf("\033[32m%.0f%%\033[0m", changePercent)
	} else {
		coloredChangePercent = fmt.Sprintf("\033[31m%.0f%%\033[0m", changePercent)
	}

	fmt.Printf("The performance of the \033[0;34m%s\033[0m is %s relative to \033[0;34m%s\033[0m\n", newFuncionName, coloredChangePercent, oldFuncionName)
}
