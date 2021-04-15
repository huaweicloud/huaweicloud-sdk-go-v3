package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type RevokeSqlserverDbUserPrivilegeResponse struct {
	// 操作结果。

	Resp           *string `json:"resp,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o RevokeSqlserverDbUserPrivilegeResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "RevokeSqlserverDbUserPrivilegeResponse struct{}"
	}

	return strings.Join([]string{"RevokeSqlserverDbUserPrivilegeResponse", string(data)}, " ")
}
