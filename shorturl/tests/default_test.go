package tests

import (
	"encoding/json"
	beetest "github.com/astaxie/beego/testing"
	"io/ioutil"
	"testing"
)

type ShortResult struct {
	UrlShort string
	UrlLong  string
}

func TestShort(t *testing.T) {
	request := beetest.Post("/v1/shorten")
	request.Param("longurl", "http://www.beego.me/")
	response, _ := request.Response()
	defer response.Body.Close()
	contents, _ := ioutil.ReadAll(response.Body)
	var s ShortResult
	json.Unmarshal(contents, &s)
	if s.UrlShort == "" {
		t.Fatal("shorturl is empty")
	}
}

func TestExpand(t *testing.T) {
	request := beetest.Get("/v1/expand")
	request.Param("shorturl", "5laZF")
	response, _ := request.Response()
	defer response.Body.Close()
	contents, _ := ioutil.ReadAll(response.Body)
	var s ShortResult
	json.Unmarshal(contents, &s)
	if s.UrlLong == "" {
		t.Fatal("urllong is empty")
	}
}
