package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateImmutableRuleRequest Request Object
type UpdateImmutableRuleRequest struct {

	// 企业仓库实例ID
	InstanceId string `json:"instance_id"`

	// 命名空间名称
	NamespaceName string `json:"namespace_name"`

	// 不可变Tag策略ID
	ImmutableRuleId int32 `json:"immutable_rule_id"`

	Body *UpdateImmutableRuleBody `json:"body,omitempty"`
}

func (o UpdateImmutableRuleRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateImmutableRuleRequest struct{}"
	}

	return strings.Join([]string{"UpdateImmutableRuleRequest", string(data)}, " ")
}
