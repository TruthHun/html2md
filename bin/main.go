package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/TruthHun/html2md"
	"github.com/astaxie/beego/logs"
)

func main() {
	args := os.Args[1:]
	if len(args) != 2 {
		usg := `
Usage:
	html2md input_html_file output_markdown_file

For example:
	html2md input.html output.md
		`
		fmt.Println(usg)
	} else {
		if b, err := ioutil.ReadFile(args[0]); err != nil {
			logs.Error(err)
		} else {
			md := html2md.Convert(string(b))
			if err := ioutil.WriteFile(args[1], []byte(md), 0777); err != nil {
				logs.Error(err)
			} else {
				logs.Info("success")
			}
		}

	}
}
