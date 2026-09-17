package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CreateUpgradeWorkFlowRequestBody struct {

	// **参数解释：** API类型，固定值\"WorkFlowTask\"，该值不可修改。 **约束限制：** 该值不可修改 **取值范围：** - WorkFlowTask  **默认取值：** WorkFlowTask
	Kind string `json:"kind"`

	// **参数解释：** API版本，固定值\"v3\"，该值不可修改。 **约束限制：** 该值不可修改 **取值范围：** - v3  **默认取值：** v3
	ApiVersion string `json:"apiVersion"`

	Spec *WorkFlowSpec `json:"spec"`
}

func (o CreateUpgradeWorkFlowRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateUpgradeWorkFlowRequestBody struct{}"
	}

	return strings.Join([]string{"CreateUpgradeWorkFlowRequestBody", string(data)}, " ")
}
