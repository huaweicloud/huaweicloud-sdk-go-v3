package model

import (
	"encoding/json"

	"strings"
)

//
type PwdPasswordUserDomain struct {
	// IAM用户所属账号名。
	Name string `json:"name"`
}

func (o PwdPasswordUserDomain) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "PwdPasswordUserDomain struct{}"
	}

	return strings.Join([]string{"PwdPasswordUserDomain", string(data)}, " ")
}
