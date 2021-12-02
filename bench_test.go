package main

import (
	"fmt"
	"os"
	"testing"
	"uk.ac.bris.cs/gameoflife/gol"
)

func BenchmarkGol(b *testing.B) {
	os.Stdout = nil
	//for threads := 1; threads <= 16; threads++ {
	p := gol.Params{
		Turns:       100,
		Threads:     1,
		ImageWidth:  16,
		ImageHeight: 16,
	}
	testName := fmt.Sprintf("%dx%dx%d-%d", p.ImageWidth, p.ImageHeight, p.Turns, p.Threads)
	b.Run(testName, func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			events := make(chan gol.Event)
			go gol.Run(p, events, nil)
			for range events {
			}
		}
	})

	p = gol.Params{
		Turns:       100,
		Threads:     1,
		ImageWidth:  64,
		ImageHeight: 64,
	}
	testName = fmt.Sprintf("%dx%dx%d-%d", p.ImageWidth, p.ImageHeight, p.Turns, p.Threads)
	b.Run(testName, func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			events := make(chan gol.Event)
			go gol.Run(p, events, nil)
			for range events {
			}
		}
	})

	p = gol.Params{
		Turns:       100,
		Threads:     1,
		ImageWidth:  128,
		ImageHeight: 128,
	}
	testName = fmt.Sprintf("%dx%dx%d-%d", p.ImageWidth, p.ImageHeight, p.Turns, p.Threads)
	b.Run(testName, func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			events := make(chan gol.Event)
			go gol.Run(p, events, nil)
			for range events {
			}
		}
	})

	p = gol.Params{
		Turns:       100,
		Threads:     1,
		ImageWidth:  256,
		ImageHeight: 256,
	}
	testName = fmt.Sprintf("%dx%dx%d-%d", p.ImageWidth, p.ImageHeight, p.Turns, p.Threads)
	b.Run(testName, func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			events := make(chan gol.Event)
			go gol.Run(p, events, nil)
			for range events {
			}
		}
	})

	p = gol.Params{
		Turns:       100,
		Threads:     1,
		ImageWidth:  512,
		ImageHeight: 512,
	}
	testName = fmt.Sprintf("%dx%dx%d-%d", p.ImageWidth, p.ImageHeight, p.Turns, p.Threads)
	b.Run(testName, func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			events := make(chan gol.Event)
			go gol.Run(p, events, nil)
			for range events {
			}
		}
	})

	//}
}
