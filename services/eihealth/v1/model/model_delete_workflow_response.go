package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
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
