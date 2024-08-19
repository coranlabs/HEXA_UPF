// SPDX-License-Identifier: Apache-2.0
// Copyright 2024 CORAN LABS

package internal

import (
	"time"

	"github.com/coranlabs/HEXA_UPF/src/logger"
)

func (association *NodeAssociation) ScheduleHeartbeat(conn *PfcpConn) {
	association.HeartbeatsActive = true
	logger.AppLog.Infof("Heartbeat started for SMF with nodeid %s", association.Addr)
	for {
		SendHeartbeatRequest(conn, sequence, association.Addr)
		time.Sleep(5 * time.Second)
	}
}
func SendHeartbeatRequest(conn *PfcpConn, sequenceID uint32, associationAddr string) {}

func (association *NodeAssociation) HandleHeartbeatTimeout() bool {}
