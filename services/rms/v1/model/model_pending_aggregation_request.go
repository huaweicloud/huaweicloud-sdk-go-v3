package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// PendingAggregationRequest PendingAggregationRequest对象。
type PendingAggregationRequest struct {

	// 请求聚合数据的帐号ID。
	RequesterAccountId *string `json:"requester_account_id,omitempty"`
}

func (o PendingAggregationRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PendingAggregationRequest struct{}"
	}

	return strings.Join([]string{"PendingAggregationRequest", string(data)}, " ")
}
