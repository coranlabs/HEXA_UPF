// SPDX-License-Identifier: Apache-2.0
// Copyright 2024 CORAN LABS

package internal

import (
	"github.com/coranlabs/HEXA_UPF/src/logger"
)

func (s *Session) CreatePDR() error {

	logger.AppLog.Debug("create pdr function called")
	logger.AppLog.Debugf("pdrs extracted: %v", s.pdrs)
	return nil
}
