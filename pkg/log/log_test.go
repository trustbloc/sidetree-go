/*
 * Copyright Gen Digital Inc. All Rights Reserved.
 *
 * SPDX-License-Identifier: Apache-2.0
 */

package log

import (
	"bytes"
	"encoding/json"
	"errors"
	"log/slog"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestDefaultLogger(t *testing.T) {
	logger := New()

	logger.Error("Error message", slog.String("error", "some error"))
}

func TestCustomHandler(t *testing.T) {
	var buf bytes.Buffer

	SetHandler(slog.NewJSONHandler(&buf, &slog.HandlerOptions{
		AddSource: true,
		Level:     slog.LevelDebug,
	}))

	logger := New()

	err := errors.New("some error")

	logger.Debug("Debug message", slog.String("error", err.Error()))

	var doc map[string]interface{}
	require.NoError(t, json.Unmarshal(buf.Bytes(), &doc))

	msg, exists := doc["error"]
	require.True(t, exists)
	require.Equal(t, err.Error(), msg)

	msg, exists = doc["level"]
	require.True(t, exists)
	require.Equal(t, "DEBUG", msg)

	msg, exists = doc["msg"]
	require.True(t, exists)
	require.Equal(t, "Debug message", msg)
}
