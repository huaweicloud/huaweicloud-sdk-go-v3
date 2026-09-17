package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowNodePoolConfigurationsResponse Response Object
type ShowNodePoolConfigurationsResponse struct {

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

func (o ShowNodePoolConfigurationsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowNodePoolConfigurationsResponse struct{}"
	}

	return strings.Join([]string{"ShowNodePoolConfigurationsResponse", string(data)}, " ")
}
