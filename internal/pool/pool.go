package pool

import (
	"sync"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type ConnPool struct {
	conns []*grpc.ClientConn
	mu    sync.Mutex
}

var pool *ConnPool

// NewConnPool creates a new connection pool with the giver size.
func NewConnPool(size int, target string) (*ConnPool, error) {
	p := &ConnPool{
		conns: make([]*grpc.ClientConn, size),
	}

	var err error

	creds := insecure.NewCredentials()

	for i := 0; i < size; i++ {
		p.conns[i], err = grpc.Dial(target, grpc.WithTransportCredentials(creds), grpc.WithBlock())
		if err != nil {
			return nil, err
		}
	}
	return p, nil
}

// Get gets a connection from the pool
func (p *ConnPool) Get() *grpc.ClientConn {
	p.mu.Lock()
	defer p.mu.Unlock()

	conn := p.conns[0]

	p.conns = append(p.conns[1:], p.conns[0])

	return conn
}

// Close close all connections in the pool.
func (p *ConnPool) Close() {
	p.mu.Lock()
	defer p.mu.Unlock()

	for _, conn := range p.conns {
		conn.Close()
	}
}

// GetDefaultConnPool returns the connection pool.
func GetDefaultConnPool() *ConnPool {
	return pool
}

func init() {
	// create a connection pool with 10 connections.
	var err error
	pool, err = NewConnPool(10, "localhost:50051")
	if err != nil {
		panic(err)
	}
}
