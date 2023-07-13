package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateWorkflowResponse Response Object
type CreateWorkflowResponse struct {

	// 流程id
	Id             *string `json:"id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateWorkflowResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateWorkflowResponse struct{}"
	}

	return strings.Join([]string{"CreateWorkflowResponse", string(data)}, " ")
}
