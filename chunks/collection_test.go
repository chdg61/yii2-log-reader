package chunks

import (
	"testing"
)

func TestNewCollection(t *testing.T) {
	collection := NewCollection()

	if len(collection.Ip) > 0 {
		t.Error("Not empty ip collection")
	}

	if len(collection.Time) > 0 {
		t.Error("Not empty time collection")
	}

	if len(collection.ChunkType) > 0 {
		t.Error("Not empty type collection")
	}

	if len(collection.Application) > 0 {
		t.Error("Not empty application collection")
	}

	if len(collection.Token) > 0 {
		t.Error("Not empty token collection")
	}
}

func TestAddChunkIp(t *testing.T) {
	collection := NewCollection()

	chunk1 := Chunk{Ip: Ip("111")}
	chunk2 := Chunk{Ip: Ip("111")}
	chunk3 := Chunk{Ip: Ip("222")}

	collection.addChunkIp(&chunk1)
	collection.addChunkIp(&chunk2)
	collection.addChunkIp(&chunk3)

	if len(collection.Ip) != 2 {
		t.Error("Error add chunk ip")
	}

	s, lenIp := collection.Ip[Ip("111")]

	if lenIp == false {
		t.Error("Error len 111")
	}

	if len(s) != 2 {
		t.Error("Error count 111 element")
	}

	s, lenIp = collection.Ip[Ip("222")]

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
		Ip:    Ip("ip_adres_1"),
		Token: Token("token1"),
	}
	chunk2 := Chunk{
		Ip:    Ip("ip_adres_1"),
		Token: Token("token1"),
	}
	chunk3 := Chunk{
		Ip:    Ip("ip_adres_2"),
		Token: Token("token2"),
	}

	collection.AddChunk(&chunk1)
	collection.AddChunk(&chunk2)
	collection.AddChunk(&chunk3)

	if len(collection.Ip) != 2 {
		t.Error("Error add chunk ip")
	}

	s, lenIp := collection.Ip[Ip("ip_adres_1")]

	if lenIp == false {
		t.Error("Error len ip_adres_1")
	}

	if len(s) != 2 {
		t.Error("Error count ip_adres_1 element")
	}

	s, lenIp = collection.Ip[Ip("ip_adres_2")]

	if lenIp == false {
		t.Error("Error len ip_adres_2")
	}

	if len(s) != 1 {
		t.Error("Error count ip_adres_2 element")
	}

	// token
	if len(collection.Token) != 2 {
		t.Error("Error add chunk ip")
	}

	s, lenToken := collection.Token[Token("token1")]

	if lenToken == false {
		t.Error("Error len token1")
	}

	if len(s) != 2 {
		t.Error("Error count token1 element")
	}

	s, lenToken = collection.Token[Token("token2")]

	if lenToken == false {
		t.Error("Error len token1")
	}

	if len(s) != 1 {
		t.Error("Error count token1 element")
	}
}
