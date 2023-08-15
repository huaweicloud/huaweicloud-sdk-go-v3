package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CheckPutEventsReqSources struct {

	// 事件通道ID
	ChannelId *string `json:"channel_id,omitempty"`

	// 事件源名称
	SourceName *string `json:"source_name,omitempty"`
}

func (o CheckPutEventsReqSources) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CheckPutEventsReqSources struct{}"
	}

	return strings.Join([]string{"CheckPutEventsReqSources", string(data)}, " ")
}
