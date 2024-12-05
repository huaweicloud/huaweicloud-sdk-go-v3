package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type SetDisasterRecoverySettingsRequestBody struct {

	// 容灾切换的故障节点比例列表。
	DisasterRecoverySettings *[]SwitchoverRatioInfo `json:"disaster_recovery_settings,omitempty"`
}

func (o SetDisasterRecoverySettingsRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SetDisasterRecoverySettingsRequestBody struct{}"
	}

	return strings.Join([]string{"SetDisasterRecoverySettingsRequestBody", string(data)}, " ")
}
