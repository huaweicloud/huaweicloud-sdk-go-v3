package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowAssociatedResourcesRequest Request Object
type ShowAssociatedResourcesRequest struct {

	// 项目ID
	ProjectId string `json:"project_id"`

	// 区域ID
	RegionId string `json:"region_id"`

	// 资源ID
	ResourceId string `json:"resource_id"`

	// 资源类型
	ResourceType string `json:"resource_type"`
}

func (o ShowAssociatedResourcesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowAssociatedResourcesRequest struct{}"
	}

	return strings.Join([]string{"ShowAssociatedResourcesRequest", string(data)}, " ")
}
