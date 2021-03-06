package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type AddMemberRequest struct {
	// 备份副本id

	BackupId string `json:"backup_id"`

	Body *AddMembersReq `json:"body,omitempty"`
}

func (o AddMemberRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "AddMemberRequest struct{}"
	}

	return strings.Join([]string{"AddMemberRequest", string(data)}, " ")
}
