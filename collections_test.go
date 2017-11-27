package main

import (
	"testing"
)

func TestCreateEmptyCollections(t *testing.T) {
	collections := CreateEmptyCollections();

	if len(collections.ip) > 0 {
		t.Error("Not empty ip collection")
	}

	if len(collections.time) > 0 {
		t.Error("Not empty time collection")
	}

	if len(collections.chunkType) > 0 {
		t.Error("Not empty type collection")
	}

	if len(collections.application) > 0 {
		t.Error("Not empty application collection")
	}

	if len(collections.token) > 0 {
		t.Error("Not empty token collection")
	}
}

func TestAddChunkIp(t *testing.T) {
	collections := CreateEmptyCollections()

	chunk1 := Chunk{ip: "111"}
	chunk2 := Chunk{ip: "111"}
	chunk3 := Chunk{ip: "222"}

	collections.addChunkIp(&chunk1)
	collections.addChunkIp(&chunk2)
	collections.addChunkIp(&chunk3)

	if len(collections.ip) != 2 {
		t.Error("Error add chunk ip")
	}

	s, lenIp := collections.ip["111"]

	if lenIp == false {
		t.Error("Error len 111")
	}

	if len(s) != 2 {
		t.Error("Error count 111 element")
	}

	s, lenIp = collections.ip["222"]

	if lenIp == false {
		t.Error("Error len 222")
	}

	if len(s) != 1 {
		t.Error("Error count 222 element")
	}
}

/*func (c *Collections) AddChunk(chunk *Chunk) {
	c.addChunkIp(chunk)
	c.addChunkToken(chunk)
	c.addChunkApplication(chunk)
	c.addChunkTime(chunk)
	c.addChunkType(chunk)
}


func (c Collections) addChunkToken(chunk *Chunk) {
	chunkToken := chunk.token
	c.token.checkOrCreateKey(chunkToken)
	c.token.addChunk(chunkToken, chunk)
}

func (c Collections) addChunkApplication(chunk *Chunk) {
	chunkApplication := chunk.application
	c.application.checkOrCreateKey(chunkApplication)
	c.application.addChunk(chunkApplication, chunk)
}

func (c Collections) addChunkTime(chunk *Chunk) {
	chunkTime := chunk.time
	c.time.checkOrCreateKey(chunkTime)
	c.time.addChunk(chunkTime, chunk)
}

func (c Collections) addChunkType(chunk *Chunk) {
	chunkType := chunk.chunkType
	c.chunkType.checkOrCreateKey(chunkType)
	c.chunkType.addChunk(chunkType, chunk)
}

func (s StringGroupCollections) addChunk(key string, chunk *Chunk)  {
	s[key] = append(s[key], *chunk)
}

func (s StringGroupCollections) checkOrCreateKey(key string) {
	_, ok := s[key]
	if ok == false {
		s[key] = []Chunk{}
	}
}

func (s TypeGroupCollections) addChunk(key ChunkType, chunk *Chunk)  {
	s[key] = append(s[key], *chunk)
}

func (s TypeGroupCollections) checkOrCreateKey(key ChunkType) {
	_, ok := s[key]
	if ok == false {
		s[key] = []Chunk{}
	}
}


func (s TimeGroupCollections) addChunk(key time.Time, chunk *Chunk)  {
	s[key] = append(s[key], *chunk)
}

func (s TimeGroupCollections) checkOrCreateKey(key time.Time) {
	_, ok := s[key]
	if ok == false {
		s[key] = []Chunk{}
	}
}*/
