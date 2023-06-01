package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 投递详情
type DeliverDetail struct {

	// 投递时间   格式 yyyy/mm/dd xx:yy:zz
	DeliverTime *string `json:"deliverTime,omitempty"`

	// 投递状态     SUCCESS Or  FAILED
	DeliverStatus *string `json:"deliverStatus,omitempty"`

	// 投递耗时，单位ms
	DeliverConsuming *string `json:"deliverConsuming,omitempty"`

	// 投递响应码
	DeliverRspCode *string `json:"deliverRspCode,omitempty"`
}

func (o DeliverDetail) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeliverDetail struct{}"
	}

	return strings.Join([]string{"DeliverDetail", string(data)}, " ")
}
