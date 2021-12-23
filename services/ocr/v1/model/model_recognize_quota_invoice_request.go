package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type RecognizeQuotaInvoiceRequest struct {
	Body *QuotaInvoiceRequestBody `json:"body,omitempty"`
}

func (o RecognizeQuotaInvoiceRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RecognizeQuotaInvoiceRequest struct{}"
	}

	return strings.Join([]string{"RecognizeQuotaInvoiceRequest", string(data)}, " ")
}
