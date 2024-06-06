package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type DeleteLtsConfigsRequestBody struct {

	// 日志配置列表
	LogConfigs []InstanceDeleteLtsConfig `json:"log_configs"`
}

func (o DeleteLtsConfigsRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteLtsConfigsRequestBody struct{}"
	}

	return strings.Join([]string{"DeleteLtsConfigsRequestBody", string(data)}, " ")
}
