package domain

import (
	"encoding/json"
	"fmt"
	"strings"
)

func CreateDocumentRequestFromDTE(codeSII, userID, businessID, idempotencyKey string, dte any) (CreateDocumentRequest, error) {
	if strings.TrimSpace(codeSII) == "" {
		return CreateDocumentRequest{}, fmt.Errorf("domain: code_sii is required")
	}

	raw, err := json.Marshal(dte)
	if err != nil {
		return CreateDocumentRequest{}, fmt.Errorf("domain: marshal dte data: %w", err)
	}

	return CreateDocumentRequest{
		UserID:         userID,
		BusinessID:     businessID,
		CodeSII:        codeSII,
		DataDTE:        string(raw),
		IdempotencyKey: idempotencyKey,
	}, nil
}

func DTE33ToRequest(userID, businessID, idempotencyKey string, dte Dte33Data) (CreateDocumentRequest, error) {
	return CreateDocumentRequestFromDTE("33", userID, businessID, idempotencyKey, dte)
}

func DTE34ToRequest(userID, businessID, idempotencyKey string, dte Dte34Data) (CreateDocumentRequest, error) {
	return CreateDocumentRequestFromDTE("34", userID, businessID, idempotencyKey, dte)
}

func DTE39ToRequest(userID, businessID, idempotencyKey string, dte Dte39Data) (CreateDocumentRequest, error) {
	return CreateDocumentRequestFromDTE("39", userID, businessID, idempotencyKey, dte)
}

func DTE41ToRequest(userID, businessID, idempotencyKey string, dte Dte41Data) (CreateDocumentRequest, error) {
	return CreateDocumentRequestFromDTE("41", userID, businessID, idempotencyKey, dte)
}

func DTE46ToRequest(userID, businessID, idempotencyKey string, dte Dte46Data) (CreateDocumentRequest, error) {
	return CreateDocumentRequestFromDTE("46", userID, businessID, idempotencyKey, dte)
}

func DTE52ToRequest(userID, businessID, idempotencyKey string, dte Dte52Data) (CreateDocumentRequest, error) {
	return CreateDocumentRequestFromDTE("52", userID, businessID, idempotencyKey, dte)
}

func DTE56ToRequest(userID, businessID, idempotencyKey string, dte Dte56Data) (CreateDocumentRequest, error) {
	return CreateDocumentRequestFromDTE("56", userID, businessID, idempotencyKey, dte)
}

func DTE61ToRequest(userID, businessID, idempotencyKey string, dte Dte61Data) (CreateDocumentRequest, error) {
	return CreateDocumentRequestFromDTE("61", userID, businessID, idempotencyKey, dte)
}
