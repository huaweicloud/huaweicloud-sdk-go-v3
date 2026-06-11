package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowDisasterRecoverySettingsResponse Response Object
type ShowDisasterRecoverySettingsResponse struct {

	// **参数解释：** 容灾切换的故障节点比例列表。 **取值范围：** 不涉及。
	DisasterRecoverySettings *[]QuerySwitchoverRatioInfo `json:"disaster_recovery_settings,omitempty"`

	// **参数解释：** 参数修改历史记录总条数。 **取值范围：** 不涉及。
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
