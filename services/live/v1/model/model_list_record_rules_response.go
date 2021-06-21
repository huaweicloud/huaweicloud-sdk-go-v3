package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type ListRecordRulesResponse struct {
	// 查询结果的总元素数量

	Total *int32 `json:"total,omitempty"`
	// 录制配置数组

	RecordConfig   *[]RecordRule `json:"record_config,omitempty"`
	HttpStatusCode int           `json:"-"`
}

func (o ListRecordRulesResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ListRecordRulesResponse struct{}"
	}

	return strings.Join([]string{"ListRecordRulesResponse", string(data)}, " ")
}
