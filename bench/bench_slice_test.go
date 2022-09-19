package bench

import (
	"testing"
)

func BenchmarkReorganizeStringSlice(b *testing.B) {
	for i := 0; i < b.N; i++ {
		r := ReorganizeStringSlice(QueryString)
		_ = r
	}
}
