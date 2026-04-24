package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// GetWorkloadIdentityAuthorizerConfigurationRequest Request Object
type GetWorkloadIdentityAuthorizerConfigurationRequest struct {

	// The name of the workload identity.
	WorkloadIdentityName string `json:"workload_identity_name"`
}

func (o GetWorkloadIdentityAuthorizerConfigurationRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "GetWorkloadIdentityAuthorizerConfigurationRequest struct{}"
	}

	return strings.Join([]string{"GetWorkloadIdentityAuthorizerConfigurationRequest", string(data)}, " ")
}
