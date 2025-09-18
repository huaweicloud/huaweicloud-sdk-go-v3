package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListActionsPipelineRunsByRunIdsResponse Response Object
type ListActionsPipelineRunsByRunIdsResponse struct {
	ErrorCode *string `json:"error_code,omitempty"`

	ErrorMsg       *string `json:"error_msg,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ListActionsPipelineRunsByRunIdsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListActionsPipelineRunsByRunIdsResponse struct{}"
	}

	return strings.Join([]string{"ListActionsPipelineRunsByRunIdsResponse", string(data)}, " ")
}
