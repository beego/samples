package controllers

import (
	"github.com/jinzhu/gorm"
	 _ "github.com/jinzhu/gorm/dialects/mysql"
	 "samples/WebIM/models"

)


// "github.com/astaxie/beego"
// "github.com/gorilla/websocket"
//

type AttributesController struct {
	baseController
}

// Get method handles GET requests for AttributeController.
func (this *AttributesController) Get() {
	// Safe check.
 	db, _ := gorm.Open("mysql", "newuser:password@/mb?charset=utf8&parseTime=True&loc=Local")
 	var attributes []models.Attribute
 	db.Find(&attributes)

  this.Data["json"] = &attributes
  this.ServeJSON()
}
