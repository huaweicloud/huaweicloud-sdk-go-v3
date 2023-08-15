package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RecognizeTollInvoiceResponse Response Object
type RecognizeTollInvoiceResponse struct {
	Result         *TollInvoiceResult `json:"result,omitempty"`
	HttpStatusCode int                `json:"-"`
}

func (o RecognizeTollInvoiceResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RecognizeTollInvoiceResponse struct{}"
	}

	return strings.Join([]string{"RecognizeTollInvoiceResponse", string(data)}, " ")
}
