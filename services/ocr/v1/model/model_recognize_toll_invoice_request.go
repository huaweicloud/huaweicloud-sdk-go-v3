package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type RecognizeTollInvoiceRequest struct {
	Body *TollInvoiceRequestBody `json:"body,omitempty"`
}

func (o RecognizeTollInvoiceRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RecognizeTollInvoiceRequest struct{}"
	}

	return strings.Join([]string{"RecognizeTollInvoiceRequest", string(data)}, " ")
}
