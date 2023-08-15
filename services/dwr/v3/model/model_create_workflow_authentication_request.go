package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateWorkflowAuthenticationRequest Request Object
type CreateWorkflowAuthenticationRequest struct {
}

func (o CreateWorkflowAuthenticationRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateWorkflowAuthenticationRequest struct{}"
	}

	return strings.Join([]string{"CreateWorkflowAuthenticationRequest", string(data)}, " ")
}
