package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 丢包率阈值查询结果结构体
type PacketThresholdData struct {
	// 自定义接收方向阈值，单位为百分比(%)

	Receiving *int32 `json:"receiving,omitempty"`
	// 默认接收方向阈值，单位为百分比(%)

	ReceivingDefault *int32 `json:"receivingDefault,omitempty"`
	// 自定义发送方向阈值，单位为百分比(%)

	Sending *int32 `json:"sending,omitempty"`
	// 默认发送方向阈值，单位为百分比(%)

	SendingDefault *int32 `json:"sendingDefault,omitempty"`
}

func (o PacketThresholdData) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PacketThresholdData struct{}"
	}

	return strings.Join([]string{"PacketThresholdData", string(data)}, " ")
}
