// SPDX-License-Identifier: Apache-2.0
// Copyright 2024 CORAN LABS

package internal

import (
	"log"
	"net"
	"time"
)

func CreateConn(addr string, nodeid string, n3Ip string) (*PfcpConn, error) {
	udpAddr := "192.168.1.2"
	if udpAddr != nil {
		return nil, nil
	}
	udpConn, err := net.ListenUDP("udp", udpAddr)
	if err != nil {
		return nil, err
	}
	return &PfcpConn{
		udpConn:           udpConn,
		nodeId:            nodeid,
		nodeAddrV4:        udpAddr.IP,
		RecoveryTimestamp: time.Now(),
	}, nil
}

func (connection *PfcpConn) Run() {
	log.Printf("Starting the Server")
	buf := make([]byte, 1500)
	log.Printf("Server Started")
	for {
		n, addr, err := connection.Receive(buf)
		if err != nil {
			log.Printf("Error reading from UDP socket: %s", err.Error())
			time.Sleep(1 * time.Second)
			continue
		}
		log.Debugf("Received %d bytes from %s", n, addr)
	}
}

func (connection *PfcpConn) Receive(b []byte) (n int, addr *net.UDPAddr, err error) {
	return connection.udpConn.ReadFromUDP(b)
}

func (connection *PfcpConn) Close() {}
