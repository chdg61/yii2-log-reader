package main

import (
	"testing"
)

func TestNewCollection(t *testing.T) {
	collection := NewCollection()

	if len(collection.ip) > 0 {
		t.Error("Not empty ip collection")
	}

	if len(collection.time) > 0 {
		t.Error("Not empty time collection")
	}

	if len(collection.chunkType) > 0 {
		t.Error("Not empty type collection")
	}

	if len(collection.application) > 0 {
		t.Error("Not empty application collection")
	}

	if len(collection.token) > 0 {
		t.Error("Not empty token collection")
	}
}

func TestAddChunkIp(t *testing.T) {
	collection := NewCollection()

	chunk1 := Chunk{ip: Ip("111")}
	chunk2 := Chunk{ip: Ip("111")}
	chunk3 := Chunk{ip: Ip("222")}

	collection.addChunkIp(&chunk1)
	collection.addChunkIp(&chunk2)
	collection.addChunkIp(&chunk3)

	if len(collection.ip) != 2 {
		t.Error("Error add chunk ip")
	}

	s, lenIp := collection.ip[Ip("111")]

	if lenIp == false {
		t.Error("Error len 111")
	}

	if len(s) != 2 {
		t.Error("Error count 111 element")
	}

	s, lenIp = collection.ip[Ip("222")]

	if lenIp == false {
		t.Error("Error len 222")
	}

	if len(s) != 1 {
		t.Error("Error count 222 element")
	}
}

func TestAddChunk(t *testing.T) {
	collection := NewCollection()

	chunk1 := Chunk{
		ip: Ip("ip_adres_1"),
		token: Token("token1"),
	}
	chunk2 := Chunk{
		ip: Ip("ip_adres_1"),
		token: Token("token1"),
	}
	chunk3 := Chunk{
		ip: Ip("ip_adres_2"),
		token: Token("token2"),
	}

	collection.AddChunk(&chunk1)
	collection.AddChunk(&chunk2)
	collection.AddChunk(&chunk3)

	if len(collection.ip) != 2 {
		t.Error("Error add chunk ip")
	}

	s, lenIp := collection.ip[Ip("ip_adres_1")]

	if lenIp == false {
		t.Error("Error len ip_adres_1")
	}

	if len(s) != 2 {
		t.Error("Error count ip_adres_1 element")
	}

	s, lenIp = collection.ip[Ip("ip_adres_2")]

	if lenIp == false {
		t.Error("Error len ip_adres_2")
	}

	if len(s) != 1 {
		t.Error("Error count ip_adres_2 element")
	}

	// token
	if len(collection.token) != 2 {
		t.Error("Error add chunk ip")
	}

	s, lenToken := collection.token[Token("token1")]

	if lenToken == false {
		t.Error("Error len token1")
	}

	if len(s) != 2 {
		t.Error("Error count token1 element")
	}

	s, lenToken = collection.token[Token("token2")]

	if lenToken == false {
		t.Error("Error len token1")
	}

	if len(s) != 1 {
		t.Error("Error count token1 element")
	}
}

/*func (c *Collection) AddChunk(chunk *Chunk) {
	c.addChunkIp(chunk)
	c.addChunkToken(chunk)
	c.addChunkApplication(chunk)
	c.addChunkTime(chunk)
	c.addChunkType(chunk)
}


func (c Collection) addChunkApplication(chunk *Chunk) {
	chunkApplication := chunk.application
	c.application.checkOrCreateKey(chunkApplication)
	c.application.addChunk(chunkApplication, chunk)
}

func (c Collection) addChunkTime(chunk *Chunk) {
	chunkTime := chunk.time
	c.time.checkOrCreateKey(chunkTime)
	c.time.addChunk(chunkTime, chunk)
}

func (c Collection) addChunkType(chunk *Chunk) {
	chunkType := chunk.chunkType
	c.chunkType.checkOrCreateKey(chunkType)
	c.chunkType.addChunk(chunkType, chunk)
}

func (s StringGroupCollection) addChunk(key string, chunk *Chunk)  {
	s[key] = append(s[key], *chunk)
}

func (s StringGroupCollection) checkOrCreateKey(key string) {
	_, ok := s[key]
	if ok == false {
		s[key] = []Chunk{}
	}
}

func (s TypeGroupCollection) addChunk(key ChunkType, chunk *Chunk)  {
	s[key] = append(s[key], *chunk)
}

func (s TypeGroupCollection) checkOrCreateKey(key ChunkType) {
	_, ok := s[key]
	if ok == false {
		s[key] = []Chunk{}
	}
}


func (s TimeGroupCollection) addChunk(key time.Time, chunk *Chunk)  {
	s[key] = append(s[key], *chunk)
}

func (s TimeGroupCollection) checkOrCreateKey(key time.Time) {
	_, ok := s[key]
	if ok == false {
		s[key] = []Chunk{}
	}
}*/
