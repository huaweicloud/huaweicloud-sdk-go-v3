package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ActionBaseResp 操作任务响应体。
type ActionBaseResp struct {

	// 错误码。
	ErrorCode string `json:"error_code"`

	// 错误描述。
	ErrorMsg string `json:"error_msg"`

	// 任务ID。
	Id string `json:"id"`

	// 任务名称。
	Name string `json:"name"`

	// 操作结果。
	Status string `json:"status"`
}

func (o ActionBaseResp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ActionBaseResp struct{}"
	}

	return strings.Join([]string{"ActionBaseResp", string(data)}, " ")
}
