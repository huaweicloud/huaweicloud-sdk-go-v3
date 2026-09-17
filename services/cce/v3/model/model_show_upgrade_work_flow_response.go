package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowUpgradeWorkFlowResponse Response Object
type ShowUpgradeWorkFlowResponse struct {

	// **参数解释：** API类型，固定值\"WorkFlowTask\"，该值不可修改。 **约束限制：** 固定值 **取值范围：** - WorkFlowTask  **默认取值：** WorkFlowTask
	Kind *string `json:"kind,omitempty"`

	// **参数解释：** API版本，固定值\"v3\"，该值不可修改。 **约束限制：** 固定值 **取值范围：** - v3  **默认取值：** v3
	ApiVersion *string `json:"apiVersion,omitempty"`

	Metadata *Metadata `json:"metadata,omitempty"`

	Spec *WorkFlowSpec `json:"spec,omitempty"`

	Status         *WorkFlowStatus `json:"status,omitempty"`
	HttpStatusCode int             `json:"-"`
}

func (o ShowUpgradeWorkFlowResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowUpgradeWorkFlowResponse struct{}"
	}

	return strings.Join([]string{"ShowUpgradeWorkFlowResponse", string(data)}, " ")
}
