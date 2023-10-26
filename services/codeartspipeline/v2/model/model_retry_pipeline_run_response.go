package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RetryPipelineRunResponse Response Object
type RetryPipelineRunResponse struct {

	// 操作是否成功
	Success        *bool `json:"success,omitempty"`
	HttpStatusCode int   `json:"-"`
}

func (o RetryPipelineRunResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RetryPipelineRunResponse struct{}"
	}

	return strings.Join([]string{"RetryPipelineRunResponse", string(data)}, " ")
}
