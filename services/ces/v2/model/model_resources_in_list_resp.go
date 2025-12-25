package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ResourcesInListResp struct {

	// **参数解释**： 资源分组ID，监控范围为资源分组时，存在该值。 **取值范围**： 以rg开头，后跟22个字母或数字。
	ResourceGroupId *string `json:"resource_group_id,omitempty"`

	// **参数解释**： 资源分组名称，监控范围为资源分组时，存在该值。 **取值范围**： 长度为[1,128]个字符。
	ResourceGroupName *string `json:"resource_group_name,omitempty"`

	// **参数解释**： 维度信息。
	Dimensions *[]MetricDimensionResp `json:"dimensions,omitempty"`
}

func (o ResourcesInListResp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResourcesInListResp struct{}"
	}

	return strings.Join([]string{"ResourcesInListResp", string(data)}, " ")
}
