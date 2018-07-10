package main

import (
	"fmt"
	"io/ioutil"
	"time"
	"math/rand"
)

type Input struct {
	data []byte
	current int
}

func NewInputFromFile(filename string) Input {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return Input{
			data: make([]byte, 0),
			current: 0,
		}
	} else {
		return Input {
			data: data,
			current: 0,
		}
	}
}

func (i Input)HasNext() bool {
	return i.current < len(i.data)
}

func (i *Input) GetNext() byte {
	if i.HasNext() {
		result := i.data[i.current]
		i.current++
		return result
	} else {
		return byte(0)
	}
}

type Ticker struct {
	hitsPerSecond int
	maxDeviationInMs int
	lastTick time.Time
}

func NewTicker(hitsPerSecond, maxDeviationInMs int) Ticker {
	return Ticker{
		hitsPerSecond: hitsPerSecond,
		maxDeviationInMs: maxDeviationInMs,
		lastTick: time.Now(),
	}
}

func (t *Ticker) Wait() {
	nextRegularTick := t.lastTick.Add(time.Duration(60000 / t.hitsPerSecond) * time.Millisecond)
	nextSimulatedTick := nextRegularTick.Add(time.Duration(rand.Int31n(int32(t.maxDeviationInMs))) * time.Millisecond)
	time.Sleep(time.Until(nextSimulatedTick))
	t.lastTick = time.Now()
}

func main() {
	file := NewInputFromFile("main.go")
	ticker := NewTicker(4000, 200)
	for file.HasNext() {
		fmt.Printf("%s", string(file.GetNext()))
		ticker.Wait()
	}
}
