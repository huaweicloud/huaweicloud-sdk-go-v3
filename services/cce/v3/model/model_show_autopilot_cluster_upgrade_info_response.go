package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowAutopilotClusterUpgradeInfoResponse Response Object
type ShowAutopilotClusterUpgradeInfoResponse struct {

	// **参数解释：** 类型 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	Kind *string `json:"kind,omitempty"`

	// **参数解释：** API版本 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	ApiVersion *string `json:"apiVersion,omitempty"`

	Metadata *Metadata `json:"metadata,omitempty"`

	Spec *UpgradeInfoSpec `json:"spec,omitempty"`

	Status         *UpgradeInfoStatus `json:"status,omitempty"`
	HttpStatusCode int                `json:"-"`
}

func (o ShowAutopilotClusterUpgradeInfoResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowAutopilotClusterUpgradeInfoResponse struct{}"
	}

	return strings.Join([]string{"ShowAutopilotClusterUpgradeInfoResponse", string(data)}, " ")
}
