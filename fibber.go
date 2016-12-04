package main

import (
	"flag"
	"fmt"
	"sync"
	"unsafe"
)

const padSize = 0
const fibberCount = 16

type fibber struct {
	last int
	val  int
	sync.Mutex
	pad [padSize]int
}

func (f *fibber) fibon() {
	f.Lock()
	defer f.Unlock()
	if f.val == 0 {
		f.val = 1
	} else if f.last == 0 {
		f.last = 1
	} else {
		last := f.val
		f.val = f.val + f.last
		f.last = last
	}
}

func (f *fibber) String() string {
	return fmt.Sprintf("%v", f.val)
}

func describe(fibbers [fibberCount]fibber) {
	fmt.Printf("Addresses of our fibbers (%d bytes):\n", unsafe.Sizeof(fibber{}))
	var plast unsafe.Pointer
	for i := range fibbers {
		p := unsafe.Pointer(&fibbers[i])
		if plast != nil {
			fmt.Printf("+ %d", uintptr(p)-uintptr(plast))
		} else {
			fmt.Printf("+ 0")
		}
		plast = p
		fmt.Printf("\t0x%x\n", p)
	}
}

func fibberize(fib *fibber, runs int, w *sync.WaitGroup) {
	for i := 0; i < runs; i++ {
		fib.fibon()
		// fmt.Printf("%s\n", fib)
	}
	w.Done()
}

func main() {
	var fibbers [fibberCount]fibber

	describe(fibbers)

	funcs := flag.Int("f", 0, "num gofunc")
	runs := flag.Int("r", 0, "num runs per")
	overlap := flag.Int("o", 0, "overlap count")
	flag.Parse()

	var w sync.WaitGroup
	o := 0
	for i := 0; i < *funcs; i++ {
		w.Add(1)
		go fibberize(&fibbers[o], *runs, &w)
		o++
		if o == *overlap {
			o = 0
		}
	}

	w.Wait()
}
