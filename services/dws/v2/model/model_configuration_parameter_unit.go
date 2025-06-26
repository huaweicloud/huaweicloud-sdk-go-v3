package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ConfigurationParameterUnit **参数解释**： 集群使用的参数配置值。 **取值范围**： 不涉及。
type ConfigurationParameterUnit struct {

	// **参数解释**： 参数类型。包括：cn、dn。 **取值范围**： cn、dn。
	Type string `json:"type"`

	// **参数解释**： 参数值。 **取值范围**： 不涉及。
	Value string `json:"value"`

	// **参数解释**： 参数默认值。 **取值范围**： 不涉及。
	DefaultValue string `json:"default_value"`
}

func (o ConfigurationParameterUnit) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ConfigurationParameterUnit struct{}"
	}

	return strings.Join([]string{"ConfigurationParameterUnit", string(data)}, " ")
}
