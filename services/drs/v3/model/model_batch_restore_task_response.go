package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type BatchRestoreTaskResponse struct {
	// 批量续传返回列表

	Results *[]RetryTaskResp `json:"results,omitempty"`
	// 总数

	Count          *int32 `json:"count,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o BatchRestoreTaskResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchRestoreTaskResponse struct{}"
	}

	return strings.Join([]string{"BatchRestoreTaskResponse", string(data)}, " ")
}
