package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type RecognizeQuotaInvoiceResponse struct {
	Result         *QuotaInvoiceResult `json:"result,omitempty"`
	HttpStatusCode int                 `json:"-"`
}

func (o RecognizeQuotaInvoiceResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RecognizeQuotaInvoiceResponse struct{}"
	}

	return strings.Join([]string{"RecognizeQuotaInvoiceResponse", string(data)}, " ")
}
