package colorjson_test

import (
	"io/ioutil"
	"testing"
)
import "github.com/relvacode/colorjson"

func benchmarkMarshall(i int, b *testing.B) {
	simpleMap := make(map[string]interface{})
	simpleMap["a"] = 1
	simpleMap["b"] = "bee"
	simpleMap["c"] = [3]float64{1, 2, 3}
	simpleMap["d"] = [3]string{"one", "two", "three"}

	// run the Fib function b.N times
	for n := 0; n < b.N; n++ {
		colorjson.Marshal(ioutil.Discard, simpleMap)
	}
}

func BenchmarkMarshall(b *testing.B)     { benchmarkMarshall(100, b) }
func BenchmarkMarshall1k(b *testing.B)   { benchmarkMarshall(1000, b) }
