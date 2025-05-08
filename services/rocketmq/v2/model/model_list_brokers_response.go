package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListBrokersResponse Response Object
type ListBrokersResponse struct {

	// 总数。
	Total float32 `json:"total,omitempty"`

	// 下个分页的offset。
	NextOffset *int32 `json:"next_offset,omitempty"`

	// 上个分页的offset。
	PreviousOffset *int32 `json:"previous_offset,omitempty"`

	// 代理列表。
	Brokers        *[]ListBrokersRespBrokers `json:"brokers,omitempty"`
	HttpStatusCode int                       `json:"-"`
}

func (o ListBrokersResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListBrokersResponse struct{}"
	}

	return strings.Join([]string{"ListBrokersResponse", string(data)}, " ")
}
