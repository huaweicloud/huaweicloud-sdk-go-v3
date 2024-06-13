package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowRemediationConfigurationRequest Request Object
type ShowRemediationConfigurationRequest struct {

	// 规则ID
	PolicyAssignmentId string `json:"policy_assignment_id"`
}

func (o ShowRemediationConfigurationRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowRemediationConfigurationRequest struct{}"
	}

	return strings.Join([]string{"ShowRemediationConfigurationRequest", string(data)}, " ")
}
