package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ChangeFlinkJobStatusReportRequestBody IEF Flink job状态上报的请求body体。
type ChangeFlinkJobStatusReportRequestBody struct {

	// 作业信息列表
	Jobs []Job `json:"jobs"`

	// 消息id
	MessageId string `json:"message_id"`

	// 消息确认topic
	MsgConfirmTopic *string `json:"msg_confirm_topic,omitempty"`
}

func (o ChangeFlinkJobStatusReportRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ChangeFlinkJobStatusReportRequestBody struct{}"
	}

	return strings.Join([]string{"ChangeFlinkJobStatusReportRequestBody", string(data)}, " ")
}
