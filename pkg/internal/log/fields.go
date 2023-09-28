/*
Copyright Gen Digital Inc. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package log

import (
	"log/slog"
)

// Log Fields.
const (
	FieldData              = "data"
	FieldSuffix            = "suffix"
	FieldSuffixes          = "suffixes"
	FieldOperationType     = "operationType"
	FieldOperation         = "operation"
	FieldTransactionTime   = "transactionTime"
	FieldTransactionNumber = "transactionNumber"
	FieldPatch             = "patch"
	FieldIsBatch           = "isBatch"
)

// WithError sets the error field.
func WithError(err error) slog.Attr {
	return slog.String("error", err.Error())
}

// WithData sets the data field.
func WithData(value []byte) slog.Attr {
	return slog.String(FieldData, string(value))
}

// WithSuffix sets the suffix field.
func WithSuffix(value string) slog.Attr {
	return slog.String(FieldSuffix, value)
}

// WithSuffixes sets the suffixes field.
func WithSuffixes(value ...string) slog.Attr {
	return slog.Any(FieldSuffixes, value)
}

// WithOperationType sets the operation-type field.
func WithOperationType(value string) slog.Attr {
	return slog.Any(FieldOperationType, value)
}

// WithOperation sets the operation field.
func WithOperation(value interface{}) slog.Attr {
	return slog.Any(FieldOperation, value)
}

// WithTransactionTime sets the transaction-time field.
func WithTransactionTime(value uint64) slog.Attr {
	return slog.Uint64(FieldTransactionTime, value)
}

// WithTransactionNumber sets the transaction-number field.
func WithTransactionNumber(value uint64) slog.Attr {
	return slog.Uint64(FieldTransactionNumber, value)
}

// WithPatch sets the patch field.
func WithPatch(value interface{}) slog.Attr {
	return slog.Any(FieldPatch, value)
}

// WithIsBatch sets the is-batch field.
func WithIsBatch(value bool) slog.Attr {
	return slog.Bool(FieldIsBatch, value)
}
