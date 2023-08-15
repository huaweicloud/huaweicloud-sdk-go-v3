package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateWorkflowTriggerStatusResponse Response Object
type UpdateWorkflowTriggerStatusResponse struct {
	Body           *string `json:"body,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o UpdateWorkflowTriggerStatusResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateWorkflowTriggerStatusResponse struct{}"
	}

	return strings.Join([]string{"UpdateWorkflowTriggerStatusResponse", string(data)}, " ")
}
