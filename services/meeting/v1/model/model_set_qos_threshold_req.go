package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 设置企业用户指定类型的阈值的请求体
type SetQosThresholdReq struct {
	Latency *SetThresholdData `json:"latency,omitempty"`

	Jitter *SetThresholdData `json:"jitter,omitempty"`

	PacketLoss *SetPacketThresholdData `json:"packetLoss,omitempty"`

	ClientCpuMax *SetCpuThresholdData `json:"clientCpuMax,omitempty"`

	SystemCpuMax *SetCpuThresholdData `json:"systemCpuMax,omitempty"`
}

func (o SetQosThresholdReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SetQosThresholdReq struct{}"
	}

	return strings.Join([]string{"SetQosThresholdReq", string(data)}, " ")
}
