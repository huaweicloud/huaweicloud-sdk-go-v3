package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DbConfig 数据库配置。
type DbConfig struct {

	// 参数名。
	ParamName *string `json:"param_name,omitempty"`

	// 参数值。
	Value *string `json:"value,omitempty"`
}

func (o DbConfig) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DbConfig struct{}"
	}

	return strings.Join([]string{"DbConfig", string(data)}, " ")
}
