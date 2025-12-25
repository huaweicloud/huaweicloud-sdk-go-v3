package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ReadWrite 读写记录
type ReadWrite struct {

	// 数值
	AcceptCount *int64 `json:"accept_count,omitempty"`

	// 数值
	AcceptRate *int64 `json:"accept_rate,omitempty"`

	// UUID
	ChannelId *string `json:"channel_id,omitempty"`

	// 采集通道实例的数量
	ChannelInstanceCount *int32 `json:"channel_instance_count,omitempty"`

	HeartBeat *HeartBeat `json:"heart_beat,omitempty"`

	// 最后一次接收到心跳信号的时间
	HeartBeatTime *int64 `json:"heart_beat_time,omitempty"`

	// 最后一次传输的时间
	LatestTransmissionTime *sdktime.SdkTime `json:"latest_transmission_time,omitempty"`

	// UUID
	MinionId *string `json:"minion_id,omitempty"`

	// 数值
	SendCount *int64 `json:"send_count,omitempty"`

	// 数值
	SendRate *int64 `json:"send_rate,omitempty"`
}

func (o ReadWrite) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ReadWrite struct{}"
	}

	return strings.Join([]string{"ReadWrite", string(data)}, " ")
}
