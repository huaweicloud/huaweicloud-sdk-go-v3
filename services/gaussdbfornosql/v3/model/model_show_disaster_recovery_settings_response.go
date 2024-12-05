package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowDisasterRecoverySettingsResponse Response Object
type ShowDisasterRecoverySettingsResponse struct {

	// 容灾切换的故障节点比例列表。
	DisasterRecoverySettings *[]SwitchoverRatioInfo `json:"disaster_recovery_settings,omitempty"`

	// 总记录数。
	TotalCount     *int32 `json:"total_count,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ShowDisasterRecoverySettingsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowDisasterRecoverySettingsResponse struct{}"
	}

	return strings.Join([]string{"ShowDisasterRecoverySettingsResponse", string(data)}, " ")
}
