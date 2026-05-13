package main

import (
	"encoding/json"
	"fmt"
	"gen/lib/imports"
	"net/http"
	"os"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/marcbran/jpoet/pkg/jpoet"
)

func main() {
	err := run(os.Args[1])
	if err != nil {
		panic(err)
	}
}

func run(outDir string) error {
	elements, err := pullElements()
	if err != nil {
		return err
	}
	err = generate(elements, outDir)
	if err != nil {
		return err
	}
	return nil
}

func pullElements() ([]string, error) {
	res, err := http.Get("https://developer.mozilla.org/en-US/docs/Web/HTML/Element")
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		return nil, fmt.Errorf("status code error: %d %s", res.StatusCode, res.Status)
	}

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		return nil, err
	}

	var elem []string
	doc.Find("table tr").Each(func(i int, s *goquery.Selection) {
		elemText := s.Find("td").First().Text()
		elems := strings.Split(elemText, ",")
		for _, e := range elems {
			e = strings.TrimSpace(e)
			if !strings.HasPrefix(e, "<") {
				continue
			}
			if !strings.HasSuffix(e, ">") {
				continue
			}
			elem = append(elem, e[1:len(e)-1])
		}
	})
	return elem, nil
}

func generate(elements []string, outDir string) error {
	elementsJSON, err := json.Marshal(elements)
	if err != nil {
		return err
	}
	err = os.MkdirAll(outDir, 0755)
	if err != nil {
		return err
	}
	err = jpoet.Eval(
		jpoet.FileImport([]string{}),
		jpoet.FSImport(lib),
		jpoet.FSImport(imports.Fs),
		jpoet.TLACode("elements", string(elementsJSON)),
		jpoet.FileInput("./lib/main.libsonnet"),
		jpoet.Serialize(false),
		jpoet.DirectoryOutput(outDir),
	)
	if err != nil {
		return err
	}
	return nil
}
