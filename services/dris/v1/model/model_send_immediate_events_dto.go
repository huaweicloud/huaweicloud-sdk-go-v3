package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type SendImmediateEventsDto struct {
	SendConfig *SendConfig `json:"send_config"`

	ImmediateEvent *ImmediateEventDto `json:"immediate_event"`
}

func (o SendImmediateEventsDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SendImmediateEventsDto struct{}"
	}

	return strings.Join([]string{"SendImmediateEventsDto", string(data)}, " ")
}
