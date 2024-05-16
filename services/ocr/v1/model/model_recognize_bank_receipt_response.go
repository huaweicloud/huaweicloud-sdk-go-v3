package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RecognizeBankReceiptResponse Response Object
type RecognizeBankReceiptResponse struct {
	Result *BankReceiptResult `json:"result,omitempty"`

	XRequestId     *string `json:"X-Request-Id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o RecognizeBankReceiptResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RecognizeBankReceiptResponse struct{}"
	}

	return strings.Join([]string{"RecognizeBankReceiptResponse", string(data)}, " ")
}
