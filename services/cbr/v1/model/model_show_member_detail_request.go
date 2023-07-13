package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowMemberDetailRequest Request Object
type ShowMemberDetailRequest struct {

	// 备份副本id
	BackupId string `json:"backup_id"`

	// 成员id，为接收方的project_id
	MemberId string `json:"member_id"`
}

func (o ShowMemberDetailRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowMemberDetailRequest struct{}"
	}

	return strings.Join([]string{"ShowMemberDetailRequest", string(data)}, " ")
}
