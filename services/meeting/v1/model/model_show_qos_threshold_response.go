package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ShowQosThresholdResponse struct {
	Latency *ThresholdData `json:"latency,omitempty" xml:"latency"`

	Jitter *ThresholdData `json:"jitter,omitempty" xml:"jitter"`

	PacketLoss *PacketThresholdData `json:"packetLoss,omitempty" xml:"packetLoss"`

	ClientCpuMax *CpuThresholdData `json:"clientCpuMax,omitempty" xml:"clientCpuMax"`

	SystemCpuMax   *CpuThresholdData `json:"systemCpuMax,omitempty" xml:"systemCpuMax"`
	HttpStatusCode int               `json:"-"`
}

func (o ShowQosThresholdResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowQosThresholdResponse struct{}"
	}

	return strings.Join([]string{"ShowQosThresholdResponse", string(data)}, " ")
}
