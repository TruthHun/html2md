//Author:TruthHun
//Email: TruthHun@QQ.COM
//Date:  2018-02-03
package converter

import (
	"strings"

	"fmt"

	"io/ioutil"

	"github.com/PuerkitoBio/goquery"
)

var tags = map[string]string{
	//"hr":    "- - -\n\n",
	//"br":    "\n",
	//"title": "# ",
	//"h1":    "# ",
	//"h2":    "## ",
	//"h3":    "### ",
	//"h4":    "#### ",
	//"h5":    "##### ",
	//"h6":    "###### ",
	//
	//"del":        "~~",
	//"b":          "**",
	//"strong":     "**",
	//"i":          "_",
	//"em":         "_",
	//"dfn":        "_",
	//"var":        "_",
	//"cite":       "_",
	//"figure": "\n",

	"ul":         "* ",
	"ol":         "1. ",
	"li":         "1. ",
	"dl":         "- ",
	"dd":         "- ",
	"pre":        "```",
	"blockquote": "> ", //最后要加个换行
	"code":       "",
	"span":       "",
	"table":      "",
	"thead":      "",
	"tbody":      "",
	"th":         "",
	"td":         "",
	"tr":         "",
}

var closeTag = map[string]string{
	"del":     "~~",
	"b":       "**",
	"strong":  "**",
	"i":       "_",
	"em":      "_",
	"dfn":     "_",
	"var":     "_",
	"cite":    "_",
	"figure":  "\n",
	"p":       "\n",
	"div":     "\n",
	"article": "\n",
	"nav":     "\n",
	"footer":  "\n",
	"header":  "\n",
	"br":      "\n",
	"span":    "",
}
var nextLineTag = []string{
	"pre", "blockquote", "table",
}

func Convert(htmlstr string) {
	//htmlstr = htmltil.Compress(htmlstr)
	doc, _ := goquery.NewDocumentFromReader(strings.NewReader(htmlstr))
	doc = handleA(doc)
	doc = handleImg(doc)
	doc = handleHead(doc)
	doc = handleClosedTag(doc)
	doc = handleHr(doc)
	doc = handleNextLineTag(doc)
	doc = handleLi(doc)
	doc = handleTable(doc)
	h, _ := doc.Find("body").Html()
	ioutil.WriteFile("test.md", []byte(h), 0777)
}

//[ok]handle tag <a>
func handleA(doc *goquery.Document) *goquery.Document {
	doc.Find("a").Each(func(i int, selection *goquery.Selection) {
		if href, ok := selection.Attr("href"); ok {
			if cont, err := selection.Html(); err == nil {
				md := fmt.Sprintf(`[%v](%v)`, cont, href)
				selection.BeforeHtml(md)
				selection.Remove()
			}
		}
	})
	return doc
}

//[ok]handle tag ul、ol、li
func handleLi(doc *goquery.Document) *goquery.Document {
	var tags = []string{"ul", "ol"}
	for _, tag := range tags {
		doc.Find(tag).Each(func(i int, selection *goquery.Selection) {
			if !doesHasTagPreOrBlockQuote(selection) {
				if cont, err := selection.Html(); err == nil {
					cont = strings.Replace(cont, "\t", "", -1)
					cont = strings.Replace(cont, " ", "", -1)
					selection.BeforeHtml(cont)
					selection.Remove()
				}
			}
		})
	}
	doc.Find("li").Each(func(i int, selection *goquery.Selection) {
		if cont, err := selection.Html(); err == nil {
			if cont = strings.TrimSpace(cont); len(cont) > 0 {
				selection.BeforeHtml("- " + cont)
			}
			selection.Remove()
		}
	})
	return doc
}

//[ok]handle tag <hr/>
func handleHr(doc *goquery.Document) *goquery.Document {
	doc.Find("hr").Each(func(i int, selection *goquery.Selection) {
		selection.BeforeHtml("\n- - -\n")
		selection.Remove()
	})
	return doc
}

//[ok]handle tag <img/>
func handleImg(doc *goquery.Document) *goquery.Document {
	doc.Find("img").Each(func(i int, selection *goquery.Selection) {
		if src, ok := selection.Attr("src"); ok {
			alt := ""
			if val, ok := selection.Attr("alt"); ok {
				alt = val
			}
			md := fmt.Sprintf(`![%v](%v)`, alt, src)
			selection.BeforeHtml(md)
			selection.Remove()
		}
	})
	return doc
}

//[ok]handle tag h1~h6
func handleHead(doc *goquery.Document) *goquery.Document {
	heads := map[string]string{
		"title": "# ",
		"h1":    "# ",
		"h2":    "## ",
		"h3":    "### ",
		"h4":    "#### ",
		"h5":    "##### ",
		"h6":    "###### ",
	}
	for tag, replace := range heads {
		doc.Find(tag).Each(func(i int, selection *goquery.Selection) {
			text, _ := selection.Html()
			selection.BeforeHtml("\n" + replace + text + "\n")
			selection.Remove()
		})
	}
	return doc
}

//[ok]
func handleClosedTag(doc *goquery.Document) *goquery.Document {
	for tag, close := range closeTag {
		doc.Find(tag).Each(func(i int, selection *goquery.Selection) {
			text, _ := selection.Html()
			selection.BeforeHtml(close + text + close)
			selection.Remove()
		})
	}
	return doc
}

func handleNextLineTag(doc *goquery.Document) *goquery.Document {
	for _, tag := range nextLineTag {
		doc.Find(tag).Each(func(i int, selection *goquery.Selection) {
			selection.BeforeHtml("\n")
			selection.AfterHtml("\n")
		})
	}
	return doc
}

func handleTable(doc *goquery.Document) *goquery.Document {
	doc.Find("table").Each(func(i int, table *goquery.Selection) {
		rows := []string{}
		table.Find("tr").Each(func(i int, tr *goquery.Selection) {
			ths := []string{}
			tr.Find("th").Each(func(i int, trth *goquery.Selection) {
				ths = append(ths, getInnerHtml(trth))
			})
			if len(ths) > 0 {
				rows = append(rows, "|"+strings.Join(ths, "|")+"\n|-----\n")
			}
			tds := []string{}
			tr.Find("td").Each(func(i int, trtd *goquery.Selection) {
				tds = append(tds, getInnerHtml(trtd))
			})
			if len(tds) > 0 {
				rows = append(rows, "|"+strings.Join(tds, "|")+"\n")
			}
		})
		table.BeforeHtml(strings.Join(rows, ""))
		table.Remove()
	})
	return doc
}

func doesHasTagPreOrBlockQuote(selection *goquery.Selection) bool {
	sele := selection.Parent()
	if !sele.Is("body") { //非body标签
		if sele.Is("pre") || sele.Is("blockquote") {
			return true
		} else {
			return doesHasTagPreOrBlockQuote(sele)
		}
	}
	return false
}

func getInnerHtml(selection *goquery.Selection) (html string) {
	html, _ = selection.Html()
	return
}
