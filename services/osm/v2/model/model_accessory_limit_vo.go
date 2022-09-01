package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type AccessoryLimitVo struct {

	// 限制文件数量
	LimitCount *string `json:"limit_count,omitempty" xml:"limit_count"`

	// 限制文件大小，单位是M
	LimitSize *string `json:"limit_size,omitempty" xml:"limit_size"`

	// 限制文件类型
	LimitFileType *string `json:"limit_file_type,omitempty" xml:"limit_file_type"`
}

func (o AccessoryLimitVo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AccessoryLimitVo struct{}"
	}

	return strings.Join([]string{"AccessoryLimitVo", string(data)}, " ")
}
