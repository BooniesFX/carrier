<h1 align="center">Carrier</h1>

English | [中文](README_CN.md)



**Carrier** is an lightweight, high-performance, decentralized P2P network infrastructure. carrier is ease-to-use for every underlay infrastructure and enhance its network ability.

**Carrier** is writen in Go program language. It is the best network framework for all network protocol. Carrier compose all features by components. Components are pluggable, and easy to be embedded for applications.



## Vision

**Carrier** is a core module for underlay infrastructure. P2P is a necessary module for severless or distributed system. Carrier wants to provide a huge advantage for interneting everytion or connection unlimited world. P2P scenario are complicated and different for network transport. **Carrier** provides a simplest interface or protocol for applications to use. Complicated logic is handling inside, but come out with a high performance interface for outside.

## Features

- Scalable lightweight P2P network modules
- Secure, encrypted, privacy protected transport protocol
- Supports multiple protocol, like tcp,udp,kcp,quic
- Stream multiplexing, efficient and easy to control connection
- Use peer-id instead of IP for recognizing peer information
- More complete, reliable communication process
- Plugins supported, like KeepAlive、Backoff、AckReply
- Use high-performance protocol buffer for message read/write
- Support compress message by gzip or zip. Reducing overhead
- Proxy or nat is supported, easy to use for LAN
- Record route QoS metrics, and network responsibility
- Use network id to split different network.

## Dependencies

* [Golang](https://golang.org/doc/install) version 1.12 or later
* Use `go mod`  for module management

## Install

**Carrier** is not a excutable binary. Any P2P project can import it easily, or use it by `go mod` 

```shell
$ go get -v github.com/IPFS-eX/carrier
```



## Example

```go
builderOpt := []network.BuilderOption{
	network.WriteFlushLatency(1 * time.Millisecond),
	network.WriteTimeout(int(time.Duration(30))),
}
builder := network.NewBuilderWithOptions(builderOpt...)
builder.SetKeys(ed25519.RandomKeyPair())
builder.SetListenAddr(fmt.Sprintf("%s://%s:%s", protocol, intranetIP, port))
builder.SetAddress(fmt.Sprintf("%s://%s:%s", protocol, addr, port))
// add msg component
component := new(NetComponent)
component.Net = this
builder.AddComponent(component)
var err error
p2p, err = builder.Build()
if err != nil {
	return err
}
var once sync.Once
once.Do(func() {
	for k, v := range.opCodes {
		if err := opcode.RegisterMessageType(k, v); err != nil {
		}
	}
})
p2p.SetNetworkID(networkId)
p2p.SetBootstrapWaitSecond(time.Duration(15) * time.Second)
go p2p.Listen()
p2p.BlockUntilListening()
```



## Benchmarks



TCP message transport speed in duplex mode.

```
MacBook Pro (13-inch, 2017, Four Thunderbolt 3 Ports)
3.5 GHz Intel Core i7
16 GB 2133 MHz LPDDR3

// Output:
TX 15931 messages/s 
RX 10845 messages/s 
TX 11437 messages/s 
RX 10577 messages/s 
TX 10813 messages/s 
RX 10416 messages/s 
TX 10812 messages/s 
RX 10503 messages/s 
TX 10199 messages/s 
RX 10342 messages/s 
```




## License

The **Carrier** source code is available under the [LGPL-3.0](LICENSE) license.