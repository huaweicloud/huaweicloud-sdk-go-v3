package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListPartitionsResponse Response Object
type ListPartitionsResponse struct {

	// 请求结果，true为成功，false为失败
	IsSuccess *bool `json:"is_success,omitempty"`

	// 信息
	Message *string `json:"message,omitempty"`

	Partitions     *PartitionList `json:"partitions,omitempty"`
	HttpStatusCode int            `json:"-"`
}

func (o ListPartitionsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListPartitionsResponse struct{}"
	}

	return strings.Join([]string{"ListPartitionsResponse", string(data)}, " ")
}
