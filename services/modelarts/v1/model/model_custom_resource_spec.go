package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CustomResourceSpec **参数解释：** 自定义规格配置。当需要使用自定义规格时需要填写此参数。 **约束限制：** 不涉及。 **参数示例：** \"custom_spec\": {     \"arch\": \"x86\",     \"cpu\": 0.51,     \"memory\": 1024.0 }
type CustomResourceSpec struct {

	// **参数解释：** GPU个数。 **约束限制：** 不涉及。 **取值范围：** 支持配置小数，输入值不能小于0（最多支持2位小数，小数点后第3位做四舍五入处理）。 **默认取值：** 不涉及。
	Gpu *float32 `json:"gpu,omitempty"`

	// **参数解释：** 内存，单位为MB。 **约束限制：** 不涉及。 **取值范围：** 仅支持整数。 **默认取值：** 不涉及。
	Memory *int32 `json:"memory,omitempty"`

	// **参数解释：** CPU核数。 **约束限制：** 不涉及。 **取值范围：** 支持配置小数，输入值不能小于0.01（最多支持2位小数，小数点后第3位做四舍五入处理）。 **默认取值：** 不涉及。
	Cpu *float32 `json:"cpu,omitempty"`

	// **参数解释：** Ascend芯片个数。 **约束限制：** 不涉及。 **取值范围：** 不涉及。 **默认取值：** 不涉及。
	Ascend *int32 `json:"ascend,omitempty"`

	// **参数解释：** 架构类型。 **约束限制：** 不涉及。 **取值范围：** 枚举值：x86 | arm64。 **默认取值：** 不涉及。
	Arch *string `json:"arch,omitempty"`
}

func (o CustomResourceSpec) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CustomResourceSpec struct{}"
	}

	return strings.Join([]string{"CustomResourceSpec", string(data)}, " ")
}
