package main

import (
	"flag"
	"io/ioutil"
	"github.com/chdg61/yii2-log-reader/chunks"
	"github.com/chdg61/yii2-log-reader/ui"
)


func main() {
	flag.Parse()
	var fileName = flag.Arg(0)

	var file, err = ioutil.ReadFile(fileName)

	if err != nil {
		panic(err)
	}


	chinks := chunks.Parse(file)

	collection := chunks.NewCollection()
	for _, chunk := range chinks {
		collection.AddChunk(&chunk)
	}

	uiImpl := ui.NewUI()
	defer uiImpl.Destroy()

	uiImpl.AddCollection(&collection)

	uiImpl.Start()
}
