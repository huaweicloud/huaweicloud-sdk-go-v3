package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListAutopilotUpgradePlansResponse Response Object
type ListAutopilotUpgradePlansResponse struct {

	// API类型，固定值“List”，该值不可修改。
	Kind *string `json:"kind,omitempty"`

	// API版本，固定值“v3”，该值不可修改。
	ApiVersion *string `json:"apiVersion,omitempty"`

	Items          *UpgradePlan `json:"items,omitempty"`
	HttpStatusCode int          `json:"-"`
}

func (o ListAutopilotUpgradePlansResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAutopilotUpgradePlansResponse struct{}"
	}

	return strings.Join([]string{"ListAutopilotUpgradePlansResponse", string(data)}, " ")
}
