package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type UpdateL7policiesRequest struct {

	// 待更新的转发策略id
	L7policyId string `json:"l7policy_id"`

	Body *UpdateL7policiesRequestBody `json:"body,omitempty"`
}

func (o UpdateL7policiesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateL7policiesRequest struct{}"
	}

	return strings.Join([]string{"UpdateL7policiesRequest", string(data)}, " ")
}
