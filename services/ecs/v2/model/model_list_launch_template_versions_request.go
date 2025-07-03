package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListLaunchTemplateVersionsRequest Request Object
type ListLaunchTemplateVersionsRequest struct {

	// max number of resources to return
	Limit *int32 `json:"limit,omitempty"`

	// next page resource index id
	Marker *string `json:"marker,omitempty"`

	// 模板id
	LaunchTemplateId *string `json:"launch_template_id,omitempty"`

	// 镜像id
	ImageId *string `json:"image_id,omitempty"`

	// 规格id
	FlavorId *string `json:"flavor_id,omitempty"`

	// 版本
	Version *[]int32 `json:"version,omitempty"`
}

func (o ListLaunchTemplateVersionsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListLaunchTemplateVersionsRequest struct{}"
	}

	return strings.Join([]string{"ListLaunchTemplateVersionsRequest", string(data)}, " ")
}
