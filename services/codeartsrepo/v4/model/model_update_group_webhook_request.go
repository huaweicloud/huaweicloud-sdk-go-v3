package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateGroupWebhookRequest Request Object
type UpdateGroupWebhookRequest struct {

	// **参数解释：** 代码组id，代码组首页，Group ID后的数字Id
	GroupId int32 `json:"group_id"`

	// **参数解释：**  Webhook id。
	HookId int32 `json:"hook_id"`

	Body *WebhookParamsDto `json:"body,omitempty"`
}

func (o UpdateGroupWebhookRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateGroupWebhookRequest struct{}"
	}

	return strings.Join([]string{"UpdateGroupWebhookRequest", string(data)}, " ")
}
