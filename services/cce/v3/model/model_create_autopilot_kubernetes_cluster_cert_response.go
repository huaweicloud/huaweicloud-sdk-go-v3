package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateAutopilotKubernetesClusterCertResponse Response Object
type CreateAutopilotKubernetesClusterCertResponse struct {

	// **参数解释**： API类型 **约束限制**： 该值不可修改 **取值范围**： 不涉及 **默认取值**： Config
	Kind *string `json:"kind,omitempty"`

	// **参数解释**： API版本 **约束限制**： 该值不可修改 **取值范围**： 不涉及 **默认取值**： v1
	ApiVersion *string `json:"apiVersion,omitempty"`

	// **参数解释**： 当前未使用该字段 **约束限制**： 不涉及 **取值范围**： 不涉及 **默认取值**： 默认为空
	Preferences *interface{} `json:"preferences,omitempty"`

	// **参数解释**： 集群列表。 **约束限制**： 不涉及
	Clusters *[]Clusters `json:"clusters,omitempty"`

	// **参数解释**： 存放了指定用户的一些证书信息和ClientKey信息。 **约束限制**： 不涉及
	Users *[]Users `json:"users,omitempty"`

	// **参数解释**： 上下文列表。 **约束限制**： 不涉及
	Contexts *[]Contexts `json:"contexts,omitempty"`

	// **参数解释**： 当前上下文 **约束限制**： 不涉及 **取值范围**： - external：公网访问  - internal：私网访问  **默认取值**： - 若存在publicIp（虚拟机弹性IP）时为 external。 - 若不存在publicIp为 internal。
	CurrentContext *string `json:"current-context,omitempty"`

	PortID         *string `json:"Port-ID,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateAutopilotKubernetesClusterCertResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateAutopilotKubernetesClusterCertResponse struct{}"
	}

	return strings.Join([]string{"CreateAutopilotKubernetesClusterCertResponse", string(data)}, " ")
}
