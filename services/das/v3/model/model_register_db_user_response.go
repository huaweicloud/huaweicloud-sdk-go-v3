package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type RegisterDbUserResponse struct {
	// 数据库用户ID

	DbUserId       *string `json:"db_user_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o RegisterDbUserResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "RegisterDbUserResponse struct{}"
	}

	return strings.Join([]string{"RegisterDbUserResponse", string(data)}, " ")
}
