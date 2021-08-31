package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type RecognizeVatInvoiceRequest struct {
	Body *VatInvoiceRequestBody `json:"body,omitempty"`
}

func (o RecognizeVatInvoiceRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "RecognizeVatInvoiceRequest struct{}"
	}

	return strings.Join([]string{"RecognizeVatInvoiceRequest", string(data)}, " ")
}
