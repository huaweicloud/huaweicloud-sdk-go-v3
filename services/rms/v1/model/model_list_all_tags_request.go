package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListAllTagsRequest Request Object
type ListAllTagsRequest struct {

	// 标签键名
	Key *string `json:"key,omitempty"`

	// 分页参数，通过上一个请求中返回的marker信息作为输入，获取当前页
	Marker *string `json:"marker,omitempty"`

	// 最大的返回数量。
	Limit *int32 `json:"limit,omitempty"`
}

func (o ListAllTagsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAllTagsRequest struct{}"
	}

	return strings.Join([]string{"ListAllTagsRequest", string(data)}, " ")
}
