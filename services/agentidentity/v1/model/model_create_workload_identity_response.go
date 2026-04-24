package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateWorkloadIdentityResponse Response Object
type CreateWorkloadIdentityResponse struct {
	WorkloadIdentity *WorkloadIdentity `json:"workload_identity,omitempty"`
	HttpStatusCode   int               `json:"-"`
}

func (o CreateWorkloadIdentityResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateWorkloadIdentityResponse struct{}"
	}

	return strings.Join([]string{"CreateWorkloadIdentityResponse", string(data)}, " ")
}
