<h1 align="center">Carrier</h1>



**Carrier** 是一个轻量级，高性能的、去中心化P2P底层网络模块。**Carrier** 致力成为一个可以让任何底层基础设施能够容易集成并运用的网络框架。

**Carrier** 使用高性能的Go语言实现。提供各种网络协议支持。**Carrier** 是将各种特性实现成组件的形式，允许上层应用可插拔得选择需要的模块。

## 愿景

**Carrier** 致力成为一个核心的底层基础设施。P2P模块是一个完整的系统中必要的环节，在服务化，分布式系统中也不可或缺。**Carrier** 希望在万物互联、连接无限制的网络中表现出自己的优势。P2P通信的场景是特别复杂的，在不同的环境，不同的网络会有不一样的情况。**Carrier** 提供最简单的接口给应用层使用，而在内部实现非常复杂的逻辑处理，但是还保持者高效的性能。

## 特性

- 可扩展的轻量级P2P网络模块
- 安全、加密、隐私保护的数据通信协议
- 多种协议支持，支持 tcp、udp、kcp、quic等
- 流式复用，高效，方便的连接管理
- 使用Peer-ID标识节点信息，支持连接动态IP
- 更完备、可靠的数据通信过程
- 支持各种组件的装载，支持KeepAlive、Backoff、AckReply
- 使用高性能的Protobuf 序列化消息内容，传输高效
- 支持压缩消息内容，包括zip, gz等
- 支持proxy、nat等内网通信
- 链路状态测量、网络质量统计
- 支持网络ID，通过网络ID区分不同的P2P网络

## 依赖

* [Golang](https://golang.org/doc/install) 1.12及以后的版本
* 使用`go mod` 进行包管理

## 安装

**Carrier** 不会编译成可执行程序，任何P2P项目可以直接`import` 使用。也可以通过`go mod` 形式安装

```shell
$ go get -v github.com/IPFS-eX/Carrier
```



## 例子

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



## 性能测试



双工通信模式，`tcp` 消息响应速度

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





## 许可证

**Carrier** 所有源码遵循[LGPL-3.0](LICENSE) 协议。