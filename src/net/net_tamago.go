// Copyright 2018 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Fake networking for js/wasm. It is intended to allow tests of other package to pass.

//go:build tamago

package net

import (
	"context"
	"errors"
	"io"
	"os"
	"syscall"
	"time"
)

type netFD struct {
	net   string
	laddr Addr
	raddr Addr
}

func (sd *sysDialer) dialTCP(ctx context.Context, la, ra *TCPAddr) (Conn, error) {
	return nil, errors.New("not implemented")
}

func (sd *sysDialer) dialUDP(ctx context.Context, la, ra *UDPAddr) (Conn, error) {
	return nil, errors.New("not implemented")
}

func (sd *sysDialer) dialIP(ctx context.Context, la, ra *IPAddr) (*IPConn, error) {
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

func (c *TCPConn) readFrom(r io.Reader) (int64, error) {
	return -1, errors.New("not implemented")
}

func (c *IPConn) readFrom(p []byte) (int, *IPAddr, error) {
	return -1, nil, errors.New("not implemented")
}

func (c *IPConn) writeTo(b []byte, addr *IPAddr) (int, error) {
	return -1, errors.New("not implemented")
}

func (c *IPConn) readMsg(b, oob []byte) (n, oobn, flags int, addr *IPAddr, err error) {
	return -1, 0, 0, nil, errors.New("not implemented")
}

func (c *IPConn) writeMsg(b, oob []byte, addr *IPAddr) (n, oobn int, err error) {
	return -1, 0, errors.New("not implemented")
}

func (fd *netFD) Read(p []byte) (n int, err error) {
	return -1, errors.New("not implemented")
}

func (fd *netFD) Write(p []byte) (n int, err error) {
	return -1, errors.New("not implemented")
}

func (fd *netFD) Close() error {
	return errors.New("not implemented")
}

func (fd *netFD) SetDeadline(t time.Time) error {
	return errors.New("not implemented")
}
func (fd *netFD) SetReadDeadline(t time.Time) error {
	return errors.New("not implemented")
}
func (fd *netFD) SetWriteDeadline(t time.Time) error {
	return errors.New("not implemented")
}

func (fd *netFD) dup() (*os.File, error) {
	return nil, errors.New("not implemented")
}

func (fd *netFD) closeRead() error {
	return errors.New("not implemented")
}

func setReadBuffer(fd *netFD, bytes int) error {
	return errors.New("not implemented")
}

func setWriteBuffer(fd *netFD, bytes int) error {
	return errors.New("not implemented")
}

func (p *ipStackCapabilities) probe() {
	p.ipv4Enabled = true
}

func lookupProtocol(ctx context.Context, name string) (proto int, err error) {
	return lookupProtocolMap(name)
}

func (*Resolver) lookupHost(ctx context.Context, host string) (addrs []string, err error) {
	return nil, syscall.ENOPROTOOPT
}

func (*Resolver) lookupIP(ctx context.Context, network, host string) (addrs []IPAddr, err error) {
	return nil, syscall.ENOPROTOOPT
}

func (*Resolver) lookupPort(ctx context.Context, network, service string) (port int, err error) {
	return goLookupPort(network, service)
}

func (*Resolver) lookupCNAME(ctx context.Context, name string) (cname string, err error) {
	return "", syscall.ENOPROTOOPT
}

func (*Resolver) lookupSRV(ctx context.Context, service, proto, name string) (cname string, srvs []*SRV, err error) {
	return "", nil, syscall.ENOPROTOOPT
}

func (*Resolver) lookupMX(ctx context.Context, name string) (mxs []*MX, err error) {
	return nil, syscall.ENOPROTOOPT
}

func (*Resolver) lookupNS(ctx context.Context, name string) (nss []*NS, err error) {
	return nil, syscall.ENOPROTOOPT
}

func (*Resolver) lookupTXT(ctx context.Context, name string) (txts []string, err error) {
	return nil, syscall.ENOPROTOOPT
}

func (*Resolver) lookupAddr(ctx context.Context, addr string) (ptrs []string, err error) {
	return nil, syscall.ENOPROTOOPT
}
