package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateEndpointPolicyRequestBody 更新网关型终端节点policy请求结构体
type UpdateEndpointPolicyRequestBody struct {

	// 只涉及开启双端固定的网关型终端节点
	PolicyStatement []PolicyStatement `json:"policy_statement"`
}

func (o UpdateEndpointPolicyRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateEndpointPolicyRequestBody struct{}"
	}

	return strings.Join([]string{"UpdateEndpointPolicyRequestBody", string(data)}, " ")
}
