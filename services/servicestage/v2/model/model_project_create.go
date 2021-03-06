package model

import (
	"encoding/json"

	"strings"
)

type ProjectCreate struct {
	// 项目名称。

	Name string `json:"name"`
}

func (o ProjectCreate) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ProjectCreate struct{}"
	}

	return strings.Join([]string{"ProjectCreate", string(data)}, " ")
}
