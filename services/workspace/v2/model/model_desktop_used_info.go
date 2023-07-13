package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DesktopUsedInfo 桌面使用时间。
type DesktopUsedInfo struct {

	// 日期，格式：yyyy-MM-dd（UTC时间）。
	Date *string `json:"date,omitempty"`

	// 总共在线时间单位:小时数（h）,精确到两位小数，如：1.32。
	UseTime *string `json:"use_time,omitempty"`
}

func (o DesktopUsedInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DesktopUsedInfo struct{}"
	}

	return strings.Join([]string{"DesktopUsedInfo", string(data)}, " ")
}
