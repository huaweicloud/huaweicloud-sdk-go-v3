package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ResourcesInListResp struct {

	// 资源分组ID，监控范围为资源分组时存在该值
	ResourceGroupId *string `json:"resource_group_id,omitempty"`

	// 资源分组名称，监控范围为资源分组时存在该值
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
