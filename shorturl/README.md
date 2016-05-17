# Shorturl

[中文文档](./README_ZH.md)

This sample is a API application based on beego. It has two API func:

- /v1/shorten
- /v1/expand

## Installation

```
cd $GOPATH/src/samples/shorturl
bee run
```

## Usage:

```
# shortening url example
http://localhost:8080/v1/shorten/?longurl=http://google.com

{
  "UrlShort": "5laZG",
  "UrlLong": "http://google.com"
}

# expanding url example
http://localhost:8080/v1/expand/?shorturl=5laZI

{
  "UrlShort": "5laZG",
  "UrlLong": "http://google.com"
}
```
