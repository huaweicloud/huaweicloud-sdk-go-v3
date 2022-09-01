package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListInstanceTopicsResponse struct {

	// topic总数。
	Total *int32 `json:"total,omitempty" xml:"total"`

	// 分页查询的大小。
	Size *int32 `json:"size,omitempty" xml:"size"`

	// 剩余分区数。
	RemainPartitions *int32 `json:"remain_partitions,omitempty" xml:"remain_partitions"`

	// 分区总数。
	MaxPartitions *int32 `json:"max_partitions,omitempty" xml:"max_partitions"`

	// topic列表。
	Topics         *[]TopicEntity `json:"topics,omitempty" xml:"topics"`
	HttpStatusCode int            `json:"-"`
}

func (o ListInstanceTopicsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListInstanceTopicsResponse struct{}"
	}

	return strings.Join([]string{"ListInstanceTopicsResponse", string(data)}, " ")
}
