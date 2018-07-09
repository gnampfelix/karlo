package main

import (
	"fmt"
	"io/ioutil"
	"time"
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

func main() {
	file := NewInputFromFile("main.go")
	for file.HasNext() {
		fmt.Printf("%s", string(file.GetNext()))
		time.Sleep(200 * time.Millisecond)
	}
}
