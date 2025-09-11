package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchDeleteOneClickAlarmsResponse Response Object
type BatchDeleteOneClickAlarmsResponse struct {

	// **参数解释**： 成功删除的一键告警ID列表。
	OneClickAlarmIds *[]string `json:"one_click_alarm_ids,omitempty"`
	HttpStatusCode   int       `json:"-"`
}

func (o BatchDeleteOneClickAlarmsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDeleteOneClickAlarmsResponse struct{}"
	}

	return strings.Join([]string{"BatchDeleteOneClickAlarmsResponse", string(data)}, " ")
}
