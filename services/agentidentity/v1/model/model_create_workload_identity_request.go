package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateWorkloadIdentityRequest Request Object
type CreateWorkloadIdentityRequest struct {
	Body *CreateWorkloadIdentityReqBody `json:"body,omitempty"`
}

func (o CreateWorkloadIdentityRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateWorkloadIdentityRequest struct{}"
	}

	return strings.Join([]string{"CreateWorkloadIdentityRequest", string(data)}, " ")
}
