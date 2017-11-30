package main

import (
	"time"
)

type StringGroupCollection map[string][]Chunk
type TimeGroupCollection map[time.Time][]Chunk
type TypeGroupCollection map[ChunkType][]Chunk

type Collection struct {
	time        TimeGroupCollection
	ip          StringGroupCollection
	token       StringGroupCollection
	chunkType   TypeGroupCollection
	application StringGroupCollection
}

func NewCollection() Collection {
	return Collection{
		time:        TimeGroupCollection{},
		ip:          StringGroupCollection{},
		token:       StringGroupCollection{},
		chunkType:   TypeGroupCollection{},
		application: StringGroupCollection{},
	}
}

func (c *Collection) AddChunk(chunk *Chunk) {
	c.addChunkIp(chunk)
	c.addChunkToken(chunk)
	c.addChunkApplication(chunk)
	c.addChunkTime(chunk)
	c.addChunkType(chunk)
}

func (c *Collection) addChunkIp(chunk *Chunk) {
	chunkIp := chunk.ip
	c.ip.checkOrCreateKey(chunkIp)
	c.ip.addChunk(chunkIp, chunk)
}

func (c Collection) addChunkToken(chunk *Chunk) {
	chunkToken := chunk.token
	c.token.checkOrCreateKey(chunkToken)
	c.token.addChunk(chunkToken, chunk)
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
}
