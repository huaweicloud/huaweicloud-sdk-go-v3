package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// EventExInfo 交通事件扩展信息
type EventExInfo struct {

	// **参数说明**：识别出交通事件时所对应的交通参与者。
	Participants *[]EventParticipant `json:"participants,omitempty"`

	CongestionInfo *CongestionInfo `json:"congestion_info,omitempty"`
}

func (o EventExInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EventExInfo struct{}"
	}

	return strings.Join([]string{"EventExInfo", string(data)}, " ")
}
