package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DesktopSysprepInfo 查询sysprep版本信息响应。
type DesktopSysprepInfo struct {

	// 桌面id。
	DesktopId *string `json:"desktop_id,omitempty"`

	// sysprep版本。
	SysprepVersion *string `json:"sysprep_version,omitempty"`

	// 是否支持休眠。
	SupportHibernate *bool `json:"support_hibernate,omitempty"`

	// 是否支持修改路由。
	SupportUpdateRoute *bool `json:"support_update_route,omitempty"`
}

func (o DesktopSysprepInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DesktopSysprepInfo struct{}"
	}

	return strings.Join([]string{"DesktopSysprepInfo", string(data)}, " ")
}
