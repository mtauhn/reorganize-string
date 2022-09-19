package bench

import (
	"testing"
)

func BenchmarkReorganizeString(b *testing.B) {
	for i := 0; i < b.N; i++ {
		r := ReorganizeString(QueryString)
		_ = r
	}
}
