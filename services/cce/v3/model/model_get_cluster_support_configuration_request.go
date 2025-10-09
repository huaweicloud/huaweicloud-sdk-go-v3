package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// GetClusterSupportConfigurationRequest Request Object
type GetClusterSupportConfigurationRequest struct {

	// **参数解释**： 该参数用于过滤集群架构 **约束限制**： 不涉及 **取值范围**： - ARM64: 仅获取鲲鹏集群支持的配置项  **默认取值**： 不涉及
	ClusterType *string `json:"clusterType,omitempty"`

	// **参数解释**： 该参数用于获取指定集群版本支持的配置项 **约束限制**： 不涉及 **取值范围**： 不涉及 **默认取值**： 不涉及
	ClusterVersion *string `json:"clusterVersion,omitempty"`

	// **参数解释**： 该参数用于获取指定集群支持的配置项 **约束限制**： 不涉及 **取值范围**： 不涉及 **默认取值**： 不涉及
	ClusterID *string `json:"clusterID,omitempty"`

	// **参数解释**： 该参数用于过滤掉集群网络模型相关配置项 **约束限制**： 不涉及 **取值范围**： - eni: 过滤掉云原生网络2.0模型相关配置  **默认取值**： 不涉及
	NetworkMode *string `json:"networkMode,omitempty"`
}

func (o GetClusterSupportConfigurationRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "GetClusterSupportConfigurationRequest struct{}"
	}

	return strings.Join([]string{"GetClusterSupportConfigurationRequest", string(data)}, " ")
}
