package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ConsumerDetailResp struct {

	// Topic关联代理（当查询topic消费“详情”才显示此参数）。
	Brokers *[]Brokers `json:"brokers,omitempty"`
}

func (o ConsumerDetailResp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ConsumerDetailResp struct{}"
	}

	return strings.Join([]string{"ConsumerDetailResp", string(data)}, " ")
}
