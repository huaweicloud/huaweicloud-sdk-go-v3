package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// InvokeSlowLogArchiveResponse Response Object
type InvokeSlowLogArchiveResponse struct {

	// 是否成功
	Success *bool `json:"success,omitempty"`

	// 错误码
	ErrorCode *string `json:"error_code,omitempty"`

	// 错误信息
	ErrorMessage   *string `json:"error_message,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o InvokeSlowLogArchiveResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "InvokeSlowLogArchiveResponse struct{}"
	}

	return strings.Join([]string{"InvokeSlowLogArchiveResponse", string(data)}, " ")
}
