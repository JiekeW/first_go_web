package models

import (
	"github.com/astaxie/beego/orm"
	"time"
)

type MoomaAccessToken struct {
	Id                    int `json:"id"`
	Providing_platform    string `json:"providing_platform"`
	Access_token          string `json:"access_token"`
	Expiry_second         string `json:"expiry_second"`
	Sync_time             time.Time `orm:"auto_now_add;type(datetime)" json:"sync_time"`
}

func init() {
	// 需要在 init 中注册定义的 model
	orm.RegisterModel(new(MoomaAccessToken))
}