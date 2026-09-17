package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type PostcheckClusterRequestBody struct {

	// **参数解释：** API版本，固定值\"v3\"，该值不可修改。 **约束限制：** 该值不可修改 **取值范围：** - v3  **默认取值：** v3
	ApiVersion string `json:"apiVersion"`

	// **参数解释：** API类型，固定值\"PostCheckTask\"，该值不可修改。 **约束限制：** 该值不可修改 **取值范围：** - PostCheckTask  **默认取值：** PostCheckTask
	Kind string `json:"kind"`

	Spec *PostcheckSpec `json:"spec"`
}

func (o PostcheckClusterRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PostcheckClusterRequestBody struct{}"
	}

	return strings.Join([]string{"PostcheckClusterRequestBody", string(data)}, " ")
}
