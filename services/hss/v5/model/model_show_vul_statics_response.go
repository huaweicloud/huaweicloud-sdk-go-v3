package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowVulStaticsResponse Response Object
type ShowVulStaticsResponse struct {

	// **参数解释**: 需紧急修复的漏洞数 **取值范围**: 取值0-2147483647
	NeedUrgentRepair *int32 `json:"need_urgent_repair,omitempty"`

	// **参数解释**: 未完成修复的漏洞数 **取值范围**: 取值0-2147483647
	Unrepair *int32 `json:"unrepair,omitempty"`

	// **参数解释**: 存在漏洞的服务器数 **取值范围**: 取值0-2147483647
	ExistedVulHosts *int32 `json:"existed_vul_hosts,omitempty"`

	// **参数解释**: 今日处理漏洞数 **取值范围**: 取值0-2147483647
	TodayHandle *int32 `json:"today_handle,omitempty"`

	// **参数解释**: 累计处理漏洞数 **取值范围**: 取值0-2147483647
	AllHandle *int32 `json:"all_handle,omitempty"`

	// **参数解释**: 已支持漏洞数 **取值范围**: 取值0-2147483647
	Supported *int32 `json:"supported,omitempty"`

	// **参数解释**: 漏洞库更新时间（时间戳，单位为毫秒） **取值范围**: 取值0-9223372036854775807
	VulLibraryUpdateTime *int64 `json:"vul_library_update_time,omitempty"`
	HttpStatusCode       int    `json:"-"`
}

func (o ShowVulStaticsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowVulStaticsResponse struct{}"
	}

	return strings.Join([]string{"ShowVulStaticsResponse", string(data)}, " ")
}
