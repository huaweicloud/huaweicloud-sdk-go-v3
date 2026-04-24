package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateWorkloadIdentityResponse Response Object
type UpdateWorkloadIdentityResponse struct {
	WorkloadIdentity *WorkloadIdentity `json:"workload_identity,omitempty"`
	HttpStatusCode   int               `json:"-"`
}

func (o UpdateWorkloadIdentityResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateWorkloadIdentityResponse struct{}"
	}

	return strings.Join([]string{"UpdateWorkloadIdentityResponse", string(data)}, " ")
}
