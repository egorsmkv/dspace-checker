# Dspace checker

## Overview

A simple Go app that collects scientific articles from Dspace repos using Geziyor

## Install

```bash
go build -o dspace-checker
```

## Usage

Add sources into the `sources.txt` file and then run the app:

```bash
./dspace-checker
```

It will generate the file `out.jsonl` with all results.

Program's output will be like the following:

```text
Scraping Started
URL already visited http://dspace.idgu.edu.ua/jspui/browse?type=dateissued&sort_by=2&order=DESC&rpp=20&etal=0&submit_browse=%D0%9E%D0%BD%D0%BE%D0%B2%D0%B8%D1%82%D0%B8
Crawled: (200) <GET https://dspace.megu.edu.ua:8443/jspui/browse?type=dateissued&sort_by=2&order=DESC&rpp=20&etal=0&submit_browse=Update>
Crawled: (200) <GET http://lib.pnu.edu.ua:8080/browse?type=dateissued&sort_by=2&order=DESC&rpp=20&etal=0&submit_browse=Update>
Crawled: (200) <GET http://dspace.idgu.edu.ua/jspui/browse?type=dateissued&sort_by=2&order=DESC&rpp=20&etal=0&submit_browse=%D0%9E%D0%BD%D0%BE%D0%B2%D0%B8%D1%82%D0%B8>
Crawled: (200) <GET http://dspace.zsmu.edu.ua/browse?type=dateissued&sort_by=2&order=DESC&rpp=20&etal=0>
Crawled: (200) <GET http://dspace.oneu.edu.ua/jspui/browse?type=dateissued&sort_by=2&order=DESC&rpp=20&etal=0>
Crawled: (200) <GET http://eir.zntu.edu.ua/browse?type=dateissued&sort_by=2&order=DESC&rpp=20&etal=0&submit_browse=Update>
Scraping Finished
```
