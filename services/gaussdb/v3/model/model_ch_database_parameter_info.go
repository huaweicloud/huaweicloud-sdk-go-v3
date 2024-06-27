package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ChDatabaseParameterInfo 数据库配置参数。
type ChDatabaseParameterInfo struct {

	// 参数名称。
	ParamName string `json:"param_name"`

	// 参数类型。
	DataType string `json:"data_type"`

	// 参数默认值。
	DefaultValue string `json:"default_value"`

	// 参数取值范围。
	ValueRange string `json:"value_range"`

	// 参数描述。
	Description string `json:"description"`
}

func (o ChDatabaseParameterInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ChDatabaseParameterInfo struct{}"
	}

	return strings.Join([]string{"ChDatabaseParameterInfo", string(data)}, " ")
}
