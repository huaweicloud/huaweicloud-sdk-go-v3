package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SupportVersions **参数解释**： 插件支持升级的集群版本。 **约束限制**： 不涉及 **取值范围**： 不涉及 **默认取值**： 不涉及
type SupportVersions struct {

	// **参数解释**： 支持的集群类型。 **约束限制**： CCE Autopilot集群仅支持VirtualMachine类型 **取值范围**： - VirtualMachine：CCE集群，控制节点架构为X86 - BareMetal：裸金属集群，控制节点部署在裸金属服务器上 - ARM64：鲲鹏集群，控制节点架构为鲲鹏  **默认取值**： 不涉及
	ClusterType string `json:"clusterType"`

	// **参数解释**： 支持的集群版本（正则表达式）。 **约束限制**： 不涉及 **取值范围**： 不涉及 **默认取值**： 不涉及
	ClusterVersion []string `json:"clusterVersion"`

	// **参数解释**： 作用的集群类型。 **约束限制**： 不涉及 **取值范围**： - CCE：CCE Standard集群 - Turbo：CCE Turbo集群 - Autopilot：CCE Autopilot集群  **默认取值**： 为空时默认为CCE Standard，CCE Turbo集群
	Category *[]string `json:"category,omitempty"`
}

func (o SupportVersions) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SupportVersions struct{}"
	}

	return strings.Join([]string{"SupportVersions", string(data)}, " ")
}
