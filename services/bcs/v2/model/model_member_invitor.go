package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 联盟成员中的邀请方
type MemberInvitor struct {

	// 邀请方BCS服务实例id
	InvitorBcsId *string `json:"invitor_bcs_id,omitempty" xml:"invitor_bcs_id"`

	// 邀请方BCS服务实例名称
	InvitorBcsName *string `json:"invitor_bcs_name,omitempty" xml:"invitor_bcs_name"`

	// 邀请方project id
	InvitorProjectId *string `json:"invitor_project_id,omitempty" xml:"invitor_project_id"`

	// 邀请方租户id
	InvitorUserId *string `json:"invitor_user_id,omitempty" xml:"invitor_user_id"`

	// 邀请方租户名
	InvitorUsername *string `json:"invitor_username,omitempty" xml:"invitor_username"`
}

func (o MemberInvitor) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "MemberInvitor struct{}"
	}

	return strings.Join([]string{"MemberInvitor", string(data)}, " ")
}
