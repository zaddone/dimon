package models

import (
	"crypto/md5"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/lib/pq"
	"net/url"
	"strings"
        "os"
)

func Getenv(envName ,defaultVal string )(string){
    if str:=os.Getenv(envName);len(str)>0{
        return str
    }else{
        return defaultVal
    }
}
func getpq()(string){
    var dburl ="host=" + Getenv("POSTGRESQL_PORT_5432_TCP_ADDR","127.0.0.1")+
        " port=" + Getenv("POSTGRESQL_PORT_5432_TCP_PORT","5432")+
        " user=" + Getenv("POSTGRESQL_ENV_POSTGRES_USER","postgres")+
        " password=" + Getenv("POSTGRESQL_ENV_POSTGRES_PASSWORD","test2015")+
        " dbname=" + Getenv("POSTGRESQL_NAME","dimon")+
        " sslmode=disable"
    return dburl
}
func init() {
	dburl :=getpq()
	orm.RegisterDataBase("default", "postgres", dburl)
	orm.RegisterModel(new(User), new(Post), new(Tag), new(Option), new(TagPost), new(Mood), new(Photo), new(Album), new(Link))
	if beego.AppConfig.String("runmode") == "dev" {
		orm.Debug = true
	}
}

func Md5(buf []byte) string {
	hash := md5.New()
	hash.Write(buf)
	return fmt.Sprintf("%x", hash.Sum(nil))
}

func Rawurlencode(str string) string {
	return strings.Replace(url.QueryEscape(str), "+", "%20", -1)
}

func GetOptions() map[string]string {
	if !Cache.IsExist("options") {
		var result []*Option
		o := orm.NewOrm()
		o.QueryTable(&Option{}).All(&result)
		options := make(map[string]string)
		for _, v := range result {
			options[v.Name] = v.Value
		}
		Cache.Put("options", options)
	}
	v := Cache.Get("options")
	return v.(map[string]string)
}

func GetLatestBlog() []*Post {
	if !Cache.IsExist("latestblog") {
		var result []*Post
		query := new(Post).Query().Filter("status", 0).Filter("urltype", 0)
		count, _ := query.Count()
		if count > 0 {
			query.OrderBy("-posttime").Limit(8).All(&result)
		}
		Cache.Put("latestblog", result)
	}
	v := Cache.Get("latestblog")
	return v.([]*Post)
}

func GetHotBlog() []*Post {
	if !Cache.IsExist("hotblog") {
		var result []*Post
		new(Post).Query().Filter("status", 0).Filter("urltype", 0).OrderBy("-views").Limit(5).All(&result)
		Cache.Put("hotblog", result)
	}
	v := Cache.Get("hotblog")
	return v.([]*Post)
}

func GetLinks() []*Link {
	if !Cache.IsExist("links") {
		var result []*Link
		new(Link).Query().OrderBy("-rank").All(&result)
		Cache.Put("links", result)
		fmt.Println(result)
	}
	v := Cache.Get("links")
	return v.([]*Link)
}

//返回带前缀的表名
func TableName(str string) string {
	return fmt.Sprintf("%s%s", beego.AppConfig.String("dbprefix"), str)
}
