package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ListAggregationAuthorizationsRequest struct {

	// 授权的帐号ID。
	AccountId *string `json:"account_id,omitempty"`

	// 最大的返回数量
	Limit *int32 `json:"limit,omitempty"`

	// 分页参数，通过上一个请求中返回的marker信息作为输入，获取当前页
	Marker *string `json:"marker,omitempty"`
}

func (o ListAggregationAuthorizationsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAggregationAuthorizationsRequest struct{}"
	}

	return strings.Join([]string{"ListAggregationAuthorizationsRequest", string(data)}, " ")
}
