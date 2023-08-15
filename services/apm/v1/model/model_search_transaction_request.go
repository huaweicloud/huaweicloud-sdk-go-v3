package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SearchTransactionRequest Request Object
type SearchTransactionRequest struct {

	// 应用id。
	XBusinessId int64 `json:"x-business-id"`

	Body *TxSearchRequest `json:"body,omitempty"`
}

func (o SearchTransactionRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SearchTransactionRequest struct{}"
	}

	return strings.Join([]string{"SearchTransactionRequest", string(data)}, " ")
}
