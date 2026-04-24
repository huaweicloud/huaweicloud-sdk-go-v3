package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// GetWorkloadIdentityRequest Request Object
type GetWorkloadIdentityRequest struct {

	// The name of the workload identity.
	WorkloadIdentityName string `json:"workload_identity_name"`
}

func (o GetWorkloadIdentityRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "GetWorkloadIdentityRequest struct{}"
	}

	return strings.Join([]string{"GetWorkloadIdentityRequest", string(data)}, " ")
}
