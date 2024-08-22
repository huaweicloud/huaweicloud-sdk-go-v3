package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateEndpointPolicyRequestBody 更新终端节点policy请求结构体
type UpdateEndpointPolicyRequestBody struct {

	// Gateway类型终端节点策略信息，仅限OBS、SFS的终端节点服务的enable_policy值为true时支持该参数。
	PolicyStatement *[]PolicyStatement `json:"policy_statement,omitempty"`

	// 终端节点策略信息，仅当终端节点服务的enable_policy值为true时支持该参数，默认值为完全访问权限。（OBS、SFS的终端节点服务暂不支持该参数）
	PolicyDocument *interface{} `json:"policy_document,omitempty"`
}

func (o UpdateEndpointPolicyRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateEndpointPolicyRequestBody struct{}"
	}

	return strings.Join([]string{"UpdateEndpointPolicyRequestBody", string(data)}, " ")
}
