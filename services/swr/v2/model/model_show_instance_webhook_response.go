package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowInstanceWebhookResponse Response Object
type ShowInstanceWebhookResponse struct {

	// 触发器ID
	Id *int32 `json:"id,omitempty"`

	// 触发器名称
	Name *string `json:"name,omitempty"`

	// 触发器描述
	Description *string `json:"description,omitempty"`

	// 触发目标
	Targets *[]Target `json:"targets,omitempty"`

	// 事件类型
	EventTypes *[]string `json:"event_types,omitempty"`

	// 是否启用，可选true或false
	Enabled *bool `json:"enabled,omitempty"`

	// 命名空间ID
	NamespaceId *int32 `json:"namespace_id,omitempty"`

	// 命名空间名
	NamespaceName *string `json:"namespace_name,omitempty"`

	// 创建者
	Creator *string `json:"creator,omitempty"`

	// 创建时间
	CreatedAt *string `json:"created_at,omitempty"`

	// 更新时间
	UpdatedAt *string `json:"updated_at,omitempty"`

	// 触发作用范围规则
	ScopeRules     *[]ScopeRule `json:"scope_rules,omitempty"`
	HttpStatusCode int          `json:"-"`
}

func (o ShowInstanceWebhookResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowInstanceWebhookResponse struct{}"
	}

	return strings.Join([]string{"ShowInstanceWebhookResponse", string(data)}, " ")
}
