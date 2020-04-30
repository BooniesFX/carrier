package sender

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net/http"
	_ "net/http/pprof"
	"os"
	"os/signal"
	"runtime"
	"runtime/pprof"
	"strconv"
	"time"

	"github.com/IPFS-eX/carrier/crypto/ed25519"
	"github.com/IPFS-eX/carrier/examples/local_benchmark/messages"
	"github.com/IPFS-eX/carrier/network"
	"github.com/IPFS-eX/carrier/network/components/keepalive"
	"github.com/IPFS-eX/carrier/types/opcode"
)

var profile = flag.String("profile", "", "write cpu profile to file")
var port = flag.Uint("port", 3002, "port to listen on")
var receiver = map[string]string{
	"udp":  "udp://127.0.0.1:3001",
	"tcp":  "tcp://127.0.0.1:3001",
	"kcp":  "kcp://127.0.0.1:3001",
	"quic": "quic://127.0.0.1:3001",
}

func Run(protocol string) {
	flag.Set("logtostderr", "true")

	go func() {
		log.Println(http.ListenAndServe("127.0.0.1:7070", nil))
	}()

	runtime.GOMAXPROCS(runtime.NumCPU())
	opcode.RegisterMessageType(opcode.Opcode(1000), &messages.BasicMessage{})

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	go func() {
		<-c
		pprof.StopCPUProfile()
		os.Exit(0)
	}()

	if *profile != "" {
		f, err := os.Create(*profile)
		if err != nil {
			log.Fatal(err)
		}
		pprof.StartCPUProfile(f)
	}

	builderOpt := []network.BuilderOption{
		network.WriteFlushLatency(1 * time.Millisecond),
		network.WriteTimeout(int(time.Duration(30))),
	}
	builder := network.NewBuilderWithOptions(builderOpt...)
	builder.SetAddress(protocol + "://127.0.0.1:" + strconv.Itoa(int(*port)))
	builder.SetListenAddr(protocol + "://127.0.0.1:" + strconv.Itoa(int(*port)))
	builder.SetKeys(ed25519.RandomKeyPair())
	options := []keepalive.ComponentOption{
		keepalive.WithKeepaliveInterval(keepalive.DefaultKeepaliveInterval),
		keepalive.WithKeepaliveTimeout(keepalive.DefaultKeepaliveTimeout),
	}

	builder.AddComponent(keepalive.New(options...))
	net, err := builder.Build()
	if err != nil {
		panic(err)
	}
	net.SetNetworkID(1)
	net.SetBootstrapWaitSecond(time.Duration(15) * time.Second)

	go net.Listen()
	net.BlockUntilListening()
	peerIds := net.Bootstrap([]string{receiver[protocol]})

	time.Sleep(500 * time.Millisecond)

	fmt.Printf("Spamming messages to %s...\n", receiver[protocol])

	client := net.GetPeerClient(peerIds[0])
	if client == nil {
		panic("client is nil")
	}
	net.Components.Delete(keepalive.ComponentID)
	ctx := network.WithSignMessage(context.Background(), false)
	count := 0
	go func() {
		for range time.Tick(1 * time.Second) {
			fmt.Printf("TX %d messages/s \n", count)

			count = 0
		}

	}()
	for {
		err = client.Tell(ctx, &messages.BasicMessage{})
		count += 1
		if err != nil {
			panic(err)
		}
	}
}
