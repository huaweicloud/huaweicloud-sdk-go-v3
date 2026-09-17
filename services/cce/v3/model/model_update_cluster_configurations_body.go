package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateClusterConfigurationsBody **参数解释：** 更新指定集群配置参数内容请求体 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
type UpdateClusterConfigurationsBody struct {

	// **参数解释：** API版本，固定值**v3** **约束限制：** 固定值 **取值范围：** - v3  **默认取值：** v3
	ApiVersion string `json:"apiVersion"`

	// **参数解释：** API类型，固定值**Configuration** **约束限制：** 固定值 **取值范围：** - Configuration  **默认取值：** Configuration
	Kind string `json:"kind"`

	Metadata *ConfigurationMetadata `json:"metadata"`

	Spec *ClusterConfigurationsSpec `json:"spec"`
}

func (o UpdateClusterConfigurationsBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateClusterConfigurationsBody struct{}"
	}

	return strings.Join([]string{"UpdateClusterConfigurationsBody", string(data)}, " ")
}
