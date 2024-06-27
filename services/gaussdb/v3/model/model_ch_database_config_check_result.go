package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ChDatabaseConfigCheckResult 库配置校验检查结果。
type ChDatabaseConfigCheckResult struct {

	// 库同步配置参数名。
	ParamName string `json:"param_name"`

	// 库同步配置参数值。
	Value string `json:"value"`

	// 校验结果。 取值范围： - success：成功 - fail：失败
	CheckResult string `json:"check_result"`
}

func (o ChDatabaseConfigCheckResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ChDatabaseConfigCheckResult struct{}"
	}

	return strings.Join([]string{"ChDatabaseConfigCheckResult", string(data)}, " ")
}
