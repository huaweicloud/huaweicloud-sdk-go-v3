package model

import (
	"encoding/json"

	"strings"
)

type CommonQueryTaskRsp struct {
	// 任务总数
	Total *int32 `json:"total,omitempty"`
}

func (o CommonQueryTaskRsp) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "CommonQueryTaskRsp struct{}"
	}

	return strings.Join([]string{"CommonQueryTaskRsp", string(data)}, " ")
}
