package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type CreateWorkflowAuthenticationRequest struct {
}

func (o CreateWorkflowAuthenticationRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateWorkflowAuthenticationRequest struct{}"
	}

	return strings.Join([]string{"CreateWorkflowAuthenticationRequest", string(data)}, " ")
}
