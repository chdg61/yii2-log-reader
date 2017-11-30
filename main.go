package main

import (
	"flag"
	"io/ioutil"
)


func main() {
	flag.Parse()
	var fileName = flag.Arg(0)

	var file, err = ioutil.ReadFile(fileName)

	if err != nil {
		panic(err)
	}


	chinks := Parse(file)

	collections := CreateEmptyCollections()
	for _, chunk := range chinks {
		collections.AddChunk(&chunk)
	}

	gui := NewGui()
	gui.Start()
}
