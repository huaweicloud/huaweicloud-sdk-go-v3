package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type DiffConfigurationRequest struct {

	// **参数解释：** 需要进行比较的参数模板ID。 **约束限制：** 不涉及。 **取值范围：** 不涉及。 **默认取值：** 不涉及。
	SourceConfigurationId string `json:"source_configuration_id"`

	// **参数解释：** 需要进行比较的参数模板ID。 **约束限制：** 不涉及。 **取值范围：** 不涉及。 **默认取值：** 不涉及。
	TargetConfigurationId string `json:"target_configuration_id"`
}

func (o DiffConfigurationRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DiffConfigurationRequest struct{}"
	}

	return strings.Join([]string{"DiffConfigurationRequest", string(data)}, " ")
}
