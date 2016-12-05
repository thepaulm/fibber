package main

import "testing"

const runs = 10000000

// 0 overlap
func BenchmarkFibber_f1o0(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Run(1, runs, 0)
	}
}

func BenchmarkFibber_f2o0(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Run(2, runs/2, 0)
	}
}

func BenchmarkFibber_f4o0(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Run(4, runs/4, 0)
	}
}

func BenchmarkFibber_f8o0(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Run(8, runs/8, 0)
	}
}

func BenchmarkFibber_f24o0(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Run(24, runs/24, 0)
	}
}

func BenchmarkFibber_f32o0(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Run(32, runs/32, 0)
	}
}

// 2 overlap
func BenchmarkFibber_f1o2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Run(1, runs, 2)
	}
}

func BenchmarkFibber_f2o2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Run(2, runs/2, 2)
	}
}

func BenchmarkFibber_f4o2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Run(4, runs/4, 2)
	}
}

func BenchmarkFibber_f8o2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Run(8, runs/8, 2)
	}
}

func BenchmarkFibber_f24o2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Run(24, runs/24, 2)
	}
}

func BenchmarkFibber_f32o2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Run(32, runs/32, 2)
	}
}

// 4 overlap
func BenchmarkFibber_f1o4(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Run(1, runs, 4)
	}
}

func BenchmarkFibber_f2o4(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Run(2, runs/2, 4)
	}
}

func BenchmarkFibber_f4o4(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Run(4, runs/4, 4)
	}
}

func BenchmarkFibber_f8o4(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Run(8, runs/8, 4)
	}
}

func BenchmarkFibber_f24o4(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Run(24, runs/24, 4)
	}
}

func BenchmarkFibber_f32o4(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Run(32, runs/32, 4)
	}
}
