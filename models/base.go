package models



import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
        _ "github.com/mattn/go-sqlite3"
	"os"
)

func Getenv(envName ,defaultVal string )(string){
	if str:=os.Getenv(envName);len(str)>0{
		return str
	}else{
		return defaultVal
	}
}

func getpqDaoServer(){

	orm.RegisterDataBase(
		"default",
		"postgres",
		fmt.Sprintf(
			"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
			Getenv("POSTGRESQL_PORT_5432_TCP_ADDR",beego.AppConfig.String("psql.addr")),
			Getenv("POSTGRESQL_PORT_5432_TCP_PORT",beego.AppConfig.String("psql.port")),
			Getenv("POSTGRESQL_USER",beego.AppConfig.String("psql.user")),
			Getenv("POSTGRESQL_PASSWORD",beego.AppConfig.String("psql.pwd")),
			Getenv("POSTGRESQL_INSTANCE_NAME",beego.AppConfig.String("psql.dbname")),
		),
	)
}
func getpq(){

	orm.RegisterDataBase(
		"default",
		"postgres",
		fmt.Sprintf(
			"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
			Getenv("POSTGRESQL_PORT_5432_TCP_ADDR",beego.AppConfig.String("psql.addr")),
			Getenv("POSTGRESQL_PORT_5432_TCP_PORT",beego.AppConfig.String("psql.port")),
			Getenv("POSTGRESQL_ENV_POSTGRES_USER",beego.AppConfig.String("psql.user")),
			Getenv("POSTGRESQL_ENV_POSTGRES_PASSWORD",beego.AppConfig.String("psql.pwd")),
			Getenv("POSTGRESQL_ENV_POSTGRES_NAME",beego.AppConfig.String("psql.dbname")),
		),
	)
}
func getmysql(){

	orm.RegisterDataBase("default", "mysql",
		fmt.Sprintf(
			"%s:%s@tcp(%s:%s)/%s?charset=utf8",
			Getenv("MYSQL_USER",beego.AppConfig.String("db.user")),
			Getenv("MYSQL_ENV_MYSQL_ROOT_PASSWORD",beego.AppConfig.String("db.pwd")),
			Getenv("MYSQL_PORT_3306_TCP_ADDR",beego.AppConfig.String("db.ip")),
			Getenv("MYSQL_PORT_3306_TCP_PORT",beego.AppConfig.String("db.port")),
			Getenv("MYSQL_NAME",beego.AppConfig.String("db.dbname")),
		),
	)
}
func getsqlite(){

	orm.RegisterDataBase("default", "sqlite3","/data_sqlite/dimon.db")
}
func init() {
	orm.Debug = true
	switch beego.AppConfig.String("db"){
	case "mysql":
		getmysql()
	case "postgresql":
		getpq()
	case "sqlite":
		getsqlite()
	}
//	user := beego.AppConfig.String("db.user")
//	pwd := beego.AppConfig.String("db.pwd")
//	ip:= beego.AppConfig.String("db.ip")
//	port,_ := beego.AppConfig.Int("db.port")
//	dbname := beego.AppConfig.String("db.dbname")
	orm.Debug = true
//	orm.RegisterModel(new(Type),new(Article))
//	orm.RegisterDriver("mysql", orm.DR_MySQL)
//	orm.RegisterDataBase("default", "mysql", fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8", user, pwd,ip, port, dbname))
}
func CreateTable() {
	name := "default"                          //数据库别名
	force := false                             //不强制建数据库
	verbose := true                            //打印建表过程
	err := orm.RunSyncdb(name, force, verbose) //建表
	if err != nil {
		beego.Error(err)
	}
}
func TableName(str string) string {
	return fmt.Sprintf("%s%s", beego.AppConfig.String("db.prefix"), str)
}
