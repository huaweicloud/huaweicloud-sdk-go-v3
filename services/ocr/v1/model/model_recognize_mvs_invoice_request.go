package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type RecognizeMvsInvoiceRequest struct {
	Body *MvsInvoiceRequestBody `json:"body,omitempty" xml:"body"`
}

func (o RecognizeMvsInvoiceRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RecognizeMvsInvoiceRequest struct{}"
	}

	return strings.Join([]string{"RecognizeMvsInvoiceRequest", string(data)}, " ")
}
