package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateFastExecuteScriptRequest Request Object
type CreateFastExecuteScriptRequest struct {
	Body *HisFastScript `json:"body,omitempty"`
}

func (o CreateFastExecuteScriptRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateFastExecuteScriptRequest struct{}"
	}

	return strings.Join([]string{"CreateFastExecuteScriptRequest", string(data)}, " ")
}
