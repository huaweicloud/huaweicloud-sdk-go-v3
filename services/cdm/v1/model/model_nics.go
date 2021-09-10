package model

import (
	"encoding/json"

	"strings"
)

type Nics struct {
	// 安全组ID

	SecurityGroupId string `json:"securityGroupId"`
	// 子网ID

	NetId string `json:"net-id"`
}

func (o Nics) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "Nics struct{}"
	}

	return strings.Join([]string{"Nics", string(data)}, " ")
}
