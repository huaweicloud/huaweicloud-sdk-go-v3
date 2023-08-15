package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type Config struct {

	// 配置名，仅支持MRS组件配置页面上所展示的配置名。
	Key string `json:"key"`

	// 配置值
	Value string `json:"value"`

	// 配置文件名，仅支持MRS组件配置页面上所展示的文件名。
	ConfigFileName string `json:"config_file_name"`
}

func (o Config) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Config struct{}"
	}

	return strings.Join([]string{"Config", string(data)}, " ")
}
