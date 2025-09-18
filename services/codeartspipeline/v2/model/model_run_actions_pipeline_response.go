package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RunActionsPipelineResponse Response Object
type RunActionsPipelineResponse struct {
	ErrorCode *string `json:"error_code,omitempty"`

	ErrorMsg       *string `json:"error_msg,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o RunActionsPipelineResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RunActionsPipelineResponse struct{}"
	}

	return strings.Join([]string{"RunActionsPipelineResponse", string(data)}, " ")
}
