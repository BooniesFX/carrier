package monitor

import (
	"github.com/astaxie/beego"
	"github.com/IPFS-eX/carrier/monitor/controllers"
	_ "github.com/IPFS-eX/carrier/monitor/routers"
	"github.com/IPFS-eX/carrier/network"
)

func Run(network *network.Network) {
	controllers.InitMonitor(network)
	beego.Run()
}
