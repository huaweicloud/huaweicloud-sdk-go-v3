package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type ShowMemberDetailRequest struct {
	// 备份副本id

	BackupId string `json:"backup_id"`
	// 成员id

	MemberId string `json:"member_id"`
}

func (o ShowMemberDetailRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ShowMemberDetailRequest struct{}"
	}

	return strings.Join([]string{"ShowMemberDetailRequest", string(data)}, " ")
}
