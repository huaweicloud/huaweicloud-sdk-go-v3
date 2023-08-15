package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type HandleNotificationInvitor struct {

	// 邀请方BCS服务实例ID
	InvitorBcsId string `json:"invitor_bcs_id"`

	// 邀请方BCS服务实例名称
	InvitorBcsName string `json:"invitor_bcs_name"`

	// 邀请方project id
	InvitorProjectId string `json:"invitor_project_id"`

	// 邀请方租户id。控制台->邀请方帐号->我的凭证->API凭证->帐号ID
	InvitorUserId string `json:"invitor_user_id"`
}

func (o HandleNotificationInvitor) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "HandleNotificationInvitor struct{}"
	}

	return strings.Join([]string{"HandleNotificationInvitor", string(data)}, " ")
}
