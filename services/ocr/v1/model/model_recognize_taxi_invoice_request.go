package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type RecognizeTaxiInvoiceRequest struct {
	Body *TaxiInvoiceRequestBody `json:"body,omitempty"`
}

func (o RecognizeTaxiInvoiceRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RecognizeTaxiInvoiceRequest struct{}"
	}

	return strings.Join([]string{"RecognizeTaxiInvoiceRequest", string(data)}, " ")
}
