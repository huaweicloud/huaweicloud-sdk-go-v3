package model

import (
	"encoding/json"

	"strings"
)

// 被邀请列表
type InvitedDomain struct {
	// 被邀请方租户，hcs模式下是邀请方的租户ID

	InvitedUser string `json:"invited_user"`
}

func (o InvitedDomain) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "InvitedDomain struct{}"
	}

	return strings.Join([]string{"InvitedDomain", string(data)}, " ")
}
