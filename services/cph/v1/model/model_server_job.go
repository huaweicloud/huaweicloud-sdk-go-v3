package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ServerJob 云手机服务器任务。
type ServerJob struct {

	// 云手机服务器的唯一标识ID，云手机服务器相关任务包含此字段。
	ServerId *string `json:"server_id,omitempty"`

	// 任务的唯一标识。
	JobId *string `json:"job_id,omitempty"`

	// 错误码。
	ErrorCode *string `json:"error_code,omitempty"`

	// 错误说明。
	ErrorMsg *string `json:"error_msg,omitempty"`
}

func (o ServerJob) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ServerJob struct{}"
	}

	return strings.Join([]string{"ServerJob", string(data)}, " ")
}
