package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UserQos 会场网络质量对象
type UserQos struct {

	// 会场ID
	ParticipantID *string `json:"participantID,omitempty"`

	// 网络质量评级
	NetRate *string `json:"netRate,omitempty"`

	// 上行总带宽(kbit/s)
	BandWidthUp *int32 `json:"bandWidthUp,omitempty"`

	// 下行总带宽(kbit/s)
	BandWidthDown *int32 `json:"bandWidthDown,omitempty"`

	// 上行丢包率（千分数）
	LostPacketRateUp *int32 `json:"lostPacketRateUp,omitempty"`

	// 下行丢包率（千分数）
	LostPacketRateDown *int32 `json:"lostPacketRateDown,omitempty"`

	// 时延(ms)
	Delay *int32 `json:"delay,omitempty"`

	VideoQos *MediaQos `json:"videoQos,omitempty"`

	AudioQos *MediaQos `json:"audioQos,omitempty"`

	AuxQos *MediaQos `json:"auxQos,omitempty"`
}

func (o UserQos) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UserQos struct{}"
	}

	return strings.Join([]string{"UserQos", string(data)}, " ")
}
