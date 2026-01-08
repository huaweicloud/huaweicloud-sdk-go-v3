package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// MetricDataPointDimensions **参数解释**： 资源维度
type MetricDataPointDimensions struct {

	// **参数解释**： 资源维度名称 **取值范围**： 最小长度1，最大长度32
	Name string `json:"name"`

	// **参数解释**： 资源维度值 **取值范围**： 最小长度1，最大长度256
	Value string `json:"value"`

	// **参数解释**： 实际维度信息。 **取值范围**： 字符串长度在 1 到 1024 之间。
	OriginValue *string `json:"origin_value,omitempty"`
}

func (o MetricDataPointDimensions) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "MetricDataPointDimensions struct{}"
	}

	return strings.Join([]string{"MetricDataPointDimensions", string(data)}, " ")
}
