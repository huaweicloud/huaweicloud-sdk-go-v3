package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// VulHostAppsResponseInfoDataList **参数解释**: 软件信息 **取值范围**: 不涉及
type VulHostAppsResponseInfoDataList struct {

	// **参数解释**: app名称 **取值范围**: 字符长度1-128位
	AppName *string `json:"app_name,omitempty"`

	// **参数解释**: 存在漏洞的软件路径，仅应用漏洞、应急漏洞等存在软件路径的漏洞类型查询时存在该字段（注：该字段即将弃用，请使用paths字段获取软件路径相关信息） **取值范围**: 字符长度1-512位
	Path *string `json:"path,omitempty"`

	// **参数解释**: 存在漏洞的软件路径列表，仅应用漏洞、应急漏洞等存在软件路径的漏洞类型查询时存在该字段 **取值范围**: 不涉及
	Paths *[]VulHostAppsResponseInfoPaths `json:"paths,omitempty"`

	// **参数解释**: 软件命中的漏洞匹配规则 **取值范围**: 字符长度1-512位
	Rule *string `json:"rule,omitempty"`
}

func (o VulHostAppsResponseInfoDataList) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "VulHostAppsResponseInfoDataList struct{}"
	}

	return strings.Join([]string{"VulHostAppsResponseInfoDataList", string(data)}, " ")
}
