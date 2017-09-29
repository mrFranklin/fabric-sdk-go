/*
Copyright SecureKey Technologies Inc. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package logbridge

import (
	"github.com/hyperledger/fabric-sdk-go/pkg/logging"
)

// Log levels (from fabric-sdk-go/pkg/logging/level.go).
const (
	CRITICAL logging.Level = iota
	ERROR
	WARNING
	INFO
	DEBUG
)

// Logger bridges the SDK's logger struct
type Logger struct {
	*logging.Logger
	module string
}

// MustGetLogger bridges calls the Go SDK NewLogger
func MustGetLogger(module string) *Logger {
	fabModule := "fabric_sdk_go"
	logger := logging.NewLogger(fabModule)
	return &Logger{
		Logger: logger,
		module: fabModule,
	}
}

// Warningf bridges calls to the Go SDK logger's Warnf.
func (l *Logger) Warningf(format string, args ...interface{}) {
	l.Warnf(format, args...)
}

// Warning bridges calls to the Go SDK logger's Warn.
func (l *Logger) Warning(format string, args ...interface{}) {
	l.Warn(args...)
}

// IsEnabledFor bridges calls to the Go SDK logger's IsEnabledFor.
func (l *Logger) IsEnabledFor(level logging.Level) bool {
	return logging.IsEnabledFor(level, l.module)
}