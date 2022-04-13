package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type RecognizeInvoiceVerificationRequest struct {
	Body *InvoiceVerificationRequestBody `json:"body,omitempty"`
}

func (o RecognizeInvoiceVerificationRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RecognizeInvoiceVerificationRequest struct{}"
	}

	return strings.Join([]string{"RecognizeInvoiceVerificationRequest", string(data)}, " ")
}
