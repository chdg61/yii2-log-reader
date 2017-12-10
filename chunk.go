package main

import (
	"time"
	//"fmt"
	"strings"
	"regexp"
	"strconv"
)

type ChunkType string

func (c ChunkType) String() string  {
	return strings.ToUpper(string(c))
}

type Ip string

func (ip Ip) String() string {
	return string(ip)
}

type Token string

func (t Token) String() string {
	return string(t)
}

type Application string

func (a Application) String() string {
	return string(a)
}

type Time struct {
	time.Time
}

func (t Time) String() string {
	return t.Format("02.01.2006 15:04:05")
}

func NewTime(year int, month time.Month, day, hour, min, sec int) Time {
	t := Time{}
	t.Time = time.Date(year, month, day, hour, min, sec, 0, time.UTC)

	return t
}


type Chunk struct {
	time Time
	ip Ip
	token Token
	chunkType ChunkType
	application Application
	message string
	text string
}

func (c *Chunk) addText(text string)  {
	separator := "";
	if len(c.text) > 0 {
		separator = "\n"
	}
	c.text = strings.Join([]string{c.text, text},separator)
}

func Parse(args []byte) []Chunk {
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
				time: NewTime(year, time.Month(month), day, hour, minutes, seconds),
				ip: Ip(matches[2]),
				token: Token(matches[4]),
				application: Application(matches[6]),
				chunkType: ChunkType(matches[5]),
			}
			chunk.addText(matches[7])
			result = append(result, chunk)
			number++
			//fmt.Printf("%#v\n",chunk)
			//fmt.Println(matches[3])
			//fmt.Println("--------------")
		} else if number > 0 {
			chunk := &result[number-1]
			chunk.addText(value)
		}
	}

	return result
}