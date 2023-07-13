package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateWorkflowAuthenticationResponse Response Object
type CreateWorkflowAuthenticationResponse struct {
	XRequestId *string `json:"x-request-id,omitempty"`

	ContentLength  *string `json:"Content-Length,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateWorkflowAuthenticationResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateWorkflowAuthenticationResponse struct{}"
	}

	return strings.Join([]string{"CreateWorkflowAuthenticationResponse", string(data)}, " ")
}
