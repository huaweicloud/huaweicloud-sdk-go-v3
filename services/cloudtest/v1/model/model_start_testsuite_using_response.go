package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// StartTestsuiteUsingResponse Response Object
type StartTestsuiteUsingResponse struct {
	Error *Error `json:"error,omitempty"`

	EtTraceId *string `json:"et_trace_id,omitempty"`

	Result *TaskBasicInfoVo `json:"result,omitempty"`

	Status *string `json:"status,omitempty"`

	Warn           *Warn `json:"warn,omitempty"`
	HttpStatusCode int   `json:"-"`
}

func (o StartTestsuiteUsingResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StartTestsuiteUsingResponse struct{}"
	}

	return strings.Join([]string{"StartTestsuiteUsingResponse", string(data)}, " ")
}
