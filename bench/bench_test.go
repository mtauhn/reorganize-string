package bench

import (
	"testing"
)

func BenchmarkReorganizeString(b *testing.B) {
	var r string

	for i := 0; i < b.N; i++ {
		r = ReorganizeString(QueryString)
	}
	_ = r
}

func BenchmarkReorganizeStringSlice(b *testing.B) {
	var r string

	for i := 0; i < b.N; i++ {
		r = ReorganizeStringSlice(QueryString)
	}
	_ = r
}

func BenchmarkReorganizeStringVer3(b *testing.B) {
	var r string

	for i := 0; i < b.N; i++ {
		r = reorganizeStringVer3(QueryString)
	}
	_ = r
}
