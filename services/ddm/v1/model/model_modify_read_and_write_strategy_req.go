package model

import (
	"encoding/json"

	"strings"
)

// This is a auto read_weight Body Object
type ModifyReadAndWriteStrategyReq struct {
	// 主数据库实例与只读数据库实例的读权重集合。

	ReadWeight *interface{} `json:"read_weight"`
}

func (o ModifyReadAndWriteStrategyReq) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ModifyReadAndWriteStrategyReq struct{}"
	}

	return strings.Join([]string{"ModifyReadAndWriteStrategyReq", string(data)}, " ")
}
