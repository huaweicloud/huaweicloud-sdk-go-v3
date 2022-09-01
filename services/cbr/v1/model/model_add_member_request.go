package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type AddMemberRequest struct {

	// 备份副本id
	BackupId string `json:"backup_id" xml:"backup_id"`

	Body *AddMembersReq `json:"body,omitempty" xml:"body"`
}

func (o AddMemberRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AddMemberRequest struct{}"
	}

	return strings.Join([]string{"AddMemberRequest", string(data)}, " ")
}
