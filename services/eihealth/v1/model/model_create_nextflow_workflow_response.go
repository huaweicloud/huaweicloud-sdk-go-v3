package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateNextflowWorkflowResponse Response Object
type CreateNextflowWorkflowResponse struct {

	// 流程id
	Id             *string `json:"id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateNextflowWorkflowResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateNextflowWorkflowResponse struct{}"
	}

	return strings.Join([]string{"CreateNextflowWorkflowResponse", string(data)}, " ")
}
