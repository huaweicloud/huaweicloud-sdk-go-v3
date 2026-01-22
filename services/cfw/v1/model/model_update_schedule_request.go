package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateScheduleRequest Request Object
type UpdateScheduleRequest struct {

	// **参数解释**： 时间表ID，可以通过调用[查询时间表列表接口]获得，通过返回值中的data.records.schedule_id获得 **约束限制**： 不涉及 **取值范围**： 32位UUID **默认取值**： 不涉及
	ScheduleId string `json:"schedule_id"`

	Body *ScheduleInfoDto `json:"body,omitempty"`
}

func (o UpdateScheduleRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateScheduleRequest struct{}"
	}

	return strings.Join([]string{"UpdateScheduleRequest", string(data)}, " ")
}
