package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type DeleteMemberRequest struct {
	// 备份副本id

	BackupId string `json:"backup_id"`
	// 成员id

	MemberId string `json:"member_id"`
}

func (o DeleteMemberRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "DeleteMemberRequest struct{}"
	}

	return strings.Join([]string{"DeleteMemberRequest", string(data)}, " ")
}
