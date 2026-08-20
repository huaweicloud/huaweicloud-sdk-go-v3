package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpgradeInstanceResponse Response Object
type UpgradeInstanceResponse struct {

	// **参数解释**： 实例升级任务ID。 **取值范围**： 不涉及。
	JobId *string `json:"job_id,omitempty"`

	// **参数解释**： 实例升级调度任务ID。 **取值范围**： 不涉及。
	ScheduleId     *string `json:"schedule_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o UpgradeInstanceResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpgradeInstanceResponse struct{}"
	}

	return strings.Join([]string{"UpgradeInstanceResponse", string(data)}, " ")
}
