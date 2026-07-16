package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// WorkflowErrorInfo error info struct
type WorkflowErrorInfo struct {

	// 错误码。
	ErrorCode *string `json:"error_code,omitempty"`

	// 错误信息。
	ErrorMessage *string `json:"error_message,omitempty"`
}

func (o WorkflowErrorInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "WorkflowErrorInfo struct{}"
	}

	return strings.Join([]string{"WorkflowErrorInfo", string(data)}, " ")
}
