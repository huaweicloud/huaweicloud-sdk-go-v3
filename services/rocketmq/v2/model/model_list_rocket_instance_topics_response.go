package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListRocketInstanceTopicsResponse Response Object
type ListRocketInstanceTopicsResponse struct {

	// Topic总数。
	Total *int32 `json:"total,omitempty"`

	// 最大可创建Topic数量。
	Max *int32 `json:"max,omitempty"`

	// 剩余可创建Topic数量。
	Remaining *int32 `json:"remaining,omitempty"`

	// 下个分页的offset。
	NextOffset *int32 `json:"next_offset,omitempty"`

	// 上个分页的offset。
	PreviousOffset *int32 `json:"previous_offset,omitempty"`

	// Topic列表。
	Topics         *[]Topic `json:"topics,omitempty"`
	HttpStatusCode int      `json:"-"`
}

func (o ListRocketInstanceTopicsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListRocketInstanceTopicsResponse struct{}"
	}

	return strings.Join([]string{"ListRocketInstanceTopicsResponse", string(data)}, " ")
}
