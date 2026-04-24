package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteWorkloadIdentityRequest Request Object
type DeleteWorkloadIdentityRequest struct {

	// The name of the workload identity.
	WorkloadIdentityName string `json:"workload_identity_name"`
}

func (o DeleteWorkloadIdentityRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteWorkloadIdentityRequest struct{}"
	}

	return strings.Join([]string{"DeleteWorkloadIdentityRequest", string(data)}, " ")
}
