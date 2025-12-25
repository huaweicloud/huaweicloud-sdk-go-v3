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
}

func (o MetricDataPointDimensions) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "MetricDataPointDimensions struct{}"
	}

	return strings.Join([]string{"MetricDataPointDimensions", string(data)}, " ")
}
