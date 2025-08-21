package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListGroupWebhookLogsRequest Request Object
type ListGroupWebhookLogsRequest struct {

	// **参数解释：** 代码组id，代码组首页，Group ID后的数字Id
	GroupId int32 `json:"group_id"`

	// **参数解释：**  Webhook id。
	HookId int32 `json:"hook_id"`

	// **参数解释：** 偏移量，从0开始。
	Offset *int32 `json:"offset,omitempty"`

	// **参数解释：** 返回数量。
	Limit *int32 `json:"limit,omitempty"`

	// **参数解释：** 仓库ID
	RepositoryId *int32 `json:"repository_id,omitempty"`

	// **参数解释：** Webhook每次发送消息时会在请求体中带上uuid字段。uuid具有唯一性，合并请求相关的Webhook事件中的uuid会包含合并请求iid。此处可以使用完整的uuid或者合并请求iid。 **约束限制：** 可选。 **取值范围：** 字符串长度不少于1，不超过100。 **默认取值：** 不涉及。
	Uuid *string `json:"uuid,omitempty"`

	// **参数解释：** 检索执行时间段的起始时间
	CreatedAfter *sdktime.SdkTime `json:"created_after,omitempty"`

	// **参数解释：** 检索执行时间段的结束时间
	CreatedBefore *sdktime.SdkTime `json:"created_before,omitempty"`
}

func (o ListGroupWebhookLogsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListGroupWebhookLogsRequest struct{}"
	}

	return strings.Join([]string{"ListGroupWebhookLogsRequest", string(data)}, " ")
}
