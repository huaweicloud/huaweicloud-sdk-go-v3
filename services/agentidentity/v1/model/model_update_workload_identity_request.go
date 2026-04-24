package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateWorkloadIdentityRequest Request Object
type UpdateWorkloadIdentityRequest struct {

	// The name of the workload identity.
	WorkloadIdentityName string `json:"workload_identity_name"`

	Body *UpdateWorkloadIdentityReqBody `json:"body,omitempty"`
}

func (o UpdateWorkloadIdentityRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateWorkloadIdentityRequest struct{}"
	}

	return strings.Join([]string{"UpdateWorkloadIdentityRequest", string(data)}, " ")
}
