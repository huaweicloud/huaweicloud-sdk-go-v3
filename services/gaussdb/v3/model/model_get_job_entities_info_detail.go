package model

import (
	"encoding/json"

	"strings"
)

// 根据不同的任务，显示不同的内容。
type GetJobEntitiesInfoDetail struct {
}

func (o GetJobEntitiesInfoDetail) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "GetJobEntitiesInfoDetail struct{}"
	}

	return strings.Join([]string{"GetJobEntitiesInfoDetail", string(data)}, " ")
}
