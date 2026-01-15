package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AsyncUpdateJobResp 更新异步任务响应体。
type AsyncUpdateJobResp struct {

	// 任务ID。
	Id string `json:"id"`

	// 任务名称。
	Name string `json:"name"`

	// 操作结果。 取值：success，failed
	Status string `json:"status"`

	// 错误码。
	ErrorCode *string `json:"error_code,omitempty"`

	// 错误描述。
	ErrorMsg *string `json:"error_msg,omitempty"`
}

func (o AsyncUpdateJobResp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AsyncUpdateJobResp struct{}"
	}

	return strings.Join([]string{"AsyncUpdateJobResp", string(data)}, " ")
}
