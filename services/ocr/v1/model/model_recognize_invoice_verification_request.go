package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type RecognizeInvoiceVerificationRequest struct {
	Body *InvoiceVerificationRequestBody `json:"body,omitempty"`
}

func (o RecognizeInvoiceVerificationRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "RecognizeInvoiceVerificationRequest struct{}"
	}

	return strings.Join([]string{"RecognizeInvoiceVerificationRequest", string(data)}, " ")
}
