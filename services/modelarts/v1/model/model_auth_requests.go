package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type AuthRequests struct {

	// **参数解释**：随机UUID，用来定位使用。 **取值范围**不涉及。
	ActionId string `json:"action_id"`

	// **参数解释**：细粒度action。 **取值范围**不涉及。
	Action string `json:"action"`

	// **参数解释**：资源。 **取值范围**不涉及。
	Resource *string `json:"resource,omitempty"`

	// **参数解释**：操作对象。 **取值范围**不涉及。
	ServiceAttributes map[string]string `json:"service_attributes,omitempty"`
}

func (o AuthRequests) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AuthRequests struct{}"
	}

	return strings.Join([]string{"AuthRequests", string(data)}, " ")
}
