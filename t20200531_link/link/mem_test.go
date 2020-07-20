package main

import "testing"

func BenchmarkReverRecursion3(b *testing.B) {
	b.ReportAllocs()
	head := New()
	for i:=1 ;i<50000;i++ {
		head.AddLast(i)
	}

	for j:=1;j<10000;j++ {
		ReverRecursion3(nil,head)
	}
}


func BenchmarkReverRecursion2(b *testing.B) {
	b.ReportAllocs()
	head := New()
	for i:=1 ;i<50000;i++ {
		head.AddLast(i)
	}
	for j:=1;j<10000;j++ {
		head.ReverRecursion2()
	}
}