package models

import(
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego"
	"time"
)

type Type struct{
	Id int
	CreateTime time.Time `orm:"auto_now_add;type(datetime)"`
	ModifyTime time.Time  `orm:"auto_now;type(datetime)"`
	Name string
	Nums int
}

func (m *Type) TableName() string {
	return TableName("type")
}

func init(){
	orm.RegisterModel(new(Type))
//	orm.RegisterModelWithPrefix(beego.AppConfig.String("db.prefix"), new(Type))
}

func GetAllType() (t1 []Type) {
	var t []Type
	_,err := orm.NewOrm().QueryTable("goblog_type").OrderBy("-CreateTime").All(&t);
	if err != nil{
		beego.Error(err);
	}
	return t;
}
