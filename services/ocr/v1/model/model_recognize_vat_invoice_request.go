package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type RecognizeVatInvoiceRequest struct {
	Body *VatInvoiceRequestBody `json:"body,omitempty"`
}

func (o RecognizeVatInvoiceRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RecognizeVatInvoiceRequest struct{}"
	}

	return strings.Join([]string{"RecognizeVatInvoiceRequest", string(data)}, " ")
}
