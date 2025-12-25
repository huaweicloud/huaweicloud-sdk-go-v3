package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// IsapResourceDisplaySetting 显示设置
type IsapResourceDisplaySetting struct {

	// 列
	Columns *[]IsapResourceDisplaySettingColumns `json:"columns,omitempty"`

	// 显示格式，例如 TABLE
	Format *string `json:"format,omitempty"`
}

func (o IsapResourceDisplaySetting) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "IsapResourceDisplaySetting struct{}"
	}

	return strings.Join([]string{"IsapResourceDisplaySetting", string(data)}, " ")
}
