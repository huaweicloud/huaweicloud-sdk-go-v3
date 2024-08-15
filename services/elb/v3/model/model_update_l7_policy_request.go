package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateL7PolicyRequest Request Object
type UpdateL7PolicyRequest struct {

	// 参数解释：转发策略ID。
	L7policyId string `json:"l7policy_id"`

	Body *UpdateL7PolicyRequestBody `json:"body,omitempty"`
}

func (o UpdateL7PolicyRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateL7PolicyRequest struct{}"
	}

	return strings.Join([]string{"UpdateL7PolicyRequest", string(data)}, " ")
}
