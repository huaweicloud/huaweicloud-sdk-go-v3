package model

import (
	"encoding/json"

	"strings"
)

type CreateUsersDatabases struct {
	// 关联逻辑库名称。

	Name string `json:"name"`
}

func (o CreateUsersDatabases) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "CreateUsersDatabases struct{}"
	}

	return strings.Join([]string{"CreateUsersDatabases", string(data)}, " ")
}
