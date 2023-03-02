package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type CreateFastExecuteScriptResponse struct {

	// 工作流执行id。
	ExecutionId    *string `json:"execution_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateFastExecuteScriptResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateFastExecuteScriptResponse struct{}"
	}

	return strings.Join([]string{"CreateFastExecuteScriptResponse", string(data)}, " ")
}
