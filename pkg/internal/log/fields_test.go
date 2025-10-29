/*
Copyright Gen Digital Inc. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
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

func TestStandardFields(t *testing.T) {
	var buf bytes.Buffer

	logger := slog.New(slog.NewJSONHandler(&buf, &slog.HandlerOptions{
		AddSource: true,
		Level:     slog.LevelDebug,
	}))

	o := struct {
		field string
	}{field: "value"}

	err := errors.New("some error")

	logger.Debug("Some message",
		WithData([]byte(`{"field":"value"}`)), WithSuffix("1234"), WithOperationType("Create"),
		WithSuffixes("123", "456"), WithError(err), WithIsBatch(true), WithOperation(o),
		WithTransactionTime(1234), WithTransactionNumber(6789), WithPatch(o), WithIsBatch(true),
	)

	t.Logf("%s", buf.String())
	l := unmarshalLogData(t, buf.Bytes())

	require.Equal(t, `Some message`, l.Msg)
	require.Equal(t, `{"field":"value"}`, l.Data)
	require.Equal(t, "1234", l.Suffix)
	require.Equal(t, "Create", l.OperationType)
	require.Equal(t, `Some message`, l.Msg)
	require.Len(t, l.Suffixes, 2)
	require.True(t, l.IsBatch)
	require.Equal(t, 1234, l.TransactionTime)
	require.Equal(t, 6789, l.TransactionNumber)
	require.Equal(t, err.Error(), l.Error)
}

type mockObject struct {
	Field1 string
	Field2 int
}

type logData struct {
	Level  string `json:"level"`
	Time   string `json:"time"`
	Logger string `json:"logger"`
	Caller string `json:"caller"`
	Msg    string `json:"msg"`
	Error  string `json:"error"`

	Data              string      `json:"data"`
	Suffix            string      `json:"suffix"`
	OperationType     string      `json:"operationType"`
	Suffixes          []string    `json:"suffixes"`
	Operation         *mockObject `json:"operation"`
	TransactionTime   int         `json:"transactionTime"`
	TransactionNumber int         `json:"transactionNumber"`
	Patch             *mockObject `json:"patch"`
	IsBatch           bool        `json:"isBatch"`
}

func unmarshalLogData(t *testing.T, b []byte) *logData {
	t.Helper()

	l := &logData{}

	require.NoError(t, json.Unmarshal(b, l))

	return l
}
