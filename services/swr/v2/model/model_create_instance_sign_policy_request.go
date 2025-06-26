package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateInstanceSignPolicyRequest Request Object
type CreateInstanceSignPolicyRequest struct {

	// 企业仓库实例ID
	InstanceId string `json:"instance_id"`

	// 命名空间名称
	NamespaceName string `json:"namespace_name"`

	Body *CreateSignaturePolicyRequestBody `json:"body,omitempty"`
}

func (o CreateInstanceSignPolicyRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateInstanceSignPolicyRequest struct{}"
	}

	return strings.Join([]string{"CreateInstanceSignPolicyRequest", string(data)}, " ")
}
