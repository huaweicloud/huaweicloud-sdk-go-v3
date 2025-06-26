package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateImmutableRuleRequest Request Object
type CreateImmutableRuleRequest struct {

	// 企业仓库实例ID
	InstanceId string `json:"instance_id"`

	// 命名空间名称
	NamespaceName string `json:"namespace_name"`

	Body *CreateImmutableRuleBody `json:"body,omitempty"`
}

func (o CreateImmutableRuleRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateImmutableRuleRequest struct{}"
	}

	return strings.Join([]string{"CreateImmutableRuleRequest", string(data)}, " ")
}
