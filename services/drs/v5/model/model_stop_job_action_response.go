package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// StopJobActionResponse Response Object
type StopJobActionResponse struct {

	// 错误码。
	ErrorCode string `json:"error_code"`

	// 错误描述。
	ErrorMsg string `json:"error_msg"`

	// 任务ID。
	Id string `json:"id"`

	// 任务名称。
	Name string `json:"name"`

	// 操作结果。
	Status         string `json:"status"`
	HttpStatusCode int    `json:"-"`
}

func (o StopJobActionResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StopJobActionResponse struct{}"
	}

	return strings.Join([]string{"StopJobActionResponse", string(data)}, " ")
}
