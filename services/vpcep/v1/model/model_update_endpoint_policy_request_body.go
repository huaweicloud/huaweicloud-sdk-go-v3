package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateEndpointPolicyRequestBody 更新网关型终端节点policy请求结构体
type UpdateEndpointPolicyRequestBody struct {

	// 终端节点策略信息
	PolicyStatement *[]PolicyStatement `json:"policy_statement,omitempty"`

	// iam 5.0 策略
	PolicyDocument *interface{} `json:"policy_document,omitempty"`
}

func (o UpdateEndpointPolicyRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateEndpointPolicyRequestBody struct{}"
	}

	return strings.Join([]string{"UpdateEndpointPolicyRequestBody", string(data)}, " ")
}
