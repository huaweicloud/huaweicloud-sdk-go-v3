package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateWorkflowResponse Response Object
type UpdateWorkflowResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o UpdateWorkflowResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateWorkflowResponse struct{}"
	}

	return strings.Join([]string{"UpdateWorkflowResponse", string(data)}, " ")
}
