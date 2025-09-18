package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteActionsRunPipelineResponse Response Object
type DeleteActionsRunPipelineResponse struct {
	ErrorCode *string `json:"error_code,omitempty"`

	ErrorMsg       *string `json:"error_msg,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o DeleteActionsRunPipelineResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteActionsRunPipelineResponse struct{}"
	}

	return strings.Join([]string{"DeleteActionsRunPipelineResponse", string(data)}, " ")
}
