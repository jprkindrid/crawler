# Web Crawler

## Overview
This project is a simple web crawler written in Go, from the boot.dev guided course. It is designed to fetch and analyze web pages, extract URLs, normalize them, and generate a report of internal links found on a given website.

## Features
- Fetch HTML content from web pages.
- Extract URLs from the HTML content.
- Normalize URLs to ensure consistency.
- Generate a report summarizing the internal links found on the website.
- Sort the report by the number of internal links.

## How to Run
1. Clone the repository:
   ```bash
   git clone https://github.com/jprkindrid/crawler.git
   cd crawler
   ```
2. Build the project:
   ```bash
   go build -o crawler
   ```
3. Run the application:
   ```bash
   ./crawler <baseURL> <maxConcurrencyCount> <maxPagesToCrawl>
   ```
   Replace `<baseURL>` with the URL of the website you want to crawl.
   `<maxConcurrencyCount>` sets the amount of concurrency channels used to crawl the website. 3 By default
   `<maxPagesToCrawl>` sets a depth limit for the amount of pages you want to crawl on the website, to prevent infinite crawling. 10 by default.

## Example Output
```
=============================
REPORT for https://example.com
=============================
Found 10 internal links to https://example.com/page1
Found 5 internal links to https://example.com/page2
...
```

## Contributing
Feel free to fork the repository and submit pull requests for improvements or bug fixes.

## License
This project is licensed under the MIT License. See the LICENSE file for details.
