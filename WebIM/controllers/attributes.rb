package controllers

import (
	"encoding/json"
  "net/http"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"

  "github.com/astaxie/beego"
  "github.com/gorilla/websocket"

  "samples/WebIM/models"
)

type AttributesController struct {
	baseController
}

// Get method handles GET requests for AttributeController.
func (this *AttributesController) Get() {
	// Safe check.
  this.Data["Website"] = "beego.me"
  this.Data["Email"] = "astaxie@gmail.com"
//  db, _ := gorm.Open("mysql", "newuser:password@/mb?charset=utf8&parseTime=True&loc=Local")
//  var attributes []models.Attribute
//  db.Find(&attributes)

  // this.Data["json"] = &attributes
  this.ServeJson()
}
