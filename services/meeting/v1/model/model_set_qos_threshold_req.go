package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 设置企业用户指定类型的阈值的请求体
type SetQosThresholdReq struct {
	Latency *SetThresholdData `json:"latency,omitempty" xml:"latency"`

	Jitter *SetThresholdData `json:"jitter,omitempty" xml:"jitter"`

	PacketLoss *SetPacketThresholdData `json:"packetLoss,omitempty" xml:"packetLoss"`

	ClientCpuMax *SetCpuThresholdData `json:"clientCpuMax,omitempty" xml:"clientCpuMax"`

	SystemCpuMax *SetCpuThresholdData `json:"systemCpuMax,omitempty" xml:"systemCpuMax"`
}

func (o SetQosThresholdReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SetQosThresholdReq struct{}"
	}

	return strings.Join([]string{"SetQosThresholdReq", string(data)}, " ")
}
