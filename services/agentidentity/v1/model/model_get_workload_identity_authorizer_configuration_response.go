package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// GetWorkloadIdentityAuthorizerConfigurationResponse Response Object
type GetWorkloadIdentityAuthorizerConfigurationResponse struct {
	WorkloadIdentityAuthorizerConfiguration *WorkloadIdentityAuthorizerConfiguration `json:"workload_identity_authorizer_configuration,omitempty"`
	HttpStatusCode                          int                                      `json:"-"`
}

func (o GetWorkloadIdentityAuthorizerConfigurationResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "GetWorkloadIdentityAuthorizerConfigurationResponse struct{}"
	}

	return strings.Join([]string{"GetWorkloadIdentityAuthorizerConfigurationResponse", string(data)}, " ")
}
