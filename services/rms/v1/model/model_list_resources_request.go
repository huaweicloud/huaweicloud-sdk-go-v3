package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ListResourcesRequest struct {
	// 云服务英文简写

	Provider string `json:"provider"`
	// 云服务资源类型名称

	Type string `json:"type"`
	// 区域ID

	RegionId *string `json:"region_id,omitempty"`
	// 企业项目ID

	EpId *string `json:"ep_id,omitempty"`
	// 最大的返回数量

	Limit *int32 `json:"limit,omitempty"`
	// 分页参数，通过上一个请求中返回的marker信息作为输入，获取当前页

	Marker *string `json:"marker,omitempty"`
}

func (o ListResourcesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListResourcesRequest struct{}"
	}

	return strings.Join([]string{"ListResourcesRequest", string(data)}, " ")
}
