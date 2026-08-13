package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CountAllResourcesRequest Request Object
type CountAllResourcesRequest struct {

	// 资源ID
	Id *string `json:"id,omitempty"`

	// 资源名称
	Name *string `json:"name,omitempty"`

	// 资源类型（provider.type）
	Type *[]string `json:"type,omitempty"`

	// 区域ID列表
	RegionId *[]string `json:"region_id,omitempty"`

	// 企业项目ID列表
	EpId *[]string `json:"ep_id,omitempty"`

	// 项目ID
	ProjectId *[]string `json:"project_id,omitempty"`

	// 标签列表
	Tags *[]string `json:"tags,omitempty"`
}

func (o CountAllResourcesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CountAllResourcesRequest struct{}"
	}

	return strings.Join([]string{"CountAllResourcesRequest", string(data)}, " ")
}
