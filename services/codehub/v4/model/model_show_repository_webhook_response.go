package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowRepositoryWebhookResponse Response Object
type ShowRepositoryWebhookResponse struct {

	// **参数解释：** webhook地址。 **取值范围：** 字符串长度不少于0，不超过500。
	Url *string `json:"url,omitempty"`

	// **参数解释：** 是否启用推送事件。
	PushEvents *bool `json:"push_events,omitempty"`

	// **参数解释：** 推送事件分支过滤正则规则。 **取值范围：** 字符串长度不少于0，不超过500。
	PushEventsBranchRegexFilter *string `json:"push_events_branch_regex_filter,omitempty"`

	// **参数解释：** 是否启用Tag推送事件。
	TagPushEvents *bool `json:"tag_push_events,omitempty"`

	// **参数解释：** 是否启用合并请求事件。
	MergeRequestsEvents *bool `json:"merge_requests_events,omitempty"`

	// **参数解释：** 是否启用评论事件。
	NoteEvents *bool `json:"note_events,omitempty"`

	// **参数解释：** token值，作为返回值时会使用掩码代替实际值。 **取值范围：** 字符串长度不少于0，不超过2000。
	Token *string `json:"token,omitempty"`

	// **参数解释：** token类型，默认为X-Repo-Token。 **取值范围：** 字符串长度不少于0，不超过200。
	TokenType *string `json:"token_type,omitempty"`

	// **参数解释：** 名称。 **取值范围：** 字符串长度不少于0，不超过200。
	Name *string `json:"name,omitempty"`

	// **参数解释：** 描述。 **取值范围：** 字符串长度不少于0，不超过200。
	Description *string `json:"description,omitempty"`

	// **参数解释：** Webhook id。
	Id *int32 `json:"id,omitempty"`

	// **参数解释：** 创建时间。 **参数解释：** yyyy-MM-dd'T'HH:mm:ss.SSSXXX
	CreatedAt *string `json:"created_at,omitempty"`

	// **参数解释：** 更新时间。 **参数解释：** yyyy-MM-dd'T'HH:mm:ss.SSSXXX
	UpdatedAt      *string `json:"updated_at,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowRepositoryWebhookResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowRepositoryWebhookResponse struct{}"
	}

	return strings.Join([]string{"ShowRepositoryWebhookResponse", string(data)}, " ")
}
