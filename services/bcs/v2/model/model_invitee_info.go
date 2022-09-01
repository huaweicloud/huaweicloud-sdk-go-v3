package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 被邀请方的信息
type InviteeInfo struct {

	// 被邀请用户id
	InviteeId *string `json:"invitee_id,omitempty" xml:"invitee_id"`

	// 被邀请租户名称
	InviteeName *int64 `json:"invitee_name,omitempty" xml:"invitee_name"`

	// 被邀请的BCS服务实例名称
	InviteeBcsName *string `json:"invitee_bcs_name,omitempty" xml:"invitee_bcs_name"`

	// 被邀请的BCS服务实例id
	InviteeBcsId *string `json:"invitee_bcs_id,omitempty" xml:"invitee_bcs_id"`

	// 被邀请的项目id
	InviteeProjectId *string `json:"invitee_project_id,omitempty" xml:"invitee_project_id"`
}

func (o InviteeInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "InviteeInfo struct{}"
	}

	return strings.Join([]string{"InviteeInfo", string(data)}, " ")
}
