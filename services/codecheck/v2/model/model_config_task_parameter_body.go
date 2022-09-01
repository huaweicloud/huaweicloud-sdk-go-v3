package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 任务配置检查参数请求信息
type ConfigTaskParameterBody struct {

	// 检查工具ID
	CheckId int32 `json:"check_id" xml:"check_id"`

	// 规则集ID
	RulesetId string `json:"ruleset_id" xml:"ruleset_id"`

	// 规则集语言
	Language string `json:"language" xml:"language"`

	// off：关闭，on：开启
	Status string `json:"status" xml:"status"`

	// 检查参数信息
	TaskCheckSettings []TaskCheckSettingsItem `json:"task_check_settings" xml:"task_check_settings"`
}

func (o ConfigTaskParameterBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ConfigTaskParameterBody struct{}"
	}

	return strings.Join([]string{"ConfigTaskParameterBody", string(data)}, " ")
}
