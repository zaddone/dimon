package cache
import (
	"fmt"
	"github.com/astaxie/beego/cache"
	_"github.com/astaxie/beego/cache/redis"
	"github.com/astaxie/beego"
)
var Redis cache.Cache
func init(){

	var err error;
//	pwd := beego.AppConfig.String("redis.pwd")
//	ip:= beego.AppConfig.String("redis.ip")
//	port,_ := beego.AppConfig.Int("redis.port")
//	key:= beego.AppConfig.String("redis.key")
//	vedis, err = cache.NewCache("redis", fmt.Sprintf(`{"conn":"%s:%d"}`, ip, port))
	Redis, err = cache.NewCache("memory", {"interval":60}))
	if err != nil {
		beego.Error(err)
	}

}

