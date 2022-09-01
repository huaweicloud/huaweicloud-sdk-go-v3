package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type RecognizeQuotaInvoiceResponse struct {
	Result         *QuotaInvoiceResult `json:"result,omitempty" xml:"result"`
	HttpStatusCode int                 `json:"-"`
}

func (o RecognizeQuotaInvoiceResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RecognizeQuotaInvoiceResponse struct{}"
	}

	return strings.Join([]string{"RecognizeQuotaInvoiceResponse", string(data)}, " ")
}
