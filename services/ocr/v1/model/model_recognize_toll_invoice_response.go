package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type RecognizeTollInvoiceResponse struct {
	Result         *TollInvoiceResult `json:"result,omitempty"`
	HttpStatusCode int                `json:"-"`
}

func (o RecognizeTollInvoiceResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "RecognizeTollInvoiceResponse struct{}"
	}

	return strings.Join([]string{"RecognizeTollInvoiceResponse", string(data)}, " ")
}
