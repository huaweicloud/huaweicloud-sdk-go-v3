package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type IgnorePathSettingItem struct {

	// 目录或文件路径
	FilePath string `json:"file_path"`

	// 屏蔽状态，包括unchecked(不屏蔽)、all(全屏蔽)、half(半屏蔽)
	CheckboxStatus string `json:"checkbox_status"`
}

func (o IgnorePathSettingItem) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "IgnorePathSettingItem struct{}"
	}

	return strings.Join([]string{"IgnorePathSettingItem", string(data)}, " ")
}
