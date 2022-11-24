package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type CheckWorkflowAuthenticationRequest struct {
}

func (o CheckWorkflowAuthenticationRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CheckWorkflowAuthenticationRequest struct{}"
	}

	return strings.Join([]string{"CheckWorkflowAuthenticationRequest", string(data)}, " ")
}
