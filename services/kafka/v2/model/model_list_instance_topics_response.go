package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type ListInstanceTopicsResponse struct {
	// topic总数。

	Total *int32 `json:"total,omitempty"`
	// 分页查询的大小。

	Size *int32 `json:"size,omitempty"`
	// 剩余分区数。

	RemainPartitions *int32 `json:"remain_partitions,omitempty"`
	// 分区总数。

	MaxPartitions *int32 `json:"max_partitions,omitempty"`
	// topic列表。

	Topics         *[]ListInstanceTopicsRespTopics `json:"topics,omitempty"`
	HttpStatusCode int                             `json:"-"`
}

func (o ListInstanceTopicsResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ListInstanceTopicsResponse struct{}"
	}

	return strings.Join([]string{"ListInstanceTopicsResponse", string(data)}, " ")
}
