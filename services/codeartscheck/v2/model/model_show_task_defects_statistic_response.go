package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowTaskDefectsStatisticResponse Response Object
type ShowTaskDefectsStatisticResponse struct {
	Severity *StatisticSeverityV2 `json:"severity,omitempty"`

	Status *StatisticStatusV2 `json:"status,omitempty"`

	// **参数解释：** 检查开始时间，使用UTC时间，格式为yyyy-MM-ddTHH:mm:ssZ。例如：2020-09-25T12:05:00Z。 **取值范围：** 不涉及。
	LastCheckTime *string `json:"last_check_time,omitempty"`

	// **参数解释：** 检查结束时间，使用UTC时间，格式为yyyy-MM-ddTHH:mm:ssZ。例如：2020-09-25T12:05:00Z。 **取值范围：** 不涉及。
	LastCheckFinishTime *string `json:"last_check_finish_time,omitempty"`
	HttpStatusCode      int     `json:"-"`
}

func (o ShowTaskDefectsStatisticResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowTaskDefectsStatisticResponse struct{}"
	}

	return strings.Join([]string{"ShowTaskDefectsStatisticResponse", string(data)}, " ")
}
