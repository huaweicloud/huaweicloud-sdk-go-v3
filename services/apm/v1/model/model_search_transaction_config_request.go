package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SearchTransactionConfigRequest Request Object
type SearchTransactionConfigRequest struct {

	// 应用id。
	XBusinessId int64 `json:"x-business-id"`

	Body *TransactionConfigSearchRequest `json:"body,omitempty"`
}

func (o SearchTransactionConfigRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SearchTransactionConfigRequest struct{}"
	}

	return strings.Join([]string{"SearchTransactionConfigRequest", string(data)}, " ")
}
