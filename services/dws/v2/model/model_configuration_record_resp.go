package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ConfigurationRecordResp struct {

	// **参数解释**： 记录ID。 **取值范围**： 不涉及。
	Id *string `json:"id,omitempty"`

	// **参数解释**： 操作名称。 **取值范围**： 不涉及。
	Operator *string `json:"operator,omitempty"`

	// **参数解释**： 开始时间。 **取值范围**： 不涉及。
	StartTime *string `json:"start_time,omitempty"`

	// **参数解释**： 结束时间。 **取值范围**： 不涉及。
	EndTime *string `json:"end_time,omitempty"`

	// **参数解释**： 状态。 **取值范围**： - Applying：应用中； - In-Sync：已同步； - Pending-Reboot：待重启；
	Status *string `json:"status,omitempty"`

	// **参数解释**： 失败原因。 **取值范围**： 不涉及。
	FailedReason *string `json:"failed_reason,omitempty"`
}

func (o ConfigurationRecordResp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ConfigurationRecordResp struct{}"
	}

	return strings.Join([]string{"ConfigurationRecordResp", string(data)}, " ")
}
