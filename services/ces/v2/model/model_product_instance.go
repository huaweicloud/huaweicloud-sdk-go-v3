package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ProductInstance struct {

	// **参数解释** 资源首层维度，如：弹性云服务器，则维度为instance_id **约束限制** 不涉及 **取值范围** 资源维度名称，必须以字母开头，只能包含0-9/a-z/A-Z/_/-，维度的最大长度为32。 **默认取值** 不涉及
	FirstDimensionName string `json:"first_dimension_name"`

	// **参数解释** 资源首层维度值，为资源的实例ID，如：4270ff17-aba3-4138-89fa-820594c39755。 **约束限制** 不涉及 **取值范围** 长度为[1,256]个字符。 **默认取值** 不涉及
	FirstDimensionValue string `json:"first_dimension_value"`

	// **参数解释** 资源名称 **约束限制** 不涉及 **取值范围** 长度[1,128]个字符 **默认取值** 不涉及
	ResourceName string `json:"resource_name"`
}

func (o ProductInstance) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ProductInstance struct{}"
	}

	return strings.Join([]string{"ProductInstance", string(data)}, " ")
}
