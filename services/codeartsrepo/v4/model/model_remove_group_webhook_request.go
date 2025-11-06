package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RemoveGroupWebhookRequest Request Object
type RemoveGroupWebhookRequest struct {

	// **参数解释：** 代码组id，代码组首页，Group ID后的数字Id
	GroupId int32 `json:"group_id"`

	// **参数解释：**  Webhook id。
	HookId int32 `json:"hook_id"`
}

func (o RemoveGroupWebhookRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RemoveGroupWebhookRequest struct{}"
	}

	return strings.Join([]string{"RemoveGroupWebhookRequest", string(data)}, " ")
}
