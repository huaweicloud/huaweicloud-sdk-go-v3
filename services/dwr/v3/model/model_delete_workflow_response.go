package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteWorkflowResponse Response Object
type DeleteWorkflowResponse struct {
	XRequestId *string `json:"x-request-id,omitempty"`

	ContentLength  *string `json:"Content-Length,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o DeleteWorkflowResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteWorkflowResponse struct{}"
	}

	return strings.Join([]string{"DeleteWorkflowResponse", string(data)}, " ")
}
