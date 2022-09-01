package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 查询灾备初始化进度
type QueryStructProcessResp struct {

	// 任务ID
	JobId string `json:"job_id" xml:"job_id"`

	// 错误码
	ErrorCode *string `json:"error_code,omitempty" xml:"error_code"`

	// 错误信息
	ErrorMessage *string `json:"error_message,omitempty" xml:"error_message"`

	StructProcess *StructProcessResp `json:"struct_process,omitempty" xml:"struct_process"`
}

func (o QueryStructProcessResp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "QueryStructProcessResp struct{}"
	}

	return strings.Join([]string{"QueryStructProcessResp", string(data)}, " ")
}
