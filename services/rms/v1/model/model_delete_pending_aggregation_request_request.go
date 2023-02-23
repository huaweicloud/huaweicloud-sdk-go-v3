package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type DeletePendingAggregationRequestRequest struct {

	// 请求聚合数据的帐号ID。
	RequesterAccountId string `json:"requester_account_id"`
}

func (o DeletePendingAggregationRequestRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeletePendingAggregationRequestRequest struct{}"
	}

	return strings.Join([]string{"DeletePendingAggregationRequestRequest", string(data)}, " ")
}
