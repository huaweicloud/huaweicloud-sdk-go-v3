package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ExecuteTestConnectionNewRequest Request Object
type ExecuteTestConnectionNewRequest struct {
	Body *ExecuteTestConnectionNewRequestBody `json:"body,omitempty"`
}

func (o ExecuteTestConnectionNewRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExecuteTestConnectionNewRequest struct{}"
	}

	return strings.Join([]string{"ExecuteTestConnectionNewRequest", string(data)}, " ")
}
