package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ServerJobInfo struct {

	// 服务器ID。
	ServerId *string `json:"server_id,omitempty"`

	// 任务ID。
	JobId *string `json:"job_id,omitempty"`

	// 失败时的错误码。
	ErrorCode *string `json:"error_code,omitempty"`

	// 失败原因。
	ErrorMessage *string `json:"error_message,omitempty"`
}

func (o ServerJobInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ServerJobInfo struct{}"
	}

	return strings.Join([]string{"ServerJobInfo", string(data)}, " ")
}
