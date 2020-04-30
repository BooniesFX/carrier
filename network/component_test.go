package network

import (
	"fmt"
	"sync/atomic"
	"testing"

	"github.com/IPFS-eX/carrier/crypto/ed25519"
	"github.com/stretchr/testify/assert"
)

type MockComponent struct {
	*Component

	startup        int32
	receive        int32
	cleanup        int32
	peerConnect    int32
	peerDisconnect int32
}

func (p *MockComponent) Startup(net *Network) {
	atomic.AddInt32(&p.startup, 1)
}

func (p *MockComponent) Receive(ctx *ComponentContext) error {
	atomic.AddInt32(&p.receive, 1)
	return nil
}

func (p *MockComponent) Cleanup(net *Network) {
	atomic.AddInt32(&p.cleanup, 1)
}

func (p *MockComponent) PeerConnect(client *PeerClient) {
	atomic.AddInt32(&p.peerConnect, 1)
}

func (p *MockComponent) PeerDisconnect(client *PeerClient) {
	atomic.AddInt32(&p.peerDisconnect, 1)
}

func TestComponentHooks(t *testing.T) {
	host := "localhost"
	var nodes []*Network
	nodeCount := 4

	for i := 0; i < nodeCount; i++ {
		builder := NewBuilder()
		builder.SetKeys(ed25519.RandomKeyPair())
		builder.SetAddress(FormatAddress("tcp", host, uint16(GetRandomUnusedPort())))
		builder.AddComponent(new(MockComponent))

		node, err := builder.Build()
		if err != nil {
			fmt.Println(err)
		}

		go node.Listen()

		nodes = append(nodes, node)
	}

	for _, node := range nodes {
		node.BlockUntilListening()
	}

	//for i, node := range nodes {
	//	if i != 0 {
	//		node.Bootstrap(nodes[0].Address)
	//	}
	//}
	//
	//time.Sleep(500 * time.Millisecond)
	//
	//for _, node := range nodes {
	//	node.Close()
	//}
	//
	//time.Sleep(500 * time.Millisecond)
	//
	//if startup != nodeCount {
	//	t.Fatalf("startup hooks error, got: %d, expected: %d", startup, nodeCount)
	//}
	//
	//if receive < nodeCount*2 { // Cannot in specific time
	//	t.Fatalf("receive hooks error, got: %d, expected at least: %d", receive, nodeCount*2)
	//}
	//
	//if cleanup != nodeCount {
	//	t.Fatalf("cleanup hooks error, got: %d, expected: %d", cleanup, nodeCount)
	//}
	//
	//if peerConnect < nodeCount*2 {
	//	t.Fatalf("connect hooks error, got: %d, expected at least: %d", peerConnect, nodeCount*2)
	//}
	//
	//if peerDisconnect < nodeCount*2 {
	//	t.Fatalf("disconnect hooks error, got: %d, expected at least: %d", peerDisconnect, nodeCount*2)
	//}
}

func TestRegisterComponent(t *testing.T) {
	t.Parallel()

	ComponentID := (*Component)(nil)

	b := NewBuilder()
	b.AddComponentWithPriority(-99999, new(Component))

	n, err := b.Build()
	assert.Equal(t, nil, err)

	p, ok := n.Components.Get(ComponentID)
	assert.Equal(t, true, ok)

	Component := p.(*Component)
	assert.NotEqual(t, nil, Component)
}
