package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ConsumerList struct {

	// Topic列表（当查询Topic消费“列表”时才显示此参数）。
	Topics *[]string `json:"topics,omitempty"`

	// Topic总数（当查询Topic消费“列表”时才显示此参数）。
	Total *int32 `json:"total,omitempty"`
}

func (o ConsumerList) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ConsumerList struct{}"
	}

	return strings.Join([]string{"ConsumerList", string(data)}, " ")
}
