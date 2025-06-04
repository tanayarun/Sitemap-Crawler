# 🕸️ SitemapCrawler

**SitemapCrawler** is a simple and concurrent Go-based web crawler that reads a sitemap (or sitemap index), extracts all URLs, and scrapes basic SEO information for each page.

It fetches:

* Page title
* First `<h1>` tag
* Meta description
* HTTP status code

---

##  Features

*  Parses XML sitemaps and sitemap indexes
*  Collects SEO metadata from each URL
*  Randomized User-Agent for each request (to avoid detection)
*  Concurrent crawling with configurable concurrency
*  Clean separation of logic using `Parser` interface

---

##  Tech Stack

* Language: **Go 1.18+**
* HTML parsing: [`goquery`](https://github.com/PuerkitoBio/goquery)

---

##  Installation

1. Install Go 1.18 or newer.
2. Install dependency:

```bash
go get github.com/PuerkitoBio/goquery
```

3. Clone and run:

```bash
git clone https://github.com/tanayarun/Sitemap-Crawler.git
cd SitemapCrawler
go run main.go
```

---

##  Usage

By default, it scrapes:

```
https://www.quicksprout.com/sitemap.xml
```

To change the sitemap URL, edit the following line in `main()`:

```go
results := ScrapeSitemap("https://example.com/sitemap.xml", p, 10)
```

Where:

* `p` is a `Parser` (e.g. `DefaultParser`)
* `10` is the concurrency level

---

##  Project Structure

| File              | Description                                  |
| ----------------- | -------------------------------------------- |
| `main.go`         | Entry point and overall control flow         |
| `ScrapeSitemap()` | Orchestrates sitemap extraction and scraping |
| `DefaultParser`   | Extracts SEO metadata from HTML pages        |
| `makeRequest()`   | Makes GET requests with random User-Agent    |
| `extractUrls()`   | Parses `<loc>` tags from XML sitemap         |
| `scrapeUrls()`    | Concurrently processes list of URLs          |

