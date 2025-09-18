package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// StepInfo 子步骤信息
type StepInfo struct {

	// 请求的region
	Region *string `json:"region,omitempty"`

	// 应用id
	Id *int64 `json:"id,omitempty"`

	// 偏移量
	Offset *int32 `json:"offset,omitempty"`

	// 是否启用
	Enable *bool `json:"enable,omitempty"`

	// 步骤名称
	Name *string `json:"name,omitempty"`

	// 当前偏移量
	CurrentOffset *int32 `json:"current_offset,omitempty"`

	// 任务执行时长
	ElapsedTime *int64 `json:"elapsed_time,omitempty"`

	// 常见问题
	FaqUrl *string `json:"faq_url,omitempty"`
}

func (o StepInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StepInfo struct{}"
	}

	return strings.Join([]string{"StepInfo", string(data)}, " ")
}
