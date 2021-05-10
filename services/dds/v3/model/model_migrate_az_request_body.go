package model

import (
	"encoding/json"

	"strings"
)

type MigrateAzRequestBody struct {
	// 迁移到的目标单az或者多az。

	TargetAzs string `json:"target_azs"`
}

func (o MigrateAzRequestBody) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "MigrateAzRequestBody struct{}"
	}

	return strings.Join([]string{"MigrateAzRequestBody", string(data)}, " ")
}
