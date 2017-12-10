package main

import (
	"fmt"
)

type GroupCollection map[fmt.Stringer][]Chunk


func (g *GroupCollection) addChunk(key fmt.Stringer, chunk *Chunk)  {
	(*g)[key] = append((*g)[key], *chunk)
}

func (g *GroupCollection) checkOrCreateKey(key fmt.Stringer) {
	_, ok := (*g)[key]
	if ok == false {
		(*g)[key] = []Chunk{}
	}
}

type Collection struct {
	time        GroupCollection
	ip          GroupCollection
	token       GroupCollection
	chunkType   GroupCollection
	application GroupCollection
}

type each interface {
	EachCollection(func(interface{}, *[]Chunk) bool)
}



func NewCollection() Collection {
	return Collection{
		time:        GroupCollection{},
		ip:          GroupCollection{},
		token:       GroupCollection{},
		chunkType:   GroupCollection{},
		application: GroupCollection{},
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

func (c *Collection) addChunkToken(chunk *Chunk) {
	chunkToken := chunk.token
	c.token.checkOrCreateKey(chunkToken)
	c.token.addChunk(chunkToken, chunk)
}

func (c *Collection) addChunkApplication(chunk *Chunk) {
	chunkApplication := chunk.application
	c.application.checkOrCreateKey(chunkApplication)
	c.application.addChunk(chunkApplication, chunk)
}

func (c *Collection) addChunkTime(chunk *Chunk) {
	chunkTime := chunk.time
	c.time.checkOrCreateKey(chunkTime)
	c.time.addChunk(chunkTime, chunk)
}

func (c *Collection) addChunkType(chunk *Chunk) {
	chunkType := chunk.chunkType
	c.chunkType.checkOrCreateKey(chunkType)
	c.chunkType.addChunk(chunkType, chunk)
}

func (c *Collection) getGroup(key string)  {

}

func (g *GroupCollection) EachCollection(f func(key fmt.Stringer, c *[]Chunk) bool) {
	for key, chunkList := range *g {
		if f(key, &chunkList) == false {
			break
		}
	}
}
