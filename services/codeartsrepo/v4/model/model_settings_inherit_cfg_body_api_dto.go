package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type SettingsInheritCfgBodyApiDto struct {

	// **参数解释：** 继承设置。 **约束限制：** 不涉及。 **取值范围：** 不涉及。 **默认取值：** 不涉及。
	Data *[]ProjectSettingsInheritCfgDto `json:"data,omitempty"`
}

func (o SettingsInheritCfgBodyApiDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SettingsInheritCfgBodyApiDto struct{}"
	}

	return strings.Join([]string{"SettingsInheritCfgBodyApiDto", string(data)}, " ")
}
