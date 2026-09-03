package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BindingFailedDetail 技能绑定操作失败详情。
type BindingFailedDetail struct {

	// 实例 ID。
	InstanceId *string `json:"instance_id,omitempty"`

	// 技能 ID。
	SkillId *string `json:"skill_id,omitempty"`

	// 错误码，格式 WKS.XXXXXXXX。
	ErrorCode *string `json:"error_code,omitempty"`

	// 错误信息。
	ErrorMsg *string `json:"error_msg,omitempty"`
}

func (o BindingFailedDetail) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BindingFailedDetail struct{}"
	}

	return strings.Join([]string{"BindingFailedDetail", string(data)}, " ")
}
