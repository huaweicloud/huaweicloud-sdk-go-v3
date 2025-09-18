package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListActionsPipelineRunsResponse Response Object
type ListActionsPipelineRunsResponse struct {
	ErrorCode *string `json:"error_code,omitempty"`

	ErrorMsg       *string `json:"error_msg,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ListActionsPipelineRunsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListActionsPipelineRunsResponse struct{}"
	}

	return strings.Join([]string{"ListActionsPipelineRunsResponse", string(data)}, " ")
}
