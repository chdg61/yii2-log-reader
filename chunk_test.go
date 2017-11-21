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

func BenchmarkHello(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fmt.Sprintf("hello")
	}
}


/*

type Chunk struct {
	time time.Time
	ip string
	token string
	chunkType ChunkType
	application string
	message string
	text string
}

func (c *Chunk) addText(text string)  {
	c.text = strings.Join([]string{c.text, text},"\n")
}*/
/*

func Parse(args []byte) {
	str := string(args)
	strList := strings.Split(str, "\n")

	number := 0

	result := []Chunk{}

	//r, _ := regexp.Compile("\\d{4}\\s[\\d:]\\s*")
	//r, _ := regexp.Compile("\\d{4}-\\d{1,2}-\\d{1,2}\\s[\\d:]*\\s\\[.*?\\]")

	for _, value := range strList {

		r, regexpErr := regexp.Compile(GetInstant().RegexpCheck)

		if regexpErr != nil {
			panic("Wrong regexp_check settings")
		}

		//math, _ := regexp.MatchString(GetInstant().RegexpCheck, value)
		math:= r.MatchString(value)

		if math {
			r, _ := regexp.Compile(GetInstant().RegexpHeader)
			matches := r.FindStringSubmatch(value)

			r, _ = regexp.Compile("(\\d{4})-(\\d{1,2})-(\\d{1,2})\\s(\\d{1,2}):(\\d{1,2}):(\\d{1,2})")
			matchesTime := r.FindStringSubmatch(matches[1])

			year, _ := strconv.Atoi(matchesTime[1])
			month, _ := strconv.Atoi(matchesTime[2])
			day, _ := strconv.Atoi(matchesTime[3])
			hour, _ := strconv.Atoi(matchesTime[4])
			minutes, _ := strconv.Atoi(matchesTime[5])
			seconds, _ := strconv.Atoi(matchesTime[6])

			chunk := Chunk{
				time: time.Date(year, time.Month(month), day, hour, minutes, seconds, 0, time.UTC),
				ip: matches[2],
				token: matches[4],
				application: matches[6],
				chunkType: ChunkType(matches[5]),
			}
			chunk.addText(matches[7])
			result = append(result, chunk)
			number++
			fmt.Printf("%#v",chunk)
			fmt.Println(matches[3])
			fmt.Println("--------------")
		} else if number > 0 {
			chunk := &result[number-1]
			chunk.addText(value)
		}
	}

	fmt.Printf("%+v", result)
}*/
