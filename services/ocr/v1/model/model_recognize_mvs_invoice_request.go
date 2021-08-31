package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type RecognizeMvsInvoiceRequest struct {
	Body *MvsInvoiceRequestBody `json:"body,omitempty"`
}

func (o RecognizeMvsInvoiceRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "RecognizeMvsInvoiceRequest struct{}"
	}

	return strings.Join([]string{"RecognizeMvsInvoiceRequest", string(data)}, " ")
}
