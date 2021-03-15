package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type ShowCurUserRoleResponse struct {
	// 用户角色id
	UserRole       *int32 `json:"user_role,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ShowCurUserRoleResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ShowCurUserRoleResponse struct{}"
	}

	return strings.Join([]string{"ShowCurUserRoleResponse", string(data)}, " ")
}
