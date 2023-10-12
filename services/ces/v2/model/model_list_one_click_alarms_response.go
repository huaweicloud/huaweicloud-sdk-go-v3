package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListOneClickAlarmsResponse Response Object
type ListOneClickAlarmsResponse struct {

	// 一键告警列表
	OneClickAlarms *[]ListOneClickAlarmsRespOneClickAlarms `json:"one_click_alarms,omitempty"`
	HttpStatusCode int                                     `json:"-"`
}

func (o ListOneClickAlarmsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListOneClickAlarmsResponse struct{}"
	}

	return strings.Join([]string{"ListOneClickAlarmsResponse", string(data)}, " ")
}
