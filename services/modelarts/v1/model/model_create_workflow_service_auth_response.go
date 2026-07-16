package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateWorkflowServiceAuthResponse Response Object
type CreateWorkflowServiceAuthResponse struct {

	// 认证结果。
	Result         *string `json:"result,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateWorkflowServiceAuthResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateWorkflowServiceAuthResponse struct{}"
	}

	return strings.Join([]string{"CreateWorkflowServiceAuthResponse", string(data)}, " ")
}
