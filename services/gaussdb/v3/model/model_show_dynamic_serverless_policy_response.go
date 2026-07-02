package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowDynamicServerlessPolicyResponse Response Object
type ShowDynamicServerlessPolicyResponse struct {

	// **参数解释**：  当前动态Serverless算力。  **取值范围**：  available_vcpus中的可选算力，大于等于min_vcpus，并且小于等于max_vcpus。未开启动态Serverless时为null。
	CurrentVcpus *string `json:"current_vcpus,omitempty"`

	// **参数解释**：  最小动态Serverless算力。  **取值范围**：  available_vcpus中的可选算力，并且小于等于max_vcpus。未开启动态Serverless时为null。
	MinVcpus *string `json:"min_vcpus,omitempty"`

	// **参数解释**：  最大动态Serverless算力。  **取值范围**：  available_vcpus中的可选算力，并且大于等于min_vcpus。未开启动态Serverless时为null。
	MaxVcpus *string `json:"max_vcpus,omitempty"`

	// **参数解释**：  可选动态Serverless算力列表，不支持动态Serverless的实例该列表为空。  **取值范围**：  不涉及。
	AvailableVcpus *[]string `json:"available_vcpus,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o ShowDynamicServerlessPolicyResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowDynamicServerlessPolicyResponse struct{}"
	}

	return strings.Join([]string{"ShowDynamicServerlessPolicyResponse", string(data)}, " ")
}
