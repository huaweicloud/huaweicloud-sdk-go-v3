package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type MysqlEngineVersionInfo struct {

	// 数据库版本ID，该字段不会有重复
	Id string `json:"id"`

	// 数据库版本号，只返回两位数的大版本号
	Name string `json:"name"`

	// 兼容的开源数据库版本号，返回三位开源版本号。
	Version *string `json:"version,omitempty"`

	// 数据库版本号，返回完整的四位版本号。
	KernelVersion *string `json:"kernel_version,omitempty"`
}

func (o MysqlEngineVersionInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "MysqlEngineVersionInfo struct{}"
	}

	return strings.Join([]string{"MysqlEngineVersionInfo", string(data)}, " ")
}
