/**
 * Description:
 * Author: Yihen.Liu
 * Create: 2019-07-31
 */
package proxy

import (
	"github.com/IPFS-eX/carrier/network"
	"github.com/IPFS-eX/carrier/common/log"
)

func ProxyComponentRestart(protocol string, n *network.Network) {
	switch protocol {
	case "tcp":
		TcpComponentRestartUp(n)
	case "udp":
		UDPComponentRestartUp(n)
	case "kcp":
		KCPComponentRestartUp(n)
	case "quic":
		QuicComponentRestartUp(n)
	default:
		log.Error("Proxy component restart err, protocol:", protocol)
	}
}