/*
Copyright Gen Digital Inc. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package operation

// Property contains a key-value pair.
type Property struct {
	Key   string
	Value interface{}
}

// Operation holds minimum information required for parsing/validating client request.
type Operation struct {

	// Type defines operation type.
	Type Type

	// UniqueSuffix defines document unique suffix.
	UniqueSuffix string

	// ID defines ID
	ID string

	// OperationRequest is the original operation request
	OperationRequest []byte

	// AnchorOrigin defines anchor origin.
	AnchorOrigin interface{}

	// Properties contains an arbitrary set of implementation-specific name-value pairs.
	Properties []Property
}

// AnchoredOperation defines an anchored operation (stored in document operation store).
type AnchoredOperation struct {

	// Type defines operation type.
	Type Type `json:"type"`

	// UniqueSuffix defines document unique suffix.
	UniqueSuffix string `json:"uniqueSuffix"`

	// OperationRequest is the original operation request
	OperationRequest []byte `json:"operation"`

	// TransactionTime is the logical anchoring time (block number in case of blockchain) for this operation in the
	// anchoring system (blockchain).
	TransactionTime uint64 `json:"transactionTime"`

	// TransactionNumber is the transaction number of the transaction this operation was batched within.
	TransactionNumber uint64 `json:"transactionNumber"`

	// ProtocolVersion is the genesis time (version) of the protocol that was used for this operation.
	ProtocolVersion uint64 `json:"protocolVersion"`

	// CanonicalReference contains canonical reference that applies to this operation.
	CanonicalReference string `json:"canonicalReference,omitempty"`

	// EquivalenceReferences contains equivalence reference that applies to this operation.
	EquivalentReferences []string `json:"equivalentReferences,omitempty"`

	// AnchorOrigin is anchor origin
	AnchorOrigin interface{} `json:"anchorOrigin,omitempty"`
}

// Type defines valid values for operation type.
type Type string

const (
	// TypeCreate captures "create" operation type.
	TypeCreate Type = "create"

	// TypeUpdate captures "update" operation type.
	TypeUpdate Type = "update"

	// TypeDeactivate captures "deactivate" operation type.
	TypeDeactivate Type = "deactivate"

	// TypeRecover captures "recover" operation type.
	TypeRecover Type = "recover"
)
