package main

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParserRemoveNewLine(t *testing.T) {
	p := new(Parser)
	p.data = "this is a test \n message"
	p.replace().standardizeSpaces().applyRegex()
	assert.Equal(t, "this is a test message", p.data)
}

func TestParseCalculateSequence(t *testing.T) {
	p := new(Parser)
	p.data = "check word sequence \n check word sequence"
	p.replace().standardizeSpaces().applyRegex()
	arr := strings.Split(p.data, " ")
	m := p.createMap(arr)
	assert.Equal(t, int32(2), m["check word sequence"])
}
