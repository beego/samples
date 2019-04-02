package controllers

import (
	"samples/shorturl/models"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/cache"
	"github.com/astaxie/beego/logs"
)

var (
	urlcache cache.Cache
)

func init() {
	urlcache, _ = cache.NewCache("memory", `{"interval":0}`)
}

type ShortResult struct {
	UrlShort string
	UrlLong  string
}

type ShortController struct {
	beego.Controller
}

// Use Get rather than Post so that we can simulate easier in the browser
func (this *ShortController) Get() {
	var result ShortResult
	longurl := this.Input().Get("longurl")
	logs.Info(longurl)
	result.UrlLong = longurl
	urlmd5 := models.GetMD5(longurl)
	logs.Info(urlmd5)
	if urlcache.IsExist(urlmd5) {
		result.UrlShort = urlcache.Get(urlmd5).(string)
	} else {
		result.UrlShort = models.Generate()
		err := urlcache.Put(urlmd5, result.UrlShort, 0)
		if err != nil {
			logs.Info(err)
		}
		err = urlcache.Put(result.UrlShort, longurl, 0)
		if err != nil {
			logs.Info(err)
		}
	}
	this.Data["json"] = result
	this.ServeJSON()
}
