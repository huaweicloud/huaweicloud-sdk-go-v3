/*
 * ProjectMan
 *
 * devcloud projectman api
 *
 */

package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type ShowCurUserRoleResponse struct {
	// 用户角色id
	RoleId *string `json:"role_id,omitempty"`
}

func (o ShowCurUserRoleResponse) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"ShowCurUserRoleResponse", string(data)}, " ")
}
