package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// InnerRayJobConfig **参数解释**：Ray作业配置。 **约束限制**：不涉及。 **取值范围**：可选参数有：   - raycluster_endpoint_name：它必须是一个有效的Ray端点，可设置多个。 **默认取值**：不涉及。
type InnerRayJobConfig struct {

	// **参数解释**：作业将运行的Ray集群端点列表，可设置多个。 **约束限制**：不涉及。 **取值范围**：每个元素最小长度1，最大长度63。 **默认取值**：不涉及。
	RayclusterEndpointName *[]string `json:"raycluster_endpoint_name,omitempty"`

	// **参数解释**：集群中所需的资源规格。 **约束限制**：最小值为1，最大值为4。
	TaskResources *[]string `json:"task_resources,omitempty"`
}

func (o InnerRayJobConfig) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "InnerRayJobConfig struct{}"
	}

	return strings.Join([]string{"InnerRayJobConfig", string(data)}, " ")
}
