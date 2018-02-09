package main

import (
	"io/ioutil"

	"github.com/TruthHun/html2md"
)

func main() {
	b, _ := ioutil.ReadFile("gitbook.html")
	md := html2md.Convert(string(b))
	ioutil.WriteFile("gitbook.md", []byte(md), 0777)

	b, _ = ioutil.ReadFile("github.io.html")
	md = html2md.Convert(string(b))
	ioutil.WriteFile("github.io.md", []byte(md), 0777)

	b, _ = ioutil.ReadFile("readthedoc.html")
	md = html2md.Convert(string(b))
	ioutil.WriteFile("readthedoc.md", []byte(md), 0777)
}
