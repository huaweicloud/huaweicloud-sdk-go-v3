package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchRunDesktopsResponse Response Object
type BatchRunDesktopsResponse struct {

	// 错误码，失败时返回。
	ErrorCode *string `json:"error_code,omitempty"`

	// 错误描述。
	ErrorMsg *string `json:"error_msg,omitempty"`

	// 操作失败桌面列表。
	FailedOperationList *[]VmOperateResult `json:"failed_operation_list,omitempty"`

	// 任务ID，冷迁移启动时返回。
	JobId          *string `json:"job_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o BatchRunDesktopsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchRunDesktopsResponse struct{}"
	}

	return strings.Join([]string{"BatchRunDesktopsResponse", string(data)}, " ")
}
