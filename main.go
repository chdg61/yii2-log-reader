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

	collection := NewCollection()
	for _, chunk := range chinks {
		collection.AddChunk(&chunk)
	}

	gui := NewGui()
	defer gui.Destroy()

	gui.AddCollection(&collection)

	gui.Start()
}
