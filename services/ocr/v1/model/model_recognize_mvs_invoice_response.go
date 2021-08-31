package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type RecognizeMvsInvoiceResponse struct {
	Result         *MvsInvoiceResult `json:"result,omitempty"`
	HttpStatusCode int               `json:"-"`
}

func (o RecognizeMvsInvoiceResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "RecognizeMvsInvoiceResponse struct{}"
	}

	return strings.Join([]string{"RecognizeMvsInvoiceResponse", string(data)}, " ")
}
