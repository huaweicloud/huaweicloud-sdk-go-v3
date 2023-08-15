package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// QosInfo QoS数据。
type QosInfo struct {

	// Qos类型 - aduio：音频 - video：视频 - screen：屏幕共享 - cpu：cpu
	Type *string `json:"type,omitempty"`

	Send *QosSendReceiveInfo `json:"send,omitempty"`

	Receive *QosSendReceiveInfo `json:"receive,omitempty"`

	Cpu *QosCpuInfo `json:"cpu,omitempty"`
}

func (o QosInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "QosInfo struct{}"
	}

	return strings.Join([]string{"QosInfo", string(data)}, " ")
}
