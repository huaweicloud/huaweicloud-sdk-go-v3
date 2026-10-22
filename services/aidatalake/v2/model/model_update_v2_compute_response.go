package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateV2ComputeResponse Response Object
type UpdateV2ComputeResponse struct {

	// **参数解释**：计算资源ID。 **取值范围**：长度为1~36个英文字母、数字和中划线的组合。
	Id *string `json:"id,omitempty"`

	// **参数解释**：计算资源名称。 **约束限制**：不涉及。 **取值范围**：支持小写字母、数字、中划线，必须字母开头、字母或数字结尾，长度为1~63个字符。 **默认取值**：不涉及。
	Name *string `json:"name,omitempty"`

	// **参数解释**：资源描述信息。 **取值范围**：不允许输入感叹号、小于号、大于号、等号、与号、双引号、单引号、艾特符、井号、百分号、分号、左括号、右括号、竖线、星号、反斜杠。长度为0~512个字符。
	Description *string `json:"description,omitempty"`

	ChargeMode *ChargeMode `json:"charge_mode,omitempty"`

	// **参数解释**：计算资源容量配置，由ResourceCapacityItem对象组成的数组，每个对象定义了资源类型、资源数量、资源规格和资源ID等信息。
	Capacity *[]ResourceCapacityItem `json:"capacity,omitempty"`

	// **参数解释**：计算资源额外属性键值对映射。
	Properties map[string]string `json:"properties,omitempty"`

	// **参数解释**：标签列表，由Tag对象组成的数组，每个对象定义了标签键和标签值。
	Tags *[]Tag `json:"tags,omitempty"`

	PrometheusConfig *PrometheusConfig `json:"prometheus_config,omitempty"`
	HttpStatusCode   int               `json:"-"`
}

func (o UpdateV2ComputeResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateV2ComputeResponse struct{}"
	}

	return strings.Join([]string{"UpdateV2ComputeResponse", string(data)}, " ")
}
