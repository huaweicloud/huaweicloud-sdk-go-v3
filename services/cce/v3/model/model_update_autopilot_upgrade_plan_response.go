package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateAutopilotUpgradePlanResponse Response Object
type UpdateAutopilotUpgradePlanResponse struct {

	// API类型，固定值“UpgradePlan”，该值不可修改。
	Kind *string `json:"kind,omitempty"`

	// API版本，固定值“v3”，该值不可修改。
	ApiVersion *string `json:"apiVersion,omitempty"`

	Metadata *Metadata `json:"metadata,omitempty"`

	Spec *UpgradePlanSpec `json:"spec,omitempty"`

	Status         *UpgradePlanStatus `json:"status,omitempty"`
	HttpStatusCode int                `json:"-"`
}

func (o UpdateAutopilotUpgradePlanResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateAutopilotUpgradePlanResponse struct{}"
	}

	return strings.Join([]string{"UpdateAutopilotUpgradePlanResponse", string(data)}, " ")
}
