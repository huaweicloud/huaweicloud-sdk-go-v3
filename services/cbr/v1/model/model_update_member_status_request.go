package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type UpdateMemberStatusRequest struct {

	// 成员id，成员id与项目id为同一个。
	MemberId string `json:"member_id"`

	// 备份副本id
	BackupId string `json:"backup_id"`

	Body *UpdateMember `json:"body,omitempty"`
}

func (o UpdateMemberStatusRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateMemberStatusRequest struct{}"
	}

	return strings.Join([]string{"UpdateMemberStatusRequest", string(data)}, " ")
}
