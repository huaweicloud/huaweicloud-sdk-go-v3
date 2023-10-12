package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateOneClickAlarmResponse Response Object
type CreateOneClickAlarmResponse struct {

	// 一键告警ID
	OneClickAlarmId *string `json:"one_click_alarm_id,omitempty"`
	HttpStatusCode  int     `json:"-"`
}

func (o CreateOneClickAlarmResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateOneClickAlarmResponse struct{}"
	}

	return strings.Join([]string{"CreateOneClickAlarmResponse", string(data)}, " ")
}
