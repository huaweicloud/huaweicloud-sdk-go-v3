package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// MemberInvitor 联盟成员中的邀请方
type MemberInvitor struct {

	// 邀请方BCS服务实例ID
	InvitorBcsId *string `json:"invitor_bcs_id,omitempty"`

	// 邀请方BCS服务实例名称
	InvitorBcsName *string `json:"invitor_bcs_name,omitempty"`

	// 邀请方project id
	InvitorProjectId *string `json:"invitor_project_id,omitempty"`

	// 邀请方租户id
	InvitorUserId *string `json:"invitor_user_id,omitempty"`

	// 邀请方租户名
	InvitorUsername *string `json:"invitor_username,omitempty"`
}

func (o MemberInvitor) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "MemberInvitor struct{}"
	}

	return strings.Join([]string{"MemberInvitor", string(data)}, " ")
}
