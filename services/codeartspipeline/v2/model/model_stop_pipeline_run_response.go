package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// StopPipelineRunResponse Response Object
type StopPipelineRunResponse struct {

	// 操作是否成功
	Success        *bool `json:"success,omitempty"`
	HttpStatusCode int   `json:"-"`
}

func (o StopPipelineRunResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StopPipelineRunResponse struct{}"
	}

	return strings.Join([]string{"StopPipelineRunResponse", string(data)}, " ")
}
