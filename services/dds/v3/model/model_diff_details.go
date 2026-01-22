package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type DiffDetails struct {

	// **参数解释：** 参数名称。 **取值范围：** 不涉及。
	ParameterName string `json:"parameter_name"`

	// **参数解释：** 源参数模板中的参数值。 **取值范围：** 不涉及。
	SourceValue string `json:"source_value"`

	// **参数解释：** 目标参数模板中的参数值。 **取值范围：** 不涉及。
	TargetValue string `json:"target_value"`
}

func (o DiffDetails) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DiffDetails struct{}"
	}

	return strings.Join([]string{"DiffDetails", string(data)}, " ")
}
