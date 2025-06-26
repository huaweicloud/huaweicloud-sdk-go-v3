package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ConfigurationParameter **参数解释**： 集群使用的参数配置项详情。 **取值范围**： 不涉及。
type ConfigurationParameter struct {

	// **参数解释**： 参数名称。 **取值范围**： 不涉及。
	Name string `json:"name"`

	// **参数解释**： 参数值。 **取值范围**： 不涉及。
	Values []ConfigurationParameterUnit `json:"values"`

	// **参数解释**： 参数单位。 **取值范围**： 不涉及。
	Unit string `json:"unit"`

	// **参数解释**： 参数类型。 **取值范围**： 包括boolean、string、integer、float、list。
	Type string `json:"type"`

	// **参数解释**： 是否只读。 **取值范围**： 不涉及。
	Readonly bool `json:"readonly"`

	// **参数解释**： 参数值范围。 **取值范围**： 不涉及。
	ValueRange string `json:"value_range"`

	// **参数解释**： 是否需要重启。 **取值范围**： 不涉及。
	RestartRequired bool `json:"restart_required"`

	// **参数解释**： 参数描述。 **取值范围**： 不涉及。
	Description string `json:"description"`
}

func (o ConfigurationParameter) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ConfigurationParameter struct{}"
	}

	return strings.Join([]string{"ConfigurationParameter", string(data)}, " ")
}
