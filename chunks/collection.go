package chunks

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
	Time        GroupCollection
	Ip          GroupCollection
	Token       GroupCollection
	ChunkType   GroupCollection
	Application GroupCollection
}

type each interface {
	EachCollection(func(interface{}, *[]Chunk) bool)
}



func NewCollection() Collection {
	return Collection{
		Time:        GroupCollection{},
		Ip:          GroupCollection{},
		Token:       GroupCollection{},
		ChunkType:   GroupCollection{},
		Application: GroupCollection{},
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
	chunkIp := chunk.Ip
	c.Ip.checkOrCreateKey(chunkIp)
	c.Ip.addChunk(chunkIp, chunk)
}

func (c *Collection) addChunkToken(chunk *Chunk) {
	chunkToken := chunk.Token
	c.Token.checkOrCreateKey(chunkToken)
	c.Token.addChunk(chunkToken, chunk)
}

func (c *Collection) addChunkApplication(chunk *Chunk) {
	chunkApplication := chunk.Application
	c.Application.checkOrCreateKey(chunkApplication)
	c.Application.addChunk(chunkApplication, chunk)
}

func (c *Collection) addChunkTime(chunk *Chunk) {
	chunkTime := chunk.Time
	c.Time.checkOrCreateKey(chunkTime)
	c.Time.addChunk(chunkTime, chunk)
}

func (c *Collection) addChunkType(chunk *Chunk) {
	chunkType := chunk.ChunkType
	c.ChunkType.checkOrCreateKey(chunkType)
	c.ChunkType.addChunk(chunkType, chunk)
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
