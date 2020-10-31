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

type IssueUser struct {
	// 用户id
	Id *int32 `json:"id,omitempty"`
	// 用户名
	Name *string `json:"name,omitempty"`
	// 昵称
	NickName *string `json:"nick_name,omitempty"`
}

func (o IssueUser) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"IssueUser", string(data)}, " ")
}
