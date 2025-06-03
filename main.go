package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
	"math/rand"
	"time"
	"github.com/PuerkitoBio/goquery"
)

type SeoData struct {
	URL             string
	Title           string
	H1              string
	MetaDescription string
	StatusCode      int
}

type Parser interface {

}

type DefaultParser struct {

}

var userAgents = []string {
	"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/137.0.0.0 Safari/537.36",
	"Mozilla/5.0 (X11; Linux x86_64; rv:10.0) Gecko/20100101 Firefox/10.0",
}
// add 6-7 user agents

func randomUserAgent() string{
	rand.Seed(time.Now().Unix())
	randNum := rand.Int() % len(userAgents)
	return userAganets[randNum]
}

func isSitemap(url []string) ([]string, []string){
	extractSitemapFiles := []string{}
	pages := []string{}
	for _, page := range urls{
		foundSitemap == true{
			fmt.Println("Found Sitemap", page)
			extractSitemapFiles = append(extractSitemapFiles, page)

		}else{
			pages = append(pages, page)
		}
	}
	return sitemapFiles, pages
}

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

func makeRequest(url string)(*http.Response, error){
	client := http.Client{
		Timeout: 10*time.Second,
	}
	req, err := http.NewRequest("GET", url, nil)
	req.Header.Set("User-Agent", randomUserAgent())
	if err!= nil{
		return nil, err
	}
	
	res, err := client.Do(req)
	if err != nil{
		return nil, err
	} 
	return res, nil
}

func scrapeURLs(){

}

func scrapePage(url string, parser Parser)(SeoData, error){

	res,  err := crawlPage(url)
	if err != nil {
		return Seodata{}, err
	}
	data, err := parser.getSEOData(res)
	if err!= nil {
		return Seodata{}, err
	}
	return data, nil
}

func crawlPage() {

}

func getSEOData(){

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
