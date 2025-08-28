package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type VulUrgentResponseInfoDataList struct {

	// **参数解释**: 漏洞标签列表 **取值范围**: 最小值0，最大值2147483647
	LabelList *[]string `json:"label_list,omitempty"`

	// **参数解释**: 修复优先级 **取值范围**: - Critical：紧急。 - High：高。 - Medium：中。 - Low：低。
	RepairPriority *string `json:"repair_priority,omitempty"`

	// **参数解释**: 当前漏洞修复成功次数 **取值范围**: 最小值0，最大值1000000
	RepairSuccessNum *int32 `json:"repair_success_num,omitempty"`

	// **参数解释**: 漏洞ID **取值范围**: 字符范围0-64位
	VulId *string `json:"vul_id,omitempty"`

	// **参数解释**: 漏洞名称 **取值范围**: 字符范围0-256位
	VulName *string `json:"vul_name,omitempty"`

	// **参数解释**: CVE列表 **取值范围**: 最小值1，最大值10000
	CveList *[]VulUrgentResponseInfoCveList `json:"cve_list,omitempty"`

	// **参数解释**: 是否影响业务 **取值范围**: - true：是。 - false：否。
	IsAffectedBusiness *bool `json:"is_affected_business,omitempty"`

	// **参数解释**: 主机ID **取值范围**: 字符长度1-64
	HostId *string `json:"host_id,omitempty"`

	// **参数解释**: 服务器公网IP **取值范围**: 字符长度0-128
	PublicIp *string `json:"public_ip,omitempty"`

	// **参数解释**: 服务器私网IP **取值范围**: 字符长度0-128
	PrivateIp *string `json:"private_ip,omitempty"`

	// **参数解释**: 主机名称 **取值范围**: 字符长度1-128位
	HostName *string `json:"host_name,omitempty"`

	// **参数解释**: 修复优先级 **取值范围**: - Critical：紧急。 - High：高。 - Medium：中。 - Low：低。
	AssetValue *string `json:"asset_value,omitempty"`

	// **参数解释**： 漏洞状态 **取值范围**： 字符长度0-32位
	Status *string `json:"status,omitempty"`

	// **参数解释**: 首次扫描时间 **取值范围**: 最小值0，最大值9223372036854775807
	FirstScanTime *int64 `json:"first_scan_time,omitempty"`

	// **参数解释**: 最近扫描时间 **取值范围**: 最小值0，最大值9223372036854775807
	ScanTime *int64 `json:"scan_time,omitempty"`

	// **参数解释**: 该漏洞不可进行的操作类型列表 **取值范围**: 最小值1，最大值10000
	DisabledOperateTypes *[]VulUrgentResponseInfoDisabledOperateTypes `json:"disabled_operate_types,omitempty"`
}

func (o VulUrgentResponseInfoDataList) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "VulUrgentResponseInfoDataList struct{}"
	}

	return strings.Join([]string{"VulUrgentResponseInfoDataList", string(data)}, " ")
}
