package model

import (
	"encoding/json"

	"strings"
)

//
type TokenProjectDomainResult struct {
	// 账号名。

	Name string `json:"name"`
	// 账号ID。

	Id string `json:"id"`
}

func (o TokenProjectDomainResult) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "TokenProjectDomainResult struct{}"
	}

	return strings.Join([]string{"TokenProjectDomainResult", string(data)}, " ")
}
