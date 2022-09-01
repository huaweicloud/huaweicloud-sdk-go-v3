package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// QoS数据。
type QosInfo struct {

	// Qos类型 - aduio：音频 - video：视频 - screen：屏幕共享 - cpu：cpu
	Type *string `json:"type,omitempty" xml:"type"`

	Send *QosSendReceiveInfo `json:"send,omitempty" xml:"send"`

	Receive *QosSendReceiveInfo `json:"receive,omitempty" xml:"receive"`

	Cpu *QosCpuInfo `json:"cpu,omitempty" xml:"cpu"`
}

func (o QosInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "QosInfo struct{}"
	}

	return strings.Join([]string{"QosInfo", string(data)}, " ")
}
