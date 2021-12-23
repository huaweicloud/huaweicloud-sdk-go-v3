package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type CreateLogStreamRequest struct {
	// 租户想创建的日志流所在的日志组的groupid，一般为36位字符串。

	LogGroupId string `json:"log_group_id"`

	Body *CreateLogStreamParams `json:"body,omitempty"`
}

func (o CreateLogStreamRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateLogStreamRequest struct{}"
	}

	return strings.Join([]string{"CreateLogStreamRequest", string(data)}, " ")
}
