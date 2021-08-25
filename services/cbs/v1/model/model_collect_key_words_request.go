package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type CollectKeyWordsRequest struct {
	// qabot编号，UUID格式。

	QabotId string `json:"qabot_id"`
	// 查询的起始时间，long，UTC时间，默认值为0。

	StartTime *string `json:"start_time,omitempty"`
	// 查询的结束时间，long，UTC时间，默认值为当前时间的毫秒数。

	EndTime *string `json:"end_time,omitempty"`
	// 关键词最多显示的个数，默认值为10，取值范围0-50。

	Top *int32 `json:"top,omitempty"`
}

func (o CollectKeyWordsRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "CollectKeyWordsRequest struct{}"
	}

	return strings.Join([]string{"CollectKeyWordsRequest", string(data)}, " ")
}
