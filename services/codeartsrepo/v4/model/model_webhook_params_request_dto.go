package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// WebhookParamsRequestDto **参数解释：** Webhook设置参数。
type WebhookParamsRequestDto struct {

	// **参数解释：** webhook地址。 **取值范围：** 字符串长度不少于0，不超过500。
	Url string `json:"url"`

	// **参数解释：** 是否启用推送事件。 **约束限制：** push_events，tag_push_events，merge_requests_events，note_events至少传一个为true
	PushEvents *bool `json:"push_events,omitempty"`

	// **参数解释：** 推送事件分支过滤正则规则。 **取值范围：** 字符串长度不少于0，不超过500。    **约束限制：** push_events_branch_regex_filter与push_events联动，当push_events为true时push_events_branch_regex_filter需传值正则
	PushEventsBranchRegexFilter *string `json:"push_events_branch_regex_filter,omitempty"`

	// **参数解释：** 评论事件文本过滤规则。 **取值范围：** 字符串长度不少于0，不超过50，不能超过10个。  **约束限制：** note_plain_text_filter与note_events联动，当note_events为true时note_plain_text_filter需传值正则
	NotePlainTextFilter *[]string `json:"note_plain_text_filter,omitempty"`

	// **参数解释：** 是否启用Tag推送事件。 **约束限制：** push_events，tag_push_events，merge_requests_events，note_events至少传一个为true
	TagPushEvents *bool `json:"tag_push_events,omitempty"`

	// **参数解释：** 是否启用合并请求事件。 **约束限制：** push_events，tag_push_events，merge_requests_events，note_events至少传一个为true
	MergeRequestsEvents *bool `json:"merge_requests_events,omitempty"`

	// **参数解释：** 是否启用评论事件。 **约束限制：** push_events，tag_push_events，merge_requests_events，note_events至少传一个为true
	NoteEvents *bool `json:"note_events,omitempty"`

	// **参数解释：** token值，作为返回值时会使用掩码代替实际值。 **取值范围：** 字符串长度不少于0，不超过2000。
	Token *string `json:"token,omitempty"`

	// **参数解释：** token类型，默认为X-Repo-Token。 **取值范围：** 字符串长度不少于0，不超过200。
	TokenType *string `json:"token_type,omitempty"`

	// **参数解释：** 名称。 **取值范围：** 字符串长度不少于0，不超过200。
	Name string `json:"name"`

	// **参数解释：** 描述。 **取值范围：** 字符串长度不少于0，不超过200。
	Description *string `json:"description,omitempty"`
}

func (o WebhookParamsRequestDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "WebhookParamsRequestDto struct{}"
	}

	return strings.Join([]string{"WebhookParamsRequestDto", string(data)}, " ")
}
