package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateOrUpdateRemediationConfigurationRequest Request Object
type CreateOrUpdateRemediationConfigurationRequest struct {

	// 规则ID
	PolicyAssignmentId string `json:"policy_assignment_id"`

	Body *RemediationConfigurationRequestBody `json:"body,omitempty"`
}

func (o CreateOrUpdateRemediationConfigurationRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateOrUpdateRemediationConfigurationRequest struct{}"
	}

	return strings.Join([]string{"CreateOrUpdateRemediationConfigurationRequest", string(data)}, " ")
}
