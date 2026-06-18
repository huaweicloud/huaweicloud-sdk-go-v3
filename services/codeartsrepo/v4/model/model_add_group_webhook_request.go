package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AddGroupWebhookRequest Request Object
type AddGroupWebhookRequest struct {

	// **参数解释：** 代码组id，代码组首页，Group ID后的数字Id **默认取值：** 不涉及。
	GroupId int32 `json:"group_id"`

	Body *WebhookParamsRequestDto `json:"body,omitempty"`
}

func (o AddGroupWebhookRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AddGroupWebhookRequest struct{}"
	}

	return strings.Join([]string{"AddGroupWebhookRequest", string(data)}, " ")
}
