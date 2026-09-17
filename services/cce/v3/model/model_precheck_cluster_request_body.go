package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type PrecheckClusterRequestBody struct {

	// **参数解释：** API版本，固定值\"v3\"，该值不可修改。 **约束限制：** 该值不可修改 **取值范围：** - v3  **默认取值：** v3
	ApiVersion string `json:"apiVersion"`

	// **参数解释：** API类型，固定值\"PreCheckTask\"，该值不可修改。 **约束限制：** 该值不可修改 **取值范围：** - PreCheckTask  **默认取值：** PreCheckTask
	Kind string `json:"kind"`

	Spec *PrecheckSpec `json:"spec"`
}

func (o PrecheckClusterRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PrecheckClusterRequestBody struct{}"
	}

	return strings.Join([]string{"PrecheckClusterRequestBody", string(data)}, " ")
}
