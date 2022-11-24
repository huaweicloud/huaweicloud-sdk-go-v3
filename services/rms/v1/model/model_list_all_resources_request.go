package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ListAllResourcesRequest struct {

	// 区域ID
	RegionId *string `json:"region_id,omitempty"`

	// 企业项目ID
	EpId *string `json:"ep_id,omitempty"`

	// 资源类型（provider.type）
	Type *string `json:"type,omitempty"`

	// 最大的返回数量
	Limit *int32 `json:"limit,omitempty"`

	// 分页参数，通过上一个请求中返回的marker信息作为输入，获取当前页
	Marker *string `json:"marker,omitempty"`
}

func (o ListAllResourcesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAllResourcesRequest struct{}"
	}

	return strings.Join([]string{"ListAllResourcesRequest", string(data)}, " ")
}
