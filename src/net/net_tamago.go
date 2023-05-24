// Copyright 2018 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Fake networking for js/wasm. It is intended to allow tests of other package to pass.

//go:build tamago

package net

import (
	"context"
	"errors"
)

type netFD struct {
}

func (sd *sysDialer) dialTCP(ctx context.Context, la, ra *TCPAddr) (Conn, error) {
	return nil, errors.New("not implemented")
}

func (sd *sysDialer) dialUDP(ctx context.Context, la, ra *UDPAddr) (Conn, error) {
	return nil, errors.New("not implemented")
}

func (sd *sysDialer) dialIP(ctx context.Context, la, ra *IPAddr) (Conn, error) {
	return nil, errors.New("not implemented")
}

func (sd *sysDialer) dialUnix(ctx context.Context, la, ra *UnixAddr) (Conn, error) {
	return nil, errors.New("not implemented")
}

func (sl *sysListener) listenTCP(ctx context.Context, la *TCPAddr) (Listener, error) {
	return nil, errors.New("not implemented")
}

func (sl *sysListener) listenUDP(ctx context.Context, la *UDPAddr) (*UDPConn, error) {
	return nil, errors.New("not implemented")
}

func (sl *sysListener) listenIP(ctx context.Context, la *IPAddr) (*IPConn, error) {
	return nil, errors.New("not implemented")
}

func (sl *sysListener) listenUnix(ctx context.Context, la *UnixAddr) (Listener, error) {
	return nil, errors.New("not implemented")
}

func (sl *sysListener) listenUnixgram(ctx context.Context, la *UnixAddr) (*UnixConn, error) {
	return nil, errors.New("not implemented")
}
