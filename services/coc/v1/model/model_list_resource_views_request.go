package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListResourceViewsRequest Request Object
type ListResourceViewsRequest struct {

	// 用于分页查询，查询数量，最大的返回数量。取值范围：1-500。
	Limit int32 `json:"limit"`

	// 用于分页查询。分页参数，通过上一个请求中返回的marker信息作为输入，获取当前页。
	Marker *string `json:"marker,omitempty"`
}

func (o ListResourceViewsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListResourceViewsRequest struct{}"
	}

	return strings.Join([]string{"ListResourceViewsRequest", string(data)}, " ")
}
