package model

import (
	"encoding/json"

	"strings"
)

type AccessoryLimitVo struct {
	// 限制文件数量

	LimitCount *string `json:"limit_count,omitempty"`
	// 限制文件大小，单位是M

	LimitSize *string `json:"limit_size,omitempty"`
	// 限制文件类型

	LimitFileType *string `json:"limit_file_type,omitempty"`
}

func (o AccessoryLimitVo) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "AccessoryLimitVo struct{}"
	}

	return strings.Join([]string{"AccessoryLimitVo", string(data)}, " ")
}
