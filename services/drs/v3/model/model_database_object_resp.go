package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 对象选择返回体
type DatabaseObjectResp struct {

	// 任务ID
	JobId *string `json:"job_id,omitempty" xml:"job_id"`

	// 选择对象任务成功标志
	Status *bool `json:"status,omitempty" xml:"status"`

	// 错误码
	ErrorCode *string `json:"error_code,omitempty" xml:"error_code"`

	// 错误信息
	ErrorMsg *string `json:"error_msg,omitempty" xml:"error_msg"`
}

func (o DatabaseObjectResp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DatabaseObjectResp struct{}"
	}

	return strings.Join([]string{"DatabaseObjectResp", string(data)}, " ")
}
