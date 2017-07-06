package controllers

import (
	"encoding/json"
	"strconv"

	"github.com/astaxie/beego"
	"github.com/beego/samples/todo/models"
)

type TaskController struct {
	beego.Controller
}

// Example:
//
//   req: GET /task/
//   res: 200 {"Tasks": [
//          {"ID": 1, "Title": "Learn Go", "Done": false},
//          {"ID": 2, "Title": "Buy bread", "Done": true}
//        ]}
func (this *TaskController) ListTasks() {
	res := struct{ Tasks []*models.Task }{models.DefaultTaskList.All()}
	this.Data["json"] = res
	this.ServeJSON()
}

// Examples:
//
//   req: POST /task/ {"Title": ""}
//   res: 400 empty title
//
//   req: POST /task/ {"Title": "Buy bread"}
//   res: 200
func (this *TaskController) NewTask() {
	req := struct{ Title string }{}
	if err := json.Unmarshal(this.Ctx.Input.RequestBody, &req); err != nil {
		this.Ctx.Output.SetStatus(400)
		this.Ctx.Output.Body([]byte("empty title"))
		return
	}
	t, err := models.NewTask(req.Title)
	if err != nil {
		this.Ctx.Output.SetStatus(400)
		this.Ctx.Output.Body([]byte(err.Error()))
		return
	}
	models.DefaultTaskList.Save(t)
}

// Examples:
//
//   req: GET /task/1
//   res: 200 {"ID": 1, "Title": "Buy bread", "Done": true}
//
//   req: GET /task/42
//   res: 404 task not found
func (this *TaskController) GetTask() {
	id := this.Ctx.Input.Param(":id")
	beego.Info("Task is ", id)
	intid, _ := strconv.ParseInt(id, 10, 64)
	t, ok := models.DefaultTaskList.Find(intid)
	beego.Info("Found", ok)
	if !ok {
		this.Ctx.Output.SetStatus(404)
		this.Ctx.Output.Body([]byte("task not found"))
		return
	}
	this.Data["json"] = t
	this.ServeJSON()
}

// Example:
//
//   req: PUT /task/1 {"ID": 1, "Title": "Learn Go", "Done": true}
//   res: 200
//
//   req: PUT /task/2 {"ID": 2, "Title": "Learn Go", "Done": true}
//   res: 400 inconsistent task IDs
func (this *TaskController) UpdateTask() {
	id := this.Ctx.Input.Param(":id")
	beego.Info("Task is ", id)
	intid, _ := strconv.ParseInt(id, 10, 64)
	var t models.Task
	if err := json.Unmarshal(this.Ctx.Input.RequestBody, &t); err != nil {
		this.Ctx.Output.SetStatus(400)
		this.Ctx.Output.Body([]byte(err.Error()))
		return
	}
	if t.ID != intid {
		this.Ctx.Output.SetStatus(400)
		this.Ctx.Output.Body([]byte("inconsistent task IDs"))
		return
	}
	if _, ok := models.DefaultTaskList.Find(intid); !ok {
		this.Ctx.Output.SetStatus(400)
		this.Ctx.Output.Body([]byte("task not found"))
		return
	}
	models.DefaultTaskList.Save(&t)
}
