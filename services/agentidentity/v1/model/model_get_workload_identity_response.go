package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// GetWorkloadIdentityResponse Response Object
type GetWorkloadIdentityResponse struct {
	WorkloadIdentity *WorkloadIdentity `json:"workload_identity,omitempty"`
	HttpStatusCode   int               `json:"-"`
}

func (o GetWorkloadIdentityResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "GetWorkloadIdentityResponse struct{}"
	}

	return strings.Join([]string{"GetWorkloadIdentityResponse", string(data)}, " ")
}
