package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RecognizeMvsInvoiceResponse Response Object
type RecognizeMvsInvoiceResponse struct {
	Result         *MvsInvoiceResult `json:"result,omitempty"`
	HttpStatusCode int               `json:"-"`
}

func (o RecognizeMvsInvoiceResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RecognizeMvsInvoiceResponse struct{}"
	}

	return strings.Join([]string{"RecognizeMvsInvoiceResponse", string(data)}, " ")
}
