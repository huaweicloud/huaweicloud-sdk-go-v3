package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SendAimBatchMessagesResponse Response Object
type SendAimBatchMessagesResponse struct {

	// 任务ID。
	TaskId *string `json:"task_id,omitempty"`

	// 任务状态。  - Success：发送成功 - Failed：发送失败  > 此状态仅代表任务提交状态，不代表智能信息发送结果。用户手机接收智能信息结果请以收到的回执结果为准，也可通过查询智能信息发送明细API获取或登录KooMessage控制台查看。
	Status *string `json:"status,omitempty"`

	// 短信ID列表，当目标号码存在多个时，每个号码都会返回一个SmsID。  当返回异常响应时不携带此字段。
	Result         *[]SmsDetailResponse `json:"result,omitempty"`
	HttpStatusCode int                  `json:"-"`
}

func (o SendAimBatchMessagesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SendAimBatchMessagesResponse struct{}"
	}

	return strings.Join([]string{"SendAimBatchMessagesResponse", string(data)}, " ")
}
