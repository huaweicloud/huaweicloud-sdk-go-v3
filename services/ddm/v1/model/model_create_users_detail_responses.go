package model

import (
	"encoding/json"

	"strings"
)

// This is a auto create response Object
type CreateUsersDetailResponses struct {
	// DDM实例帐号名称。

	Name string `json:"name"`
}

func (o CreateUsersDetailResponses) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "CreateUsersDetailResponses struct{}"
	}

	return strings.Join([]string{"CreateUsersDetailResponses", string(data)}, " ")
}
