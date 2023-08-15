package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CheckWorkflowAuthenticationRequest Request Object
type CheckWorkflowAuthenticationRequest struct {
}

func (o CheckWorkflowAuthenticationRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CheckWorkflowAuthenticationRequest struct{}"
	}

	return strings.Join([]string{"CheckWorkflowAuthenticationRequest", string(data)}, " ")
}
