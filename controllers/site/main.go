package site

func (this *MainController) Index() {

import (
	"github.com/zaddone/dimon/models"
)

type MainController struct {
	baseController
}

func (this *MainController) Index() {
	var list []*models.Pager
	query := new(models.Pager).Query()
	count, _ := query.Count()
	if count > 0 {
		query.OrderBy("-Id").Limit(this.pagesize, (this.page-1)*this.pagesize).All(&list)
	}
	this.Data["list"] = list
	this.TplNames = "index.tpl"
}
