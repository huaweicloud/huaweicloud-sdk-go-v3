package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateV2ComputeRequestBody **参数解释**：创建计算资源请求。 **约束限制**：name为必填字段。
type CreateV2ComputeRequestBody struct {

	// **参数解释**：计算资源名称。 **约束限制**：不涉及。 **取值范围**：支持小写字母、数字、中划线，必须字母开头、字母或数字结尾，长度为1~63个字符。 **默认取值**：不涉及。
	Name string `json:"name"`

	// **参数解释**：资源描述信息。 **约束限制**：不涉及。 **取值范围**：不允许输入感叹号、小于号、大于号、等号、与号、双引号、单引号、艾特符、井号、百分号、分号、左括号、右括号、竖线、星号、反斜杠。长度为0~512个字符。 **默认取值**：不涉及。
	Description *string `json:"description,omitempty"`

	ChargeMode *ChargeMode `json:"charge_mode,omitempty"`

	// **参数解释**：计算资源容量配置，由CapacityItem对象组成的数组，每个对象定义了资源类型、资源数量和资源规格等信息。 **约束限制**：最少1项，最多10项。
	Capacity []CapacityItem `json:"capacity"`

	// **参数解释**：计算资源额外属性键值对映射。
	Properties map[string]string `json:"properties,omitempty"`

	// **参数解释**：标签列表，由V2ResourceTag对象组成的数组，每个对象定义了标签键和标签值。 **约束限制**：最多20个标签。
	Tags *[]V2ResourceTag `json:"tags,omitempty"`

	PrometheusConfig *PrometheusConfig `json:"prometheus_config,omitempty"`
}

func (o CreateV2ComputeRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateV2ComputeRequestBody struct{}"
	}

	return strings.Join([]string{"CreateV2ComputeRequestBody", string(data)}, " ")
}
