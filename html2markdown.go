//Author:TruthHun
//Email: TruthHun@QQ.COM
//Date:  2018-02-03
package html2md

import (
	"strings"

	"fmt"

	"github.com/PuerkitoBio/goquery"
	"github.com/astaxie/beego/logs"
)

var closeTag = map[string]string{
	"del":    "~~",
	"b":      "**",
	"strong": "**",
	"i":      "_",
	"em":     "_",
	"dfn":    "_",
	"var":    "_",
	"cite":   "_",
	"br":     "\n",
	"span":   "",
}

var blockTag = []string{
	"div",
	"figure",
	"p",
	"article",
	"nav",
	"footer",
	"header",
	"section",
}
var nextlineTag = []string{
	"pre", "blockquote", "table",
}

//convert html to markdown
//将html转成markdown
func Convert(htmlstr string) (md string) {
	doc, _ := goquery.NewDocumentFromReader(strings.NewReader(htmlstr))
	handleNextLine(doc)  //<div>...
	handleBlockTag(doc)  //<div>...
	handleA(doc)         //<a>
	handleImg(doc)       //<img>
	handleHead(doc)      //h1~h6
	handleClosedTag(doc) //<strong>、<i>、eg..
	handleHr(doc)        //<hr>
	handleLi(doc)        //<ul>、<li>
	handleTable(doc)     //<table>
	md, _ = doc.Find("body").Html()
	md = strings.Replace(md, "</blockquote>", "\n</blockquote>", -1)
	return
}

func handleBlockTag(doc *goquery.Document) *goquery.Document {
	for _, tag := range blockTag {
		hasTag := true
		for hasTag {
			if tagEle := doc.Find(tag); len(tagEle.Nodes) > 0 {
				tagEle.Each(func(i int, selection *goquery.Selection) {
					selection.BeforeHtml("\n" + getInnerHtml(selection) + "\n")
					selection.Remove()
				})
			} else {
				hasTag = false
			}
		}
	}
	return doc
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
			if cont, err := selection.Html(); err == nil {
				cont = strings.Replace(cont, "\t", "", -1)
				cont = strings.Replace(cont, " ", "", -1)
				selection.BeforeHtml(cont)
				selection.Remove()
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
			text := selection.Text()
			selection.BeforeHtml("\n" + replace + text + "\n")
			selection.Remove()
		})
	}
	return doc
}

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

func handleNextLine(doc *goquery.Document) *goquery.Document {
	for _, tag := range nextlineTag {
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

func getInnerHtml(selection *goquery.Selection) (html string) {
	var err error
	html, _ = selection.Html()
	if err != nil {
		logs.Error(err)
	}
	return
}
