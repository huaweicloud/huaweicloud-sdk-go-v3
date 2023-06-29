package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteJobResp 删除任务响应体。
type DeleteJobResp struct {

	// 错误码。
	ErrorCode string `json:"error_code"`

	// 错误描述。
	ErrorMsg string `json:"error_msg"`

	// 任务ID。
	Id string `json:"id"`

	// 任务名称。
	Name string `json:"name"`

	// 删除结果。
	Status string `json:"status"`
}

func (o DeleteJobResp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteJobResp struct{}"
	}

	return strings.Join([]string{"DeleteJobResp", string(data)}, " ")
}
