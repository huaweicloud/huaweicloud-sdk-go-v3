package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteNextflowWorkflowResponse Response Object
type DeleteNextflowWorkflowResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteNextflowWorkflowResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteNextflowWorkflowResponse struct{}"
	}

	return strings.Join([]string{"DeleteNextflowWorkflowResponse", string(data)}, " ")
}
