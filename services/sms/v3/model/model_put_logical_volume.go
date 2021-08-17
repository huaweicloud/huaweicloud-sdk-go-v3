package model

import (
	"encoding/json"

	"strings"
)

// 修改的逻辑卷信息
type PutLogicalVolume struct {
	// 逻辑卷id

	Id string `json:"id"`
	// 是否迁移

	NeedMigration *bool `json:"need_migration,omitempty"`
	// 调整大小

	AdjustSize *int64 `json:"adjust_size,omitempty"`
}

func (o PutLogicalVolume) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "PutLogicalVolume struct{}"
	}

	return strings.Join([]string{"PutLogicalVolume", string(data)}, " ")
}
