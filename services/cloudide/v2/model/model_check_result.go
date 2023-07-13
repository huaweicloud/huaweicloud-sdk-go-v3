package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CheckResult struct {

	// 解析状态
	Status *string `json:"status,omitempty"`

	// 检查插件错误结果集
	Errors *[]CheckResultError `json:"errors,omitempty"`

	// 插件版本信息
	ExtensionVersionCompare *string `json:"extension_version_compare,omitempty"`
}

func (o CheckResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CheckResult struct{}"
	}

	return strings.Join([]string{"CheckResult", string(data)}, " ")
}
