package main

import(
	_ "github.com/zaddone/dimon/routers"
	"github.com/zaddone/dimon/models"
	_ "github.com/zaddone/dimon/filter"
	_ "github.com/zaddone/dimon/cache"
	"github.com/astaxie/beego"
)

func main() {
	models.CreateTable()
	beego.Run()
}
