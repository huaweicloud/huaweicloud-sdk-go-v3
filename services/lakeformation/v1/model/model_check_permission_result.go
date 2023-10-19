package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CheckPermissionResult 查询权限返回结果
type CheckPermissionResult struct {

	// 对应输入策略的检查结果
	CheckResult bool `json:"check_result"`

	// 错误信息
	ErrorMessage *string `json:"error_message,omitempty"`
}

func (o CheckPermissionResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CheckPermissionResult struct{}"
	}

	return strings.Join([]string{"CheckPermissionResult", string(data)}, " ")
}
