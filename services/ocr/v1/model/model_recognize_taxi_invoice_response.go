package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type RecognizeTaxiInvoiceResponse struct {
	Result         *TaxiInvoiceResult `json:"result,omitempty"`
	HttpStatusCode int                `json:"-"`
}

func (o RecognizeTaxiInvoiceResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RecognizeTaxiInvoiceResponse struct{}"
	}

	return strings.Join([]string{"RecognizeTaxiInvoiceResponse", string(data)}, " ")
}
