package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowGroupWebhookLogRequest Request Object
type ShowGroupWebhookLogRequest struct {

	// **参数解释：** 代码组id，代码组首页，Group ID后的数字Id
	GroupId int32 `json:"group_id"`

	// **参数解释：**  Webhook id。
	HookId int32 `json:"hook_id"`

	// **参数解释：**  Webhook 日志id。
	LogId int32 `json:"log_id"`
}

func (o ShowGroupWebhookLogRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowGroupWebhookLogRequest struct{}"
	}

	return strings.Join([]string{"ShowGroupWebhookLogRequest", string(data)}, " ")
}
