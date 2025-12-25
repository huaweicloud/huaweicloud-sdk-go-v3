package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// MetricDimensionResp **参数解释**： 资源维度。
type MetricDimensionResp struct {

	// **参数解释**： 资源维度名称。 **取值范围**： 以字母开头，长度为[1,32]个字符。
	Name *string `json:"name,omitempty"`

	// **参数解释**： 资源维度值。 **取值范围**： 长度为[0,256]个字符。
	Value *string `json:"value,omitempty"`
}

func (o MetricDimensionResp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "MetricDimensionResp struct{}"
	}

	return strings.Join([]string{"MetricDimensionResp", string(data)}, " ")
}
