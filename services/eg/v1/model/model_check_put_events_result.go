package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CheckPutEventsResult struct {

	// 事件通道id
	ChannelId *string `json:"channel_id,omitempty"`

	// 事件源名称
	SourceName *string `json:"source_name,omitempty"`

	// 发送事件是否成功检查结果
	CheckResult *bool `json:"check_result,omitempty"`

	// 发送事件是否成功检查明细
	CheckDetail *string `json:"check_detail,omitempty"`
}

func (o CheckPutEventsResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CheckPutEventsResult struct{}"
	}

	return strings.Join([]string{"CheckPutEventsResult", string(data)}, " ")
}
