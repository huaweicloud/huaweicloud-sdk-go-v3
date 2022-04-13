package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ResourcesInListResp struct {
	// 资源分组ID

	ResourceGroupId *string `json:"resource_group_id,omitempty"`
	// 资源分组名称

	ResourceGroupName *string `json:"resource_group_name,omitempty"`
	// 维度信息

	Dimensions *[]MetricDimension `json:"dimensions,omitempty"`
}

func (o ResourcesInListResp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResourcesInListResp struct{}"
	}

	return strings.Join([]string{"ResourcesInListResp", string(data)}, " ")
}
