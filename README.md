<!DOCTYPE html>
<html lang="en">
<head>
  <meta charset="UTF-8">
  <title>SitemapCrawler</title>
  <style>
    body { font-family: Arial, sans-serif; margin: 2em; }
    code, pre { background: #f4f4f4; padding: 2px 4px; border-radius: 3px; }
    pre { padding: 1em; }
    h1, h2, h3 { color: #2c3e50; }
    ul { margin-bottom: 1em; }
  </style>
</head>
<body>
  <h1>SitemapCrawler</h1>
  <p>
    <strong>SitemapCrawler</strong> is a simple Go program to crawl a website’s sitemap, extract all URLs, and collect basic SEO data (title, H1, meta description, status code) for each page.
  </p>

  <h2>Features</h2>
  <ul>
    <li>Fetches and parses XML sitemaps</li>
    <li>Recursively follows sitemap index files</li>
    <li>Scrapes SEO data from each page</li>
    <li>Supports concurrent crawling with configurable concurrency</li>
    <li>Randomizes User-Agent for each request</li>
  </ul>

  <h2>Requirements</h2>
  <ul>
    <li>Go 1.18 or newer</li>
    <li><a href="https://github.com/PuerkitoBio/goquery">github.com/PuerkitoBio/goquery</a></li>
  </ul>
  <pre>
go get github.com/PuerkitoBio/goquery
  </pre>

  <h2>Usage</h2>
  <ol>
    <li>Clone this repository.</li>
    <li>Build and run:</li>
  </ol>
  <pre>
go run main.go
  </pre>
  <p>
    By default, it scrapes the sitemap at <code>https://www.quicksprout.com/sitemap.xml</code> with a concurrency of 10.
  </p>

  <h2>Example Output</h2>
  <pre>
{URL:https://www.quicksprout.com/ Title:... H1:... MetaDescription:... StatusCode:200}
...
  </pre>

  <h2>Customization</h2>
  <p>
    To crawl a different sitemap, edit the URL in the <code>main()</code> function:
  </p>
  <pre>
results := ScrapeSitemap("https://example.com/sitemap.xml", p, 10)
  </pre>

  <h2>License</h2>
  <p>MIT License</p>
</body>
</html>
