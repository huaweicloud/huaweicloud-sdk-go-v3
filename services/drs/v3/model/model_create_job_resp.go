package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 创建任务响应体
type CreateJobResp struct {
	// 任务ID

	Id string `json:"id"`
	// 任务名称

	Name *string `json:"name,omitempty"`
	// 任务状态

	Status *string `json:"status,omitempty"`
	// 创建时间，时间戳

	CreateTime *string `json:"create_time,omitempty"`
	// 错误码

	ErrorCode *string `json:"error_code,omitempty"`
	// 错误信息

	ErrorMsg *string `json:"error_msg,omitempty"`
	// 子任务ID，有子任务时返回该字段。

	ChildIds *[]string `json:"child_ids,omitempty"`
}

func (o CreateJobResp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateJobResp struct{}"
	}

	return strings.Join([]string{"CreateJobResp", string(data)}, " ")
}
