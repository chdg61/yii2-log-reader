package main

import (
	"testing"
	"fmt"
)

func TestChunkType(t *testing.T)  {
	chunk := ChunkType("lower text")
	if chunk.String() != "LOWER TEXT" {
		t.Error("Error convert Uppercase")
	}
}

func TestChunkAddText(t *testing.T)  {
	chunk := Chunk{}

	chunk.addText("first")

	if chunk.text != "first" {
		t.Error("Error add text chunk")
	}

	chunk.addText("second")

	if chunk.text != "first\nsecond" {
		t.Error("Error add next text chunk")
	}
}

func TestNewTime(t *testing.T) {
	time := NewTime(2018,1,1,6,20,10)
	if time.Day() != 1 {
		t.Error("Day time error");
	}
	if time.Month() != 1 {
		t.Error("Month time error");
	}
	if time.Year() != 2018 {
		t.Error("Year time error");
	}
	if time.Hour() != 6 {
		t.Error("Hour time error");
	}
	if time.Minute() != 20 {
		t.Error("Minute time error");
	}
	if time.Second() != 10 {
		t.Error("Second time error");
	}
	if time.String() != "01.01.2018 06:20:10" {
		t.Errorf("Time not format %s", time.String())
	}
}

func BenchmarkHello(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fmt.Sprintf("hello")
	}
}

func TestParse(t *testing.T) {
	var text []byte = []byte(`2018-01-01 14:15:16 [192.168.1.1][23][test_hash][error][api\exception\ApiException] api\exception\ApiException: Не удалось получить список заказов in /www/vendor/api.php
	Stack trace:
	#0 /www/controller.php: api->getOrderList()
	#1 {main}
	2018-01-01 14:15:17 [192.168.1.2][23][test_hash][info][application] $_COOKIE = [
	'id' => 23
	]`)

	chunks := Parse(text)

	if len(chunks) != 2 {
		t.Error("Count chunk is not 2")
	}

	chunk1 := &chunks[0]
	chunk2 := &chunks[1]

	if chunk1.ip != `192.168.1.1` {
		t.Errorf("Parser ip error %s", chunk1.ip)
	}

	if chunk2.ip != `192.168.1.2` {
		t.Errorf("Parser ip error %s", chunk2.ip)
	}

	if chunk1.time.String() != `01.01.2018 14:15:16` {
		t.Errorf("Parse time error %s", chunk1.time)
	}

	if chunk2.time.String() != `01.01.2018 14:15:17` {
		t.Errorf("Parse time error %s", chunk1.time)
	}

	if chunk1.token != `test_hash` {
		t.Errorf("Parse token error %s", chunk1.token)
	}

	if chunk2.token != `test_hash` {
		t.Errorf("Parse token error %s", chunk2.token)
	}

	if chunk1.chunkType != "error" {
		t.Errorf("Parse chunk type error %s", chunk1.chunkType)
	}

	if chunk2.chunkType != `info` {
		t.Errorf("Parse chunk type error %s", chunk2.chunkType)
	}

	if chunk1.application != `api\exception\ApiException` {
		t.Errorf("Parse application error %s", chunk1.application)
	}

	if chunk2.application != `application` {
		t.Errorf("Parse application error %s", chunk2.application)
	}

	mockText := `api\exception\ApiException: Не удалось получить список заказов in /www/vendor/api.php
	Stack trace:
	#0 /www/controller.php: api->getOrderList()
	#1 {main}`

	if chunk1.text != mockText {
		t.Errorf("Parse message error %s", chunk1.text)
	}

	mockText = `$_COOKIE = [
	'id' => 23
	]`

	if chunk2.text != mockText {
		t.Errorf("Parse message error %s", chunk2.text)
	}
}

func TestParseFail(t *testing.T) {
	var text []byte = []byte(``)

	chunks := Parse(text)

	if len(chunks) != 0 {
		t.Error("Parse fail error")
	}
}
