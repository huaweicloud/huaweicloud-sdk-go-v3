package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type ListUserLoginProtectsResponse struct {
	// 登录状态保护信息列表。

	LoginProtects  *[]LoginProtectResult `json:"login_protects,omitempty"`
	HttpStatusCode int                   `json:"-"`
}

func (o ListUserLoginProtectsResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ListUserLoginProtectsResponse struct{}"
	}

	return strings.Join([]string{"ListUserLoginProtectsResponse", string(data)}, " ")
}
