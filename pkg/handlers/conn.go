package handlers

import (
	"sync/atomic"

	"github.com/ksysoev/wasabi"
)

const (
	// initID is the initial ID for the connection.
	// it uses a large number to avoid conflicts with the IDs generated by the frontend.
	// not solving the problem fully, but it's a good start.
	initID = 1_000_000
)

type ConnState struct {
	Conn   wasabi.Connection
	currID int64
}

func NewConnState(conn wasabi.Connection) *ConnState {

	return &ConnState{
		Conn:   conn,
		currID: initID,
	}
}

func (c *ConnState) NextID() int64 {
	return atomic.AddInt64(&c.currID, 1)
}
