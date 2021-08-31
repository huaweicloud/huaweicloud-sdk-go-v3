package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type RecognizeTaxiInvoiceRequest struct {
	Body *TaxiInvoiceRequestBody `json:"body,omitempty"`
}

func (o RecognizeTaxiInvoiceRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "RecognizeTaxiInvoiceRequest struct{}"
	}

	return strings.Join([]string{"RecognizeTaxiInvoiceRequest", string(data)}, " ")
}
