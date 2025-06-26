package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteImmutableRuleRequest Request Object
type DeleteImmutableRuleRequest struct {

	// 企业仓库实例ID
	InstanceId string `json:"instance_id"`

	// 命名空间名称
	NamespaceName string `json:"namespace_name"`

	// 不可变Tag策略ID
	ImmutableRuleId int32 `json:"immutable_rule_id"`
}

func (o DeleteImmutableRuleRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteImmutableRuleRequest struct{}"
	}

	return strings.Join([]string{"DeleteImmutableRuleRequest", string(data)}, " ")
}
