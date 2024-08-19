// SPDX-License-Identifier: Apache-2.0
// Copyright 2024 CORAN LABS

package internal

import (
	"log"
	"net"

	"github.com/coranlabs/HEXA_UPF/src/logger"
	infoElement "github.com/wmnsk/go-pfcp/ie"
	mes "github.com/wmnsk/go-pfcp/message"
)

func Handle(conn *PfcpConn, buf []byte, addr *net.UDPAddr) error {
	msg, err := mes.Parse(buf)
	stringIpAddr := addr.IP.String()
	if err != nil {

		return err
	}

	for {
		switch msg.MessageType() {
		case mes.MsgTypeAssociationSetupRequest:
			Msg, err := HandlePfcpAssociationSetupRequest(conn, msg, stringIpAddr)
			if err != nil {

				return err
			}
			return conn.SendMessage(Msg, addr)
		default:
			log.Printf("Got unexpected mes %s: %s, from: %s", msg.MessageTypeName(), msg, addr)
			return nil
		}
	}
}

func HandlePfcpAssociationSetupRequest(conn *PfcpConn, msg mes.Message, addr string) (mes.Message, error) {
	asreq := msg.(*mes.AssociationSetupRequest)
	remoteNodeID, err := asreq.NodeID.NodeID()
	logger.AppLog.Infof("Handling association Setup Request from: %s ", addr)
	logger.AppLog.Infof("nodeip: %s", remoteNodeID)
	if err != nil {
		logger.AppLog.Infof("Got Association Setup Request with invalid NodeID from: %s", addr)
		asres := mes.NewAssociationSetupResponse(asreq.SequenceNumber,
			infoElement.NewCause(infoElement.CauseMandatoryIEMissing),
		)
		return asres, nil
	}

	res := mes.NewAssociationSetupResponse(asreq.SequenceNumber,
		infoElement.NewCause(infoElement.CauseRequestRejected),
		infoElement.NewRecoveryTimeStamp(conn.RecoveryTimestamp),
	)
	return res, nil
}

func NewNodeAssociation(remoteNodeID string, addr string) *NodeAssociation {
	return &NodeAssociation{
		ID:               remoteNodeID,
		Addr:             addr,
		NextSessionID:    1,
		NextSequenceID:   1,
		Sessions:         make(map[uint64]*Session),
		HeartbeatChannel: make(chan uint32),
	}
}

func (connection *PfcpConn) SendMessage(msg mes.Message, addr *net.UDPAddr) error {
	responseBytes := make([]byte, msg.MarshalLen())
	if err := msg.MarshalTo(responseBytes); err != nil {
		logger.AppLog.Infof(err.Error())
		return err
	}
	if _, err := connection.Send(responseBytes, addr); err != nil {
		logger.AppLog.Infof(err.Error())
		return err
	}
	return nil
}

func (connection *PfcpConn) Send(b []byte, addr *net.UDPAddr) (int, error) {
	return connection.udpConn.WriteTo(b, addr)
}
