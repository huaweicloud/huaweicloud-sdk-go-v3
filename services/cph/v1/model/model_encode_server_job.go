package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// EncodeServerJob 编码容器任务。
type EncodeServerJob struct {

	// 编码服务的唯一标识ID，编码服务相关任务包含此字段。
	EncodeServerId *string `json:"encode_server_id,omitempty"`

	// 任务的唯一标识。
	JobId *string `json:"job_id,omitempty"`

	// 错误码。
	ErrorCode *string `json:"error_code,omitempty"`

	// 错误说明。
	ErrorMsg *string `json:"error_msg,omitempty"`
}

func (o EncodeServerJob) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EncodeServerJob struct{}"
	}

	return strings.Join([]string{"EncodeServerJob", string(data)}, " ")
}
