package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateBackupResponse Response Object
type CreateBackupResponse struct {

	// 系统提示信息，执行成功时，字段可能为空。执行失败时，用于显示错误信息。
	ErrorMessage *string `json:"errorMessage,omitempty"`

	// 系统提示信息，执行成功时，字段可能为空。执行失败时，用于显示错误码。
	ErrorCode *string `json:"errorCode,omitempty"`

	// 图备份任务ID。 >可以查询jobId查看任务执行状态、获取返回结果
	JobId          *string `json:"jobId,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateBackupResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateBackupResponse struct{}"
	}

	return strings.Join([]string{"CreateBackupResponse", string(data)}, " ")
}
