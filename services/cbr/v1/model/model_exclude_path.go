package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ExcludePath struct {

	// 备份目录
	PathName *string `json:"path_name,omitempty"`

	// 排除目录列表
	ExcludePathName *[]string `json:"exclude_path_name,omitempty"`
}

func (o ExcludePath) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExcludePath struct{}"
	}

	return strings.Join([]string{"ExcludePath", string(data)}, " ")
}
