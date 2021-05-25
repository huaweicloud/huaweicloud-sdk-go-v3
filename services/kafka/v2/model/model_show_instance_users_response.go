package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type ShowInstanceUsersResponse struct {
	// 用户列表。

	Users          *[]ShowInstanceUsersRespUsers `json:"users,omitempty"`
	HttpStatusCode int                           `json:"-"`
}

func (o ShowInstanceUsersResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ShowInstanceUsersResponse struct{}"
	}

	return strings.Join([]string{"ShowInstanceUsersResponse", string(data)}, " ")
}
