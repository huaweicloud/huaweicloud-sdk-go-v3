package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowTransactionDetailRequest Request Object
type ShowTransactionDetailRequest struct {

	// 应用id。
	XBusinessId int64 `json:"x-business-id"`

	Body *TxDetailRequest `json:"body,omitempty"`
}

func (o ShowTransactionDetailRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowTransactionDetailRequest struct{}"
	}

	return strings.Join([]string{"ShowTransactionDetailRequest", string(data)}, " ")
}
