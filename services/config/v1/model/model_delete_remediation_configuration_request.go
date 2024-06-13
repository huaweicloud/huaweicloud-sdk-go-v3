package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteRemediationConfigurationRequest Request Object
type DeleteRemediationConfigurationRequest struct {

	// 规则ID
	PolicyAssignmentId string `json:"policy_assignment_id"`
}

func (o DeleteRemediationConfigurationRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteRemediationConfigurationRequest struct{}"
	}

	return strings.Join([]string{"DeleteRemediationConfigurationRequest", string(data)}, " ")
}
