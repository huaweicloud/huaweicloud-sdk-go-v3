package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateInstanceRetentionPolicyRequest Request Object
type CreateInstanceRetentionPolicyRequest struct {

	// 企业仓库实例ID
	InstanceId string `json:"instance_id"`

	// 命名空间名称
	NamespaceName string `json:"namespace_name"`

	Body *CreateRetentionPolicyRequestBody `json:"body,omitempty"`
}

func (o CreateInstanceRetentionPolicyRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateInstanceRetentionPolicyRequest struct{}"
	}

	return strings.Join([]string{"CreateInstanceRetentionPolicyRequest", string(data)}, " ")
}
