package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// IsDisableResult 返回结果
type IsDisableResult struct {

	// 禁用者
	Forbiddener *string `json:"forbiddener,omitempty"`

	// 禁用描述
	Reson *string `json:"reson,omitempty"`

	// 禁用标识
	Disabled *int32 `json:"disabled,omitempty"`

	// 任务ID
	JobId *string `json:"job_id,omitempty"`

	// 明文名称
	JobName *string `json:"job_name,omitempty"`

	// 禁用时间
	ForbiddenTime float32 `json:"forbidden_time,omitempty"`
}

func (o IsDisableResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "IsDisableResult struct{}"
	}

	return strings.Join([]string{"IsDisableResult", string(data)}, " ")
}
