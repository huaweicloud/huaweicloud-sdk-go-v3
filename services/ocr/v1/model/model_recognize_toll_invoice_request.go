package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type RecognizeTollInvoiceRequest struct {
	Body *TollInvoiceRequestBody `json:"body,omitempty"`
}

func (o RecognizeTollInvoiceRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "RecognizeTollInvoiceRequest struct{}"
	}

	return strings.Join([]string{"RecognizeTollInvoiceRequest", string(data)}, " ")
}
