package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// smn 消息通知结果
type SmnInfo struct {

	// 记录迁移任务执行完毕后SMN消息是否发送成功。
	NotifyResult *bool `json:"notify_result,omitempty" xml:"notify_result"`

	// 记录SMN消息发送失败原因的错误码（迁移任务成功时为空）。
	NotifyErrorMessage *string `json:"notify_error_message,omitempty" xml:"notify_error_message"`

	// SMN Topic的名称（SMN消息发送成功时为空）。
	TopicName *string `json:"topic_name,omitempty" xml:"topic_name"`
}

func (o SmnInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SmnInfo struct{}"
	}

	return strings.Join([]string{"SmnInfo", string(data)}, " ")
}
