package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AccessConfigDeatilCreate 日志接入详细信息
type AccessConfigDeatilCreate struct {

	// 采集路径。 1. 路径必须以/或者字母:\\\\开头 2. 不能包含特殊字符<> ' | \" 且不能只输入/ 3. 第一级目录不支持通配符*：不能以/_**   /_*开头 4.**只能出现一次
	Paths []string `json:"paths"`

	// 采集路径黑名单。 1. 路径必须以/或者字母:\\\\开头 2. 不能包含特殊字符<> ' | \" 且不能只输入/ 3. 第一级目录不支持通配符*：不能以/_**   /_*开头 4.**只能出现一次
	BlackPaths *[]string `json:"black_paths,omitempty"`

	Format *AccessConfigFormatCreate `json:"format"`

	WindowsLogInfo *AccessConfigWindowsLogInfoCreate `json:"windows_log_info,omitempty"`
}

func (o AccessConfigDeatilCreate) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AccessConfigDeatilCreate struct{}"
	}

	return strings.Join([]string{"AccessConfigDeatilCreate", string(data)}, " ")
}
