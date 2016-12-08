#!/bin/bash

run() {
	s=${1}
	echo "size: $s"
	SIZE=${s} make -s size 
	go test -bench BenchmarkFibber_f8o0
}

run 0
for s in $(seq 5 8 1024); do
	run "${s}"
done
