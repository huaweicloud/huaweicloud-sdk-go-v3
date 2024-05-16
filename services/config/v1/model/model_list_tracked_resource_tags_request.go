package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListTrackedResourceTagsRequest Request Object
type ListTrackedResourceTagsRequest struct {

	// 标签键名
	Key *string `json:"key,omitempty"`

	// 分页参数，通过上一个请求中返回的marker信息作为输入，获取当前页
	Marker *string `json:"marker,omitempty"`

	// 最大的返回数量。
	Limit *int32 `json:"limit,omitempty"`

	// 是否查询已删除的资源。默认为false，不查询已删除的资源
	ResourceDeleted *bool `json:"resource_deleted,omitempty"`
}

func (o ListTrackedResourceTagsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListTrackedResourceTagsRequest struct{}"
	}

	return strings.Join([]string{"ListTrackedResourceTagsRequest", string(data)}, " ")
}
