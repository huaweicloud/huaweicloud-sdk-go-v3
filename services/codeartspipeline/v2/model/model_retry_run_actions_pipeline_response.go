package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RetryRunActionsPipelineResponse Response Object
type RetryRunActionsPipelineResponse struct {
	ErrorCode *string `json:"error_code,omitempty"`

	ErrorMsg       *string `json:"error_msg,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o RetryRunActionsPipelineResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RetryRunActionsPipelineResponse struct{}"
	}

	return strings.Join([]string{"RetryRunActionsPipelineResponse", string(data)}, " ")
}
