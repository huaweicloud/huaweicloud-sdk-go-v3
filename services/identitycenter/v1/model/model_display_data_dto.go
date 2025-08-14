package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DisplayDataDto 应用程序提供商显示信息
type DisplayDataDto struct {

	// 应用程序提供商描述
	Description *string `json:"description,omitempty"`

	// 应用程序提供商显示名
	DisplayName *string `json:"display_name,omitempty"`

	// 应用程序提供商图标
	IconUrl *string `json:"icon_url,omitempty"`
}

func (o DisplayDataDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DisplayDataDto struct{}"
	}

	return strings.Join([]string{"DisplayDataDto", string(data)}, " ")
}
