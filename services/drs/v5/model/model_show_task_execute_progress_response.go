package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowTaskExecuteProgressResponse Response Object
type ShowTaskExecuteProgressResponse struct {

	// 任务状态。
	Status *string `json:"status,omitempty"`

	// 任务执行进度。
	Progress *string `json:"progress,omitempty"`

	// 错误码。
	ErrorCode *string `json:"error_code,omitempty"`

	// 错误信息。
	ErrorMessage   *string `json:"error_message,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowTaskExecuteProgressResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowTaskExecuteProgressResponse struct{}"
	}

	return strings.Join([]string{"ShowTaskExecuteProgressResponse", string(data)}, " ")
}
