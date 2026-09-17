package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateNodePoolConfigurationResponse Response Object
type UpdateNodePoolConfigurationResponse struct {

	// **参数解释：** API版本，固定值**v3** **约束限制：** 固定值 **取值范围：** - v3  **默认取值：** v3
	ApiVersion *string `json:"apiVersion,omitempty"`

	// **参数解释：** API类型，固定值**Configuration** **约束限制：** 固定值 **取值范围：** - Configuration  **默认取值：** Configuration
	Kind *string `json:"kind,omitempty"`

	Metadata *ConfigurationMetadata `json:"metadata,omitempty"`

	Spec *ClusterConfigurationsSpec `json:"spec,omitempty"`

	// **参数解释：** Configuration的状态信息 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	Status         *interface{} `json:"status,omitempty"`
	HttpStatusCode int          `json:"-"`
}

func (o UpdateNodePoolConfigurationResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateNodePoolConfigurationResponse struct{}"
	}

	return strings.Join([]string{"UpdateNodePoolConfigurationResponse", string(data)}, " ")
}
