package main
import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"fmt"
	"samples/WebIM/models"
	"os"
)
func main() {
	db, err := gorm.Open("mysql", "newuser:password@/mb?charset=utf8&parseTime=True&loc=Local")
	fmt.Printf("%v\n", err)
	fmt.Printf("%v\n", db)
	fmt.Printf("%v\n", db.HasTable("events"))
	var event models.Event
	data := db.Find(&event)
	rows, _ := data.Rows()
	defer rows.Close()
	for rows.Next() {
		var content string
    var user string
		var Type int
		var timestamp int64
		fmt.Printf("%v\n", rows.Scan(&Type, &user, &timestamp, &content))
		fmt.Printf("%v\n", user)
		fmt.Printf("%v\n", Type)
		fmt.Printf("%v\n", content)
	}
}
