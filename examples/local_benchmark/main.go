/**
 * Description:
 * Author: Yihen.Liu
 * Create: 2019-06-14
 */
package main

import (
	"flag"

	"github.com/IPFS-eX/carrier/common/log"
	"github.com/IPFS-eX/carrier/examples/local_benchmark/receiver"
	"github.com/IPFS-eX/carrier/examples/local_benchmark/sender"
)

func main() {
	log.InitLog(1)
	protocol := flag.String("protocol", "kcp", "protocol to use (kcp/tcp/udp)")
	flag.Parse()
	go receiver.Run(*protocol)
	go sender.Run(*protocol)
	select {}
}
