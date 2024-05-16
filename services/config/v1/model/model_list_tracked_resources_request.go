package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListTrackedResourcesRequest Request Object
type ListTrackedResourcesRequest struct {

	// 区域ID
	RegionId *string `json:"region_id,omitempty"`

	// 企业项目ID
	EpId *string `json:"ep_id,omitempty"`

	// 资源类型（provider.type）
	Type *string `json:"type,omitempty"`

	// 最大的返回数量。
	Limit *int32 `json:"limit,omitempty"`

	// 分页参数，通过上一个请求中返回的marker信息作为输入，获取当前页
	Marker *string `json:"marker,omitempty"`

	// 资源ID
	Id *string `json:"id,omitempty"`

	// 资源名称
	Name *string `json:"name,omitempty"`

	// 标签列表
	Tags *[]string `json:"tags,omitempty"`

	// 是否查询已删除的资源。默认为false，不查询已删除的资源
	ResourceDeleted *bool `json:"resource_deleted,omitempty"`
}

func (o ListTrackedResourcesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListTrackedResourcesRequest struct{}"
	}

	return strings.Join([]string{"ListTrackedResourcesRequest", string(data)}, " ")
}
