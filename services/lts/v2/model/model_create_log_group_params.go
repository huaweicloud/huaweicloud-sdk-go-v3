package model

import (
	"encoding/json"

	"strings"
)

// 创建日志组参数。
type CreateLogGroupParams struct {
	// 需要创建的日志组名称。

	LogGroupName string `json:"log_group_name"`
	// 日志存储时间（天），取值范围：1-30。

	TtlInDays int32 `json:"ttl_in_days"`
}

func (o CreateLogGroupParams) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "CreateLogGroupParams struct{}"
	}

	return strings.Join([]string{"CreateLogGroupParams", string(data)}, " ")
}
