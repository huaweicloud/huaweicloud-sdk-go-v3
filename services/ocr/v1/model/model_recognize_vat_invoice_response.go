package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type RecognizeVatInvoiceResponse struct {
	Result         *VatInvoiceResult `json:"result,omitempty"`
	HttpStatusCode int               `json:"-"`
}

func (o RecognizeVatInvoiceResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RecognizeVatInvoiceResponse struct{}"
	}

	return strings.Join([]string{"RecognizeVatInvoiceResponse", string(data)}, " ")
}
