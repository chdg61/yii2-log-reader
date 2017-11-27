package main

import (
	"time"
)

type StringGroupCollections map[string][]Chunk
type TimeGroupCollections map[time.Time][]Chunk
type TypeGroupCollections map[ChunkType][]Chunk

type Collections struct {
	time TimeGroupCollections
	ip StringGroupCollections
	token StringGroupCollections
	chunkType TypeGroupCollections
	application StringGroupCollections
}

func CreateEmptyCollections() Collections {
	return Collections{
		time: TimeGroupCollections{},
		ip: StringGroupCollections{},
		token: StringGroupCollections{},
		chunkType: TypeGroupCollections{},
		application: StringGroupCollections{},
	}
}

func (c *Collections) AddChunk(chunk *Chunk) {
	c.addChunkIp(chunk)
	c.addChunkToken(chunk)
	c.addChunkApplication(chunk)
	c.addChunkTime(chunk)
	c.addChunkType(chunk)
}

func (c *Collections) addChunkIp(chunk *Chunk) {
	chunkIp := chunk.ip
	c.ip.checkOrCreateKey(chunkIp)
	c.ip.addChunk(chunkIp, chunk)
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
}
