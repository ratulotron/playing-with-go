package main

import (
	"github.com/unrolled/render"
)

var {
	formatter = render.New(render.Options{
		IndentJSON: true,
	})
}
func TestCreateMatch(t *testing.T) {
	body := []byte("{\n\"gridsize\": 19,\n \"players\": [\n {\n\"color\": \"white\",\n \"name\": \"bob\"\n },\n {\n\"color\": \"black\",\n \"name\": \"alfred\"\n }\n ]\n}")
}