package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"sort"
	"strings"
	"sync"

	"github.com/ishantsagar/sequencer/sequence"
)

var wg sync.WaitGroup

type Parser struct {
	data string
}

func main() {

	files := os.Args[1:]
	if len(files) > 0 {
		for _, f := range files {
			wg.Add(1)
			p := new(Parser)
			go p.parser(f)
		}
		wg.Wait()
	}
}

func (p *Parser) parser(file string) {
	if fileContent, err := ioutil.ReadFile(file); err != nil {
		panic(err)
	} else {
		p.data = string(fileContent)
		p.replace().standardizeSpaces().applyRegex()
		arr := strings.Split(p.data, " ")
		m := p.createMap(arr)
		pl := p.rankByWordCount(m)
		for _, val := range pl {
			fmt.Println(val)
		}
		wg.Done()
	}
}

func (p *Parser) replace() *Parser {
	p.data = strings.Replace(p.data, "\n", " ", -1)
	return p
}

func (p *Parser) applyRegex() *Parser {
	p.data = regexp.MustCompile(`[^a-zA-Z0-9 ]+`).ReplaceAllString(p.data, "")
	return p
}

func (p *Parser) standardizeSpaces() *Parser {
	p.data = strings.Join(strings.Fields(p.data), " ")
	return p
}

func (p *Parser) createMap(arr []string) (m map[string]int32) {
	m = make(map[string]int32)

	for i := 0; i < len(arr); i += 3 {
		if i+2 < len(arr) {
			s := fmt.Sprintf("%s %s %s", arr[i], arr[i+1], arr[i+2])
			if val, ok := m[s]; ok {
				m[s] = val + 1
			} else {
				m[s] = 1
			}
		}
	}
	return
}

func (p *Parser) rankByWordCount(frequency map[string]int32) sequence.SequenceList {
	pl := make(sequence.SequenceList, len(frequency))
	i := 0
	for k, v := range frequency {
		pl[i] = sequence.Sequence{Word: k, Occurrence: v}
		i++
	}
	sort.Sort(sort.Reverse(pl))
	return pl[0:100]
}
