package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type RecognizeQuotaInvoiceRequest struct {
	Body *QuotaInvoiceRequestBody `json:"body,omitempty"`
}

func (o RecognizeQuotaInvoiceRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "RecognizeQuotaInvoiceRequest struct{}"
	}

	return strings.Join([]string{"RecognizeQuotaInvoiceRequest", string(data)}, " ")
}
