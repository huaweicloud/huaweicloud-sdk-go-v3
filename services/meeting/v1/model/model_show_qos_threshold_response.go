package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ShowQosThresholdResponse struct {
	Latency *ThresholdData `json:"latency,omitempty"`

	Jitter *ThresholdData `json:"jitter,omitempty"`

	PacketLoss *PacketThresholdData `json:"packetLoss,omitempty"`

	ClientCpuMax *CpuThresholdData `json:"clientCpuMax,omitempty"`

	SystemCpuMax   *CpuThresholdData `json:"systemCpuMax,omitempty"`
	HttpStatusCode int               `json:"-"`
}

func (o ShowQosThresholdResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowQosThresholdResponse struct{}"
	}

	return strings.Join([]string{"ShowQosThresholdResponse", string(data)}, " ")
}
