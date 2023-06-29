package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SparkOptionConf spark可选配置项
type SparkOptionConf struct {

	// 参数名称
	Name *string `json:"name,omitempty"`

	// 参数值
	Value *string `json:"value,omitempty"`
}

func (o SparkOptionConf) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SparkOptionConf struct{}"
	}

	return strings.Join([]string{"SparkOptionConf", string(data)}, " ")
}
