package controllers

import (
	"shorturl/models"

	"github.com/astaxie/beego"
	_ "github.com/astaxie/beego/cache/redis"
	"github.com/garyburd/redigo/redis"
)

var Conn redis.Conn

func Prepare() {

}

type ShortResult struct {
	UrlShort string
	UrlLong  string
}

type UrlController struct {
	beego.Controller
}

func (this *UrlController) Shorturl() {
	Conn, err := redis.Dial("tcp", "localhost:6379")
	if err != nil {
		beego.Error(err)
	}
	defer Conn.Close()

	var result ShortResult

	longurl := this.Input().Get("longurl")
	beego.Info(longurl)
	result.UrlLong = longurl
	urlmd5 := models.GetMD5(longurl)
	s, err := redis.String(Conn.Do("GET", urlmd5))
	if err != nil {
		beego.Error(err)
	}
	if s != "" {
		beego.Info(s)
		result.UrlShort = s
	} else {
		str := ""
		for i := 1; i <= 6; i++ {
			rand := models.GetRand()
			str += models.GetRandStr(rand)
		}
		result.UrlShort = str
		_, err := Conn.Do("SET", urlmd5, str)
		if err != nil {
			beego.Error(err)
		}
		_, err = Conn.Do("SET", result.UrlShort, longurl)
		if err != nil {
			beego.Error(err)
		}
	}
	this.Data["json"] = result
	this.ServeJSON()
}

func (this *UrlController) Longurl() {
	Conn, err := redis.Dial("tcp", "localhost:6379")
	if err != nil {
		beego.Error(err)
	}
	defer Conn.Close()

	var result ShortResult
	shorturl := this.Input().Get("shorturl")
	result.UrlShort = shorturl

	s, err := redis.String(Conn.Do("GET", shorturl))
	if err != nil {
		beego.Error(err)
	}

	if s != "" {
		result.UrlLong = s
	} else {
		result.UrlLong = ""
	}

	this.Data["json"] = result
	this.ServeJSON()
}
