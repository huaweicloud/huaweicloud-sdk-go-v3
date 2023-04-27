package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
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
