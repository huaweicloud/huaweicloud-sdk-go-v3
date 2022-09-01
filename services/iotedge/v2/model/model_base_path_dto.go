package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type BasePathDto struct {

	// 节点日志根目录
	LogBasePath *string `json:"log_base_path,omitempty" xml:"log_base_path"`

	// 节点配置根目录
	ConfigBasePath *string `json:"config_base_path,omitempty" xml:"config_base_path"`

	// 节点数据存储根目录
	DbBasePath *string `json:"db_base_path,omitempty" xml:"db_base_path"`
}

func (o BasePathDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BasePathDto struct{}"
	}

	return strings.Join([]string{"BasePathDto", string(data)}, " ")
}
