package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateNextflowWorkflowResponse Response Object
type UpdateNextflowWorkflowResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o UpdateNextflowWorkflowResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateNextflowWorkflowResponse struct{}"
	}

	return strings.Join([]string{"UpdateNextflowWorkflowResponse", string(data)}, " ")
}
