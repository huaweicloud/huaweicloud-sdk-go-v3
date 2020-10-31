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

type BatchDeleteMembersV4RequestBody struct {
	// 用户id
	UserIds []string `json:"user_ids"`
}

func (o BatchDeleteMembersV4RequestBody) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"BatchDeleteMembersV4RequestBody", string(data)}, " ")
}
