package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowVulIndividualStatisticsResponse Response Object
type ShowVulIndividualStatisticsResponse struct {

	// **参数解释**: 需要紧急修复的漏洞数量 **取值范围**: 最小值0，最大值2147483647
	NeedUrgentRepairNum *int32 `json:"need_urgent_repair_num,omitempty"`

	// **参数解释**: 未完成修复的漏洞数量 **取值范围**: 最小值0，最大值2147483647
	UnhandledNum *int32 `json:"unhandled_num,omitempty"`

	// **参数解释**: 存在漏洞的服务器数量 **取值范围**: 最小值0，最大值2147483647
	ExistedVulHostNum *int32 `json:"existed_vul_host_num,omitempty"`

	// **参数解释**: 今日处理的漏洞数量 **取值范围**: 最小值0，最大值2147483647
	TodayHandledNum *int32 `json:"today_handled_num,omitempty"`

	// **参数解释**: 累计处理的漏洞数量 **取值范围**: 最小值0，最大值2147483647
	AllHandledNum *int32 `json:"all_handled_num,omitempty"`

	// **参数解释**: 漏洞库已支持漏洞数 **取值范围**: 最小值0，最大值2147483647
	VulLibrarySupportedNum *int32 `json:"vul_library_supported_num,omitempty"`

	// **参数解释**: 漏洞库更新时间，Unix时间戳格式，单位为毫秒 **取值范围**: 最小值0，最大值9223372036854775807
	VulLibraryUpdateTime *int64 `json:"vul_library_update_time,omitempty"`

	// **参数解释**: 未处置的Linux漏洞数量 **取值范围**: 最小值0，最大值2147483647
	LinuxVulUnhandledNum *int32 `json:"linux_vul_unhandled_num,omitempty"`

	// **参数解释**: 未处置的Windows漏洞数量 **取值范围**: 最小值0，最大值2147483647
	WindowsVulUnhandledNum *int32 `json:"windows_vul_unhandled_num,omitempty"`

	// **参数解释**: 未处置的Web CMS漏洞数量 **取值范围**: 最小值0，最大值2147483647
	WebCmsVulUnhandledNum *int32 `json:"web_cms_vul_unhandled_num,omitempty"`

	// **参数解释**: 未处置的应用漏洞数量 **取值范围**: 最小值0，最大值2147483647
	AppVulUnhandledNum *int32 `json:"app_vul_unhandled_num,omitempty"`

	// **参数解释**: 未处置的应急漏洞数量 **取值范围**: 最小值0，最大值2147483647
	UrgentVulUnhandledNum *int32 `json:"urgent_vul_unhandled_num,omitempty"`

	// **参数解释**: 修复成功待重启的主机数量 **取值范围**: 最小值0，最大值2147483647
	NeedRebootHostNum *int32 `json:"need_reboot_host_num,omitempty"`

	// **参数解释**: 单次最大可处置的漏洞数量 **取值范围**: 最小值0，最大值2147483647
	MaxVulHandleNum *int32 `json:"max_vul_handle_num,omitempty"`
	HttpStatusCode  int    `json:"-"`
}

func (o ShowVulIndividualStatisticsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowVulIndividualStatisticsResponse struct{}"
	}

	return strings.Join([]string{"ShowVulIndividualStatisticsResponse", string(data)}, " ")
}
