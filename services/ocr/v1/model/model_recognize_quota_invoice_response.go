package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type RecognizeQuotaInvoiceResponse struct {
	Result         *QuotaInvoiceResult `json:"result,omitempty"`
	HttpStatusCode int                 `json:"-"`
}

func (o RecognizeQuotaInvoiceResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "RecognizeQuotaInvoiceResponse struct{}"
	}

	return strings.Join([]string{"RecognizeQuotaInvoiceResponse", string(data)}, " ")
}
