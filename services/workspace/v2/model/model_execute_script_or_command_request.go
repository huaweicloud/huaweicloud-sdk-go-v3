package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ExecuteScriptOrCommandRequest Request Object
type ExecuteScriptOrCommandRequest struct {
	Body *ExecuteScriptOrCommandReq `json:"body,omitempty"`
}

func (o ExecuteScriptOrCommandRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExecuteScriptOrCommandRequest struct{}"
	}

	return strings.Join([]string{"ExecuteScriptOrCommandRequest", string(data)}, " ")
}
