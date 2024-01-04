package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchDeleteSubscriptionsFilterPolicesResponse Response Object
type BatchDeleteSubscriptionsFilterPolicesResponse struct {

	// 请求的唯一标识ID。
	RequestId *string `json:"request_id,omitempty"`

	// 批量结果
	BatchResult    *[]BatchResult `json:"batch_result,omitempty"`
	HttpStatusCode int            `json:"-"`
}

func (o BatchDeleteSubscriptionsFilterPolicesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDeleteSubscriptionsFilterPolicesResponse struct{}"
	}

	return strings.Join([]string{"BatchDeleteSubscriptionsFilterPolicesResponse", string(data)}, " ")
}
