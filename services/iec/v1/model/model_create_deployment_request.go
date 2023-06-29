package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateDeploymentRequest Request Object
type CreateDeploymentRequest struct {
	Body *CreateDeploymentRequestBody `json:"body,omitempty"`
}

func (o CreateDeploymentRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateDeploymentRequest struct{}"
	}

	return strings.Join([]string{"CreateDeploymentRequest", string(data)}, " ")
}
