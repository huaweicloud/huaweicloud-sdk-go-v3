package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type RecognizeTaxiInvoiceResponse struct {
	Result         *TaxiInvoiceResult `json:"result,omitempty"`
	HttpStatusCode int                `json:"-"`
}

func (o RecognizeTaxiInvoiceResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "RecognizeTaxiInvoiceResponse struct{}"
	}

	return strings.Join([]string{"RecognizeTaxiInvoiceResponse", string(data)}, " ")
}
