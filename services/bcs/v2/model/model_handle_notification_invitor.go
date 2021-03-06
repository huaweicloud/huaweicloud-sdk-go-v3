package model

import (
	"encoding/json"

	"strings"
)

type HandleNotificationInvitor struct {
	// 邀请方实例id

	InvitorBcsId string `json:"invitor_bcs_id"`
	// 邀请方实例名称，同意联盟邀请时比填

	InvitorBcsName string `json:"invitor_bcs_name"`
	// 邀请方project id

	InvitorProjectId string `json:"invitor_project_id"`
	// 邀请方租户,hcs模式下是邀请方租户id

	InvitorUserId string `json:"invitor_user_id"`
}

func (o HandleNotificationInvitor) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "HandleNotificationInvitor struct{}"
	}

	return strings.Join([]string{"HandleNotificationInvitor", string(data)}, " ")
}
