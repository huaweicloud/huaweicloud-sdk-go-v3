package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 对象选择返回体
type DatabaseObjectResp struct {
	// 任务ID

	JobId *string `json:"job_id,omitempty"`
	// 选择对象任务成功标志

	Status *bool `json:"status,omitempty"`
	// 错误码

	ErrorCode *string `json:"error_code,omitempty"`
	// 错误信息

	ErrorMsg *string `json:"error_msg,omitempty"`
}

func (o DatabaseObjectResp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DatabaseObjectResp struct{}"
	}

	return strings.Join([]string{"DatabaseObjectResp", string(data)}, " ")
}
