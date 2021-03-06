package model

import (
	"encoding/json"

	"strings"
)

//
type OsfederationGroups struct {
	// 用户组ID。

	Id string `json:"id"`
	// 用户组名称。

	Name string `json:"name"`
}

func (o OsfederationGroups) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "OsfederationGroups struct{}"
	}

	return strings.Join([]string{"OsfederationGroups", string(data)}, " ")
}
