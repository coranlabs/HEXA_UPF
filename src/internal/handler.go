// SPDX-License-Identifier: Apache-2.0
// Copyright 2024 CORAN LABS

package internal

import (
	"log"
	"net"

	mes "github.com/wmnsk/go-pfcp/message"
)

func Handle(conn *PfcpConn, buf []byte, addr *net.UDPAddr) error {
	msg, err := mes.Parse(buf)
	if err != nil {

		return err
	}

	for {
		switch msg.MessageType() {
		case mes.MsgTypeAssociationSetupRequest:
			HandlePfcpAssociationSetupRequest(conn, msg)
		default:
			log.Printf("Got unexpected mes %s: %s, from: %s", msg.MessageTypeName(), msg, addr)
			return nil
		}
	}
}

func HandlePfcpAssociationSetupRequest(conn *PfcpConn, msg mes.Message) {}

func NewNodeAssociation(remoteNodeID string, addr string) *NodeAssociation {
	return &NodeAssociation{}
}

func (association *NodeAssociation) ScheduleHeartbeat(conn *PfcpConn) {}

func (connection *PfcpConn) SendMessage(msg mes.Message, addr *net.UDPAddr) error {
	return nil
}
