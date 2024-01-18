package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// LogoffUserSessionReq 注销会话请求。
type LogoffUserSessionReq struct {

	// 会话信息id列表。
	SessionIds *[]string `json:"session_ids,omitempty"`

	// 客户端弹框级别，代表给会话发消息时的严重程度（比如info、warning、error级别） 0->info; 1-> warn; 2->serious。
	MessageType int32 `json:"message_type"`

	// 客户端弹框内容。
	Message *string `json:"message,omitempty"`

	// 弹框标题。
	Title *string `json:"title,omitempty"`

	// 延迟多长时间注销会话。
	DelayTime int32 `json:"delay_time"`

	// 事务id，用作客户端日志定位跟踪。
	TransactionId *string `json:"transaction_id,omitempty"`
}

func (o LogoffUserSessionReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "LogoffUserSessionReq struct{}"
	}

	return strings.Join([]string{"LogoffUserSessionReq", string(data)}, " ")
}
