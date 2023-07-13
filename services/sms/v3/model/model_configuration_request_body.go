package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ConfigurationRequestBody struct {

	// 配置项列表
	Configurations []ConfigBody `json:"configurations"`
}

func (o ConfigurationRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ConfigurationRequestBody struct{}"
	}

	return strings.Join([]string{"ConfigurationRequestBody", string(data)}, " ")
}
