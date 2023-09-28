/*
 * Copyright Gen Digital Inc. All Rights Reserved.
 *
 * SPDX-License-Identifier: Apache-2.0
 */

package log

import (
	"log/slog"
	"os"
	"sync"
)

var mutex sync.RWMutex

var handler slog.Handler = slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelError})

// New returns a new logger.
func New() *slog.Logger {
	mutex.RLock()
	defer mutex.RUnlock()

	return slog.New(handler)
}

// SetHandler sets the slog handler.
func SetHandler(h slog.Handler) {
	mutex.Lock()
	handler = h
	mutex.Unlock()
}
