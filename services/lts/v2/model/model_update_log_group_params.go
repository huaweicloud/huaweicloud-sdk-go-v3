package model

import (
	"encoding/json"

	"strings"
)

// 修改日志组的参数。
type UpdateLogGroupParams struct {
	// 日志存储时间 天。 取值范围为 [1, 30]

	TtlInDays int32 `json:"ttl_in_days"`
}

func (o UpdateLogGroupParams) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "UpdateLogGroupParams struct{}"
	}

	return strings.Join([]string{"UpdateLogGroupParams", string(data)}, " ")
}
