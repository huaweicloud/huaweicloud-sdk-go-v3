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

type BatchAddMembersV4RequestBody struct {
	// 添加的用户信息
	Users []BatchAddMemberRequestV4 `json:"users"`
}

func (o BatchAddMembersV4RequestBody) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"BatchAddMembersV4RequestBody", string(data)}, " ")
}
