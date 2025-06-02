package main

import "fmt"

type SeoData struct {
	URL             string
	Title           string
	H1              string
	MetaDescription string
	StatusCode      int
}

type parser interface {

}

type DefaultParser struct {

}

userAganets  

func randomUserAgent

func extractSitemapURLs(startURL string) []string {
	Worklist := make(chan []string)
	toCrawl := []string{}
	var n int
	n++

	go func() {Worklist <- []string{startURL}}()

	for ; n>0 ; n--{

	list := <- Worklist
	for _, link := range list{
		n++} // check this for loop
		go func(link string){
			response, err := makeRequest(link)
			if err != nil {
				log.Printf("Error retrieving URL: %s", link)
			}
			urls, _ := extractSitemapURLs(response)
			if err != nil {
				log.Printf("Error extracting document from response, URL:%s",link)
			}
			sitemapFiles, pages := isSitemap(urls)
			if sitemapFiles != nil {
				worklist <- sitemapFiles
			}
			for _, page := range pages{
				toCrawl = append(toCrawl, page)
			}
		}(link)
	}
	return toCrawl
}

func ScrapeSitemap(url string) []SeoData{
	results := extractSitemapURLs(url)
	res := ScrapeURLs(results)
	return res
}

func main() {

	p := DefaultParser{}
	results := ScrapeSitemap("")
	for _, res := range results {
		fmt.Println(res)
	}

}
