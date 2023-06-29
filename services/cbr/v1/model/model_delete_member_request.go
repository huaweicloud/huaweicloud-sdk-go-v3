package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteMemberRequest Request Object
type DeleteMemberRequest struct {

	// 备份副本id
	BackupId string `json:"backup_id"`

	// 成员id
	MemberId string `json:"member_id"`
}

func (o DeleteMemberRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteMemberRequest struct{}"
	}

	return strings.Join([]string{"DeleteMemberRequest", string(data)}, " ")
}
