package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type EnvVariableBase struct {

	// 变量值支持英文字母、数字、英文格式的下划线、中划线，斜线（/）、点、冒号，1 ~ 255个字符。
	VariableValue string `json:"variable_value"`
}

func (o EnvVariableBase) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EnvVariableBase struct{}"
	}

	return strings.Join([]string{"EnvVariableBase", string(data)}, " ")
}
