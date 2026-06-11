package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type SetDisasterRecoverySettingsRequestBody struct {

	// **参数解释：** 容灾切换的故障节点比例列表。 **约束限制：** 不涉及。 **取值范围：** 不涉及。 **默认取值：** 不涉及。
	DisasterRecoverySettings *[]SetSwitchoverRatioInfo `json:"disaster_recovery_settings,omitempty"`
}

func (o SetDisasterRecoverySettingsRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SetDisasterRecoverySettingsRequestBody struct{}"
	}

	return strings.Join([]string{"SetDisasterRecoverySettingsRequestBody", string(data)}, " ")
}
