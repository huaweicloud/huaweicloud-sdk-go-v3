package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowInstanceRetentionPolicyResponse Response Object
type ShowInstanceRetentionPolicyResponse struct {

	// 老化策略ID
	Id *int32 `json:"id,omitempty"`

	// 算法
	Algorithm *string `json:"algorithm,omitempty"`

	// 匹配规则
	Rules *[]RetentionRuleResponseBody `json:"rules,omitempty"`

	Trigger *TriggerConfig `json:"trigger,omitempty"`

	// 是否开启策略
	Enabled *bool `json:"enabled,omitempty"`

	// 策略名称
	Name *string `json:"name,omitempty"`

	// 命名空间ID
	NamespaceId *int32 `json:"namespace_id,omitempty"`

	// 命名空间名
	NamespaceName  *string `json:"namespace_name,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowInstanceRetentionPolicyResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowInstanceRetentionPolicyResponse struct{}"
	}

	return strings.Join([]string{"ShowInstanceRetentionPolicyResponse", string(data)}, " ")
}
