package site

import (
	"github.com/astaxie/beego"
)

type baseController struct {
	beego.Controller
	options  map[string]string
	right    string
	page     int
	pagesize int
}
