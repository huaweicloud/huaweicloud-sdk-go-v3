package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// StopExecutionResponse Response Object
type StopExecutionResponse struct {
	Body           *string `json:"body,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o StopExecutionResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StopExecutionResponse struct{}"
	}

	return strings.Join([]string{"StopExecutionResponse", string(data)}, " ")
}
