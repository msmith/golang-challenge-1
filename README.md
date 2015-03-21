# Go Challenge 1

A solution to [Go Challenge 1](http://golang-challenge.com/go-challenge1/)

Michael Smith
https://github.com/msmith

## Running the tests

```
$ go test drum
ok  	drum	0.004s
```

## Running the benchmarks

```
$ go test -bench=. -benchmem drum
PASS
Benchmark1	  300000	      4893 ns/op	     712 B/op	      35 allocs/op
Benchmark2	  500000	      3594 ns/op	     504 B/op	      26 allocs/op
Benchmark3	  300000	      4918 ns/op	     712 B/op	      35 allocs/op
Benchmark4	  500000	      3690 ns/op	     512 B/op	      26 allocs/op
Benchmark5	 1000000	      2313 ns/op	     328 B/op	      17 allocs/op
ok  	drum	9.104s
```

## Building & running the decoder

```
$ go build decode.go
$ ./decode src/drum/fixtures/*.splice
src/drum/fixtures/pattern_1.splice:

Saved with HW Version: 0.808-alpha
Tempo: 120
(0) kick        |x---|x---|x---|x---|
(1) snare       |----|x---|----|x---|
(2) clap        |----|x-x-|----|----|
(3) hh-open     |--x-|--x-|x-x-|--x-|
(4) hh-close    |x---|x---|----|x--x|
(5) cowbell     |----|----|--x-|----|


src/drum/fixtures/pattern_2.splice:

Saved with HW Version: 0.808-alpha
Tempo: 98.4
(0) kick        |x---|----|x---|----|
(1) snare       |----|x---|----|x---|
(3) hh-open     |--x-|--x-|x-x-|--x-|
(5) cowbell     |----|----|x---|----|


src/drum/fixtures/pattern_3.splice:

Saved with HW Version: 0.808-alpha
Tempo: 118
(40) kick       |x---|----|x---|----|
(1) clap        |----|x---|----|x---|
(3) hh-open     |--x-|--x-|x-x-|--x-|
(5) low-tom     |----|---x|----|----|
(12) mid-tom    |----|----|x---|----|
(9) hi-tom      |----|----|-x--|----|


src/drum/fixtures/pattern_4.splice:

Saved with HW Version: 0.909
Tempo: 240
(0) SubKick     |----|----|----|----|
(1) Kick        |x---|----|x---|----|
(99) Maracas    |x-x-|x-x-|x-x-|x-x-|
(255) Low Conga |----|x---|----|x---|


src/drum/fixtures/pattern_5.splice:

Saved with HW Version: 0.708-alpha
Tempo: 999
(1) Kick        |x---|----|x---|----|
(2) HiHat       |x-x-|x-x-|x-x-|x-x-|
```
