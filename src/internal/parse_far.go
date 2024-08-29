// SPDX-License-Identifier: Apache-2.0
// Copyright 2024 CORAN LABS

package internal

import (
	"github.com/coranlabs/HEXA_UPF/src/logger"
	infoElement "github.com/wmnsk/go-pfcp/ie"
)

func (s *Session) CreateFAR(lSeid uint64, req *infoElement.IE) error {
	logger.AppLog.Debug("create far function called")
	logger.AppLog.Debugf("fars extracted: %v", s.pdrs)
	return nil
}

func (s *Session) newForwardingParameter(ies []*infoElement.IE) error {}
