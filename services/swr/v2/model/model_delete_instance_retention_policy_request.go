package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteInstanceRetentionPolicyRequest Request Object
type DeleteInstanceRetentionPolicyRequest struct {

	// 企业仓库实例ID
	InstanceId string `json:"instance_id"`

	// 命名空间名称
	NamespaceName string `json:"namespace_name"`

	// 老化策略ID
	PolicyId int32 `json:"policy_id"`
}

func (o DeleteInstanceRetentionPolicyRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteInstanceRetentionPolicyRequest struct{}"
	}

	return strings.Join([]string{"DeleteInstanceRetentionPolicyRequest", string(data)}, " ")
}
