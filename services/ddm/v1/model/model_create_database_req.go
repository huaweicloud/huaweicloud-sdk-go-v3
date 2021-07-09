package model

import (
	"encoding/json"

	"strings"
)

// This is a auto create Body Object
type CreateDatabaseReq struct {
	// 逻辑库相关信息的集合

	Databases []CreateDatabaseDetail `json:"databases"`
}

func (o CreateDatabaseReq) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "CreateDatabaseReq struct{}"
	}

	return strings.Join([]string{"CreateDatabaseReq", string(data)}, " ")
}
