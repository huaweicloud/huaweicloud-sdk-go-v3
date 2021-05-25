package model

import (
	"encoding/json"

	"strings"
)

type BatchDeleteInstanceTopicRespTopics struct {
	// Topic名称。

	Id *string `json:"id,omitempty"`
	// 是否删除成功。

	Success *bool `json:"success,omitempty"`
}

func (o BatchDeleteInstanceTopicRespTopics) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "BatchDeleteInstanceTopicRespTopics struct{}"
	}

	return strings.Join([]string{"BatchDeleteInstanceTopicRespTopics", string(data)}, " ")
}
