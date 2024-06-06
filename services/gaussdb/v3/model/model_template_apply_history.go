package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type TemplateApplyHistory struct {

	// 应用实例ID。
	TargetId *string `json:"target_id,omitempty"`

	// 应用实例名称。
	TargetName *string `json:"target_name,omitempty"`

	// 应用结果。
	ApplyResult *string `json:"apply_result,omitempty"`

	// 应用时间。
	AppliedAt float32 `json:"applied_at,omitempty"`

	// 错误码。
	ErrorCode *string `json:"error_code,omitempty"`
}

func (o TemplateApplyHistory) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TemplateApplyHistory struct{}"
	}

	return strings.Join([]string{"TemplateApplyHistory", string(data)}, " ")
}
