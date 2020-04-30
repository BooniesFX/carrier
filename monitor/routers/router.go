package routers

import (
	"github.com/astaxie/beego"
	"github.com/IPFS-eX/carrier/monitor/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{})
}
