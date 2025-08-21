package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListRepositoryWebhookLogsRequest Request Object
type ListRepositoryWebhookLogsRequest struct {

	// **参数解释：** 仓库的ID，通过[[查询用户所有仓库](https://support.huaweicloud.com/api-codeartsrepo/ListUserAllRepositories.html)](tag:hws)[[查询用户所有仓库](https://support.huaweicloud.com/intl/en-us/api-codeartsrepo/ListUserAllRepositories.html)](tag:hws_hk)[查询项目列表](tag:hcs,hcs_sm)接口查询项目列表获取。 **约束限制：** 不涉及。
	RepositoryId int32 `json:"repository_id"`

	// **参数解释：**  Webhook id。
	HookId int32 `json:"hook_id"`

	// **参数解释：** 偏移量，从0开始。
	Offset *int32 `json:"offset,omitempty"`

	// **参数解释：** 返回数量。
	Limit *int32 `json:"limit,omitempty"`

	// **参数解释：** Webhook每次发送消息时会在请求体中带上uuid字段。uuid具有唯一性，合并请求相关的Webhook事件中的uuid会包含合并请求iid。此处可以使用完整的uuid或者合并请求iid。 **约束限制：** 可选。 **取值范围：** 字符串长度不少于1，不超过100。 **默认取值：** 不涉及。
	Uuid *string `json:"uuid,omitempty"`

	// **参数解释：** 检索执行时间段的起始时间
	CreatedAfter *sdktime.SdkTime `json:"created_after,omitempty"`

	// **参数解释：** 检索执行时间段的结束时间
	CreatedBefore *sdktime.SdkTime `json:"created_before,omitempty"`
}

func (o ListRepositoryWebhookLogsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListRepositoryWebhookLogsRequest struct{}"
	}

	return strings.Join([]string{"ListRepositoryWebhookLogsRequest", string(data)}, " ")
}
