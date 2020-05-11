package controllers

import (
	"fmt"
	"encoding/json"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/JiekeW/first_go_web/models"
)

type BaseController struct {
	beego.Controller
}

func (this *BaseController) response_data(data interface{}) {
	res := map[string]interface{}{
		"code": 0,
		"data": data,
	}
	this.Data["json"] = &res
	this.ServeJSON()
}

func (this *BaseController) response_err(err error) {
	res := map[string]interface{}{
		"code": 40040,
		"errmsg": fmt.Sprint(err),
	}
	this.Data["json"] = &res
	this.ServeJSON()
}

func (this *BaseController) check_err(err error, callback func ()) {
	if err != nil {
		this.response_err(err)
	} else {
		callback()
	}
}

type AccessTokenController struct {
	BaseController
}

func (this *AccessTokenController) Get() {
	id, _ := this.GetInt(":id")
	o := orm.NewOrm()
	if id > 0 {
		access := models.MoomaAccessToken{Id: id}
		err := o.Read(&access)
		this.check_err(err, func () {
			this.response_data(access)
		})
		
	} else {
		access := []*models.MoomaAccessToken{}
		_, err := o.QueryTable("mooma_access_token").All(&access)
		this.check_err(err, func () {
			this.response_data(access)
		})
	}
}

func (this *AccessTokenController) Post() {
	var access models.MoomaAccessToken
	data := this.Ctx.Input.RequestBody
	jsonErr := json.Unmarshal(data, &access)
	this.check_err(jsonErr, func () {
		o := orm.NewOrm()
		_, insertErr := o.Insert(&access)
		this.check_err(insertErr, func () {
			this.response_data(access)
		})
	})
}