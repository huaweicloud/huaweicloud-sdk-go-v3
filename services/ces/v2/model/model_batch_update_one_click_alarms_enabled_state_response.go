package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchUpdateOneClickAlarmsEnabledStateResponse Response Object
type BatchUpdateOneClickAlarmsEnabledStateResponse struct {

	// **参数解释**： 成功启停的告警规则ID列表。
	AlarmIds       *[]string `json:"alarm_ids,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o BatchUpdateOneClickAlarmsEnabledStateResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchUpdateOneClickAlarmsEnabledStateResponse struct{}"
	}

	return strings.Join([]string{"BatchUpdateOneClickAlarmsEnabledStateResponse", string(data)}, " ")
}
