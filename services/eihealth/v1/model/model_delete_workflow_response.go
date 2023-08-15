package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteWorkflowResponse Response Object
type DeleteWorkflowResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteWorkflowResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteWorkflowResponse struct{}"
	}

	return strings.Join([]string{"DeleteWorkflowResponse", string(data)}, " ")
}
