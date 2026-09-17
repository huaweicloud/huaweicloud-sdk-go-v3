package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListStoredQueriesRequest Request Object
type ListStoredQueriesRequest struct {

	// 最大的返回数量
	Limit *int32 `json:"limit,omitempty"`

	// 分页参数，通过上一个请求中返回的marker信息作为输入，获取当前页
	Marker *string `json:"marker,omitempty"`

	// ResourceQL 名字
	Name *string `json:"name,omitempty"`
}

func (o ListStoredQueriesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListStoredQueriesRequest struct{}"
	}

	return strings.Join([]string{"ListStoredQueriesRequest", string(data)}, " ")
}
