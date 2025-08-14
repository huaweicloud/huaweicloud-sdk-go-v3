package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DisplayDto 应用程序显示信息
type DisplayDto struct {

	// 应用程序描述
	Description string `json:"description"`

	// 应用程序显示名称
	DisplayName string `json:"display_name"`

	// 应用程序图标
	Icon *string `json:"icon,omitempty"`
}

func (o DisplayDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DisplayDto struct{}"
	}

	return strings.Join([]string{"DisplayDto", string(data)}, " ")
}
