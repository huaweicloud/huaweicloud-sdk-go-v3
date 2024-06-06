package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CreateLtsConfigs struct {

	// 日志配置列表
	LogConfigs []InstanceSaveLtsConfig `json:"log_configs"`
}

func (o CreateLtsConfigs) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateLtsConfigs struct{}"
	}

	return strings.Join([]string{"CreateLtsConfigs", string(data)}, " ")
}
