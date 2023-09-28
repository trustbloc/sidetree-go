/*
   Copyright SecureKey Technologies Inc.

   This file contains software code that is the intellectual property of SecureKey.
   SecureKey reserves all rights in the code and you may not use it without
	 written permission from SecureKey.
*/

package docutil

import (
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/trustbloc/sidetree-go/pkg/api/operation"
	"github.com/trustbloc/sidetree-go/pkg/api/protocol"
	"github.com/trustbloc/sidetree-go/pkg/document"
)

// GetTransformationInfoForPublished will create transformation info object for published document.
func GetTransformationInfoForPublished(namespace, id, suffix string,
	internalResult *protocol.ResolutionModel) protocol.TransformationInfo {
	ti := make(protocol.TransformationInfo)
	ti[document.IDProperty] = id
	ti[document.PublishedProperty] = true

	canonicalRef := ""
	if internalResult.CanonicalReference != "" {
		canonicalRef = NamespaceDelimiter + internalResult.CanonicalReference
	}

	canonicalID := namespace + canonicalRef + NamespaceDelimiter + suffix

	// we should always set canonical id if document has been published
	ti[document.CanonicalIDProperty] = canonicalID

	equivalentIDs := []string{canonicalID}

	if len(internalResult.EquivalentReferences) > 0 {
		for _, eqRef := range internalResult.EquivalentReferences {
			equivalentID := namespace + NamespaceDelimiter + eqRef + NamespaceDelimiter + suffix
			equivalentIDs = append(equivalentIDs, equivalentID)
		}
	}

	// equivalent ids should always include canonical id (if specified)
	ti[document.EquivalentIDProperty] = equivalentIDs

	return ti
}

// GetTransformationInfoForUnpublished will create transformation info object for unpublished document.
func GetTransformationInfoForUnpublished(namespace, domain, label, suffix, createRequestJCS string) protocol.TransformationInfo {
	ti := make(protocol.TransformationInfo)
	ti[document.PublishedProperty] = false

	id := fmt.Sprintf("%s:%s", namespace, suffix)

	// For interim/unpublished documents we should set optional label if specified.
	if label != "" {
		id = fmt.Sprintf("%s:%s:%s", namespace, label, suffix)
	}

	var equivalentIDs []string

	if createRequestJCS != "" {
		// we should always set short form equivalent id for long form resolution
		equivalentIDs = append(equivalentIDs, id)
	}

	// Also, if optional domain is specified, we should set equivalent id with domain hint
	if label != "" && domain != "" {
		equivalentID := id
		if !strings.Contains(label, domain) {
			equivalentID = fmt.Sprintf("%s:%s:%s:%s", namespace, domain, label, suffix)
		}

		equivalentIDs = append(equivalentIDs, equivalentID)
	}

	if len(equivalentIDs) > 0 {
		ti[document.EquivalentIDProperty] = equivalentIDs
	}

	if createRequestJCS != "" {
		id = fmt.Sprintf("%s:%s", id, createRequestJCS)
	}

	ti[document.IDProperty] = id

	return ti
}

func GetCreateResult(op *operation.Operation, pv protocol.Version) (*protocol.ResolutionModel, error) {
	// we can use operation applier to generate create response even though operation is not anchored yet
	anchored := &operation.AnchoredOperation{
		Type:             op.Type,
		UniqueSuffix:     op.UniqueSuffix,
		OperationRequest: op.OperationRequest,
		TransactionTime:  uint64(time.Now().Unix()),
		ProtocolVersion:  pv.Protocol().GenesisTime,
		AnchorOrigin:     op.AnchorOrigin,
	}

	rm := &protocol.ResolutionModel{UnpublishedOperations: []*operation.AnchoredOperation{anchored}}
	rm, err := pv.OperationApplier().Apply(anchored, rm)
	if err != nil {
		return nil, err
	}

	// if returned document is empty (e.g. applying patches failed) we can reject this request at API level
	if len(rm.Doc.JSONLdObject()) == 0 {
		return nil, errors.New("applying delta resulted in an empty document (most likely due to an invalid patch)")
	}

	return rm, nil
}
