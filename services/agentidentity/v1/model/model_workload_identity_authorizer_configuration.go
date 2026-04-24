package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type WorkloadIdentityAuthorizerConfiguration struct {

	// The name of the workload identity.
	WorkloadIdentityName string `json:"workload_identity_name"`

	AuthorizerType *AuthorizerType `json:"authorizer_type"`

	AuthorizerConfiguration *AuthorizerConfiguration `json:"authorizer_configuration,omitempty"`
}

func (o WorkloadIdentityAuthorizerConfiguration) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "WorkloadIdentityAuthorizerConfiguration struct{}"
	}

	return strings.Join([]string{"WorkloadIdentityAuthorizerConfiguration", string(data)}, " ")
}
