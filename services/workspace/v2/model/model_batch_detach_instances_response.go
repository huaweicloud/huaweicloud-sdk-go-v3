package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchDetachInstancesResponse Response Object
type BatchDetachInstancesResponse struct {

	// 操作失败桌面列表
	FailedOperationList *[]VmOperateResult `json:"failed_operation_list,omitempty"`

	// 任务ID
	JobId          *string `json:"job_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o BatchDetachInstancesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDetachInstancesResponse struct{}"
	}

	return strings.Join([]string{"BatchDetachInstancesResponse", string(data)}, " ")
}
