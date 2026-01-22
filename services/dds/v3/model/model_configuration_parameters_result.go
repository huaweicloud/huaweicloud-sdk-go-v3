package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ConfigurationParametersResult struct {

	// **参数解释：** 参数名称。 **取值范围：** 不涉及。
	Name string `json:"name"`

	// **参数解释：** 参数值。 **取值范围：** 不涉及。
	Value string `json:"value"`

	// **参数解释：** 参数描述。 **取值范围：** 不涉及。
	Description string `json:"description"`

	// **参数解释：** 参数类型。 **取值范围：** - integer - string - boolean - float - list
	Type string `json:"type"`

	// **参数解释：** 参数值范围。 **取值范围：** - integer取值0-1。 - boolean取值true/false等。
	ValueRange string `json:"value_range"`

	// **参数解释：** 参数是否需要重启。 **取值范围：** - 取值为“true”，需要重启。 - 取值为“false”，不需要重启。
	RestartRequired bool `json:"restart_required"`

	// **参数解释：** 是否只读 **取值范围：** - 取值为“true”，只读参数。 - 取值为“false”，非只读参数。
	Readonly bool `json:"readonly"`
}

func (o ConfigurationParametersResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ConfigurationParametersResult struct{}"
	}

	return strings.Join([]string{"ConfigurationParametersResult", string(data)}, " ")
}
