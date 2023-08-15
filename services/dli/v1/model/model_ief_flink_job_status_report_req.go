package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// IefFlinkJobStatusReportReq IEF Flink job状态上报的请求body体。
type IefFlinkJobStatusReportReq struct {

	// 作业信息列表
	Jobs []Jobs `json:"jobs"`

	// 消息id
	MessageId string `json:"message_id"`

	// 消息确认topic
	MsgConfirmTopic *string `json:"msg_confirm_topic,omitempty"`
}

func (o IefFlinkJobStatusReportReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "IefFlinkJobStatusReportReq struct{}"
	}

	return strings.Join([]string{"IefFlinkJobStatusReportReq", string(data)}, " ")
}
