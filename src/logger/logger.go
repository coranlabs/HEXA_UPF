// SPDX-License-Identifier: Apache-2.0
// Copyright 2024 CORAN LABS

package logger

import (
	"github.com/sirupsen/logrus"
)

var log *logrus.Logger

func init() {
	log = logrus.New()
	log.SetReportCaller(false)
}
