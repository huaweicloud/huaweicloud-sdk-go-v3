package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RecognizeVatInvoiceResponse Response Object
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
