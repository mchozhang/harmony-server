package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHandler(t *testing.T)  {
	expectResponseBody :="{\"data\":" +
		"{\"level\":{\"cells\":[[{\"col\":0,\"row\":0,\"steps\":4,\"targetRow\":3}," +
		"{\"col\":1,\"row\":0,\"steps\":1,\"targetRow\":0}," +
		"{\"col\":2,\"row\":0,\"steps\":3,\"targetRow\":1}," +
		"{\"col\":3,\"row\":0,\"steps\":1,\"targetRow\":0}," +
		"{\"col\":4,\"row\":0,\"steps\":2,\"targetRow\":4}]," +
		"[{\"col\":0,\"row\":1,\"steps\":3,\"targetRow\":4}," +
		"{\"col\":1,\"row\":1,\"steps\":1,\"targetRow\":1}," +
		"{\"col\":2,\"row\":1,\"steps\":1,\"targetRow\":1}," +
		"{\"col\":3,\"row\":1,\"steps\":3,\"targetRow\":2}," +
		"{\"col\":4,\"row\":1,\"steps\":2,\"targetRow\":0}]," +
		"[{\"col\":0,\"row\":2,\"steps\":2,\"targetRow\":2}," +
		"{\"col\":1,\"row\":2,\"steps\":2,\"targetRow\":3}," +
		"{\"col\":2,\"row\":2,\"steps\":2,\"targetRow\":0}," +
		"{\"col\":3,\"row\":2,\"steps\":3,\"targetRow\":2}," +
		"{\"col\":4,\"row\":2,\"steps\":3,\"targetRow\":2}]," +
		"[{\"col\":0,\"row\":3,\"steps\":4,\"targetRow\":3}," +
		"{\"col\":1,\"row\":3,\"steps\":2,\"targetRow\":3}," +
		"{\"col\":2,\"row\":3,\"steps\":3,\"targetRow\":4}," +
		"{\"col\":3,\"row\":3,\"steps\":3,\"targetRow\":4}," +
		"{\"col\":4,\"row\":3,\"steps\":2,\"targetRow\":2}]," +
		"[{\"col\":0,\"row\":4,\"steps\":3,\"targetRow\":3}," +
		"{\"col\":1,\"row\":4,\"steps\":2,\"targetRow\":0}," +
		"{\"col\":2,\"row\":4,\"steps\":3,\"targetRow\":4}," +
		"{\"col\":3,\"row\":4,\"steps\":2,\"targetRow\":1}," +
		"{\"col\":4,\"row\":4,\"steps\":1,\"targetRow\":1}]]," +
		"\"colors\":[\"#C62828\",\"#7C0718\",\"#33691E\",\"#43A047\",\"#F0F4C3\"],\"size\":5}}}"

	req := Request{Body: "{\"query\":\"query {level(id: 25) {size colors\\n cells{\\n targetRow\\n steps\\n row\\n col}}}\", \"variables\":{}}"}
	res, err := Handler(req)
	assert.Nil(t, err)
	assert.Equal(t, 200, res.StatusCode)
	assert.Equal(t, expectResponseBody, res.Body)
}
