package model

import (
	"encoding/json"

	"strings"
)

//
type TokenProjectResult struct {
	// 项目名。

	Name string `json:"name"`
	// 项目ID。

	Id string `json:"id"`

	Domain *TokenProjectDomainResult `json:"domain"`
}

func (o TokenProjectResult) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "TokenProjectResult struct{}"
	}

	return strings.Join([]string{"TokenProjectResult", string(data)}, " ")
}
