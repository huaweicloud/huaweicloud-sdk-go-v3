package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ConfigurationParameterValue **参数解释**： 集群参数配置详情。 **取值范围**： 不涉及。
type ConfigurationParameterValue struct {

	// **参数解释**： 参数类型。 **取值范围**： cn、dn。
	Type string `json:"type"`

	// **参数解释**： 参数名称。 **取值范围**： 不涉及。
	Name string `json:"name"`

	// **参数解释**： 参数值。 **取值范围**： 不涉及。
	Value string `json:"value"`
}

func (o ConfigurationParameterValue) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ConfigurationParameterValue struct{}"
	}

	return strings.Join([]string{"ConfigurationParameterValue", string(data)}, " ")
}
