package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type HandleNotificationInvitee struct {
	// 被邀请方服务实例id

	InviteeBcsId string `json:"invitee_bcs_id"`
	// 被邀请方服务实例名称，同意联盟邀请时必填

	InviteeBcsName string `json:"invitee_bcs_name"`
	// 被邀请方project id

	InviteeProjectId string `json:"invitee_project_id"`
	// 被邀请方租户id

	InviteeUserId string `json:"invitee_user_id"`
}

func (o HandleNotificationInvitee) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "HandleNotificationInvitee struct{}"
	}

	return strings.Join([]string{"HandleNotificationInvitee", string(data)}, " ")
}
