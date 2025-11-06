package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DimensionResp **参数解释**： 指标维度。
type DimensionResp struct {

	// **参数解释**： 监控维度名称，如ECS的维度为instance_id。各服务资源的指标维度名称可查看：“[服务维度名称](ces_03_0059.xml)”。 **取值范围**： 以字母开头，只能包含字母、数字、“_”、“-”。长度[1,32]个字符。
	Name *string `json:"name,omitempty"`

	// **参数解释**： 监控维度取值，例如ECS的ID。 **取值范围**： 长度为[0,256]个字符。
	Value *string `json:"value,omitempty"`
}

func (o DimensionResp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DimensionResp struct{}"
	}

	return strings.Join([]string{"DimensionResp", string(data)}, " ")
}
