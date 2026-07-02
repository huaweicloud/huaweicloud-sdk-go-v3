package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateDynamicServerlessPolicyRequestBody **参数解释**：  设置动态Serverless算力策略请求体。  **约束限制**：  不涉及。
type UpdateDynamicServerlessPolicyRequestBody struct {

	// **参数解释**：   最小动态Serverless算力。  **取值范围**：  取值范围可根据[查询动态Serverless算力策略](https://support.huaweicloud.com/api-taurusdb/ShowDynamicServerlessPolicy.html)接口获取，并且小于等于max_vcpus。
	MinVcpus string `json:"min_vcpus"`

	// **参数解释**：   最大动态Serverless算力。  **取值范围**：  取值范围可根据[查询动态Serverless算力策略](https://support.huaweicloud.com/api-taurusdb/ShowDynamicServerlessPolicy.html)接口获取，并且大于等于min_vcpus。
	MaxVcpus string `json:"max_vcpus"`

	// **参数解释**：  节点算力同步，修改算力范围的同时，是否将小于最小算力的节点的当前算力同步至最小算力。  **约束限制**：  不涉及。  **取值范围**： - true: 节点算力同步。 - false: 节点算力不同步。  **默认取值**：  false。
	NeedUpdateNodesComputeAbility *bool `json:"need_update_nodes_compute_ability,omitempty"`
}

func (o UpdateDynamicServerlessPolicyRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateDynamicServerlessPolicyRequestBody struct{}"
	}

	return strings.Join([]string{"UpdateDynamicServerlessPolicyRequestBody", string(data)}, " ")
}
