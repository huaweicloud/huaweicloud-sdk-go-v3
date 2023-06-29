package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ThresholdData 阈值查询结果结构体
type ThresholdData struct {

	// 自定义接收方向阈值，单位为毫秒(ms)
	Receiving *int32 `json:"receiving,omitempty"`

	// 默认接收方向阈值，单位为毫秒(ms)
	ReceivingDefault *int32 `json:"receivingDefault,omitempty"`

	// 自定义发送方向阈值，单位为毫秒(ms)
	Sending *int32 `json:"sending,omitempty"`

	// 默认发送方向阈值，单位为毫秒(ms)
	SendingDefault *int32 `json:"sendingDefault,omitempty"`
}

func (o ThresholdData) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ThresholdData struct{}"
	}

	return strings.Join([]string{"ThresholdData", string(data)}, " ")
}
