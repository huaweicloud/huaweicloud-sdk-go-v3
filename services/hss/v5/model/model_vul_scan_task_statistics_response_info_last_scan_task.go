package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// VulScanTaskStatisticsResponseInfoLastScanTask **参数解释**: 最后一次漏洞扫描任务的数据 **取值范围**: 不涉及
type VulScanTaskStatisticsResponseInfoLastScanTask struct {

	// **参数解释**: 任务ID **取值范围**: 字符长度1-128位
	Id *string `json:"id,omitempty"`

	// **参数解释**: 扫描任务类型 **取值范围**: - manual：手动扫描 - automatic：自动扫描
	ScanType *string `json:"scan_type,omitempty"`

	// **参数解释**: 任务状态 **取值范围**: - finished：已完成 - running：运行中 - failed：失败
	Status *string `json:"status,omitempty"`

	// **参数解释**: 任务开始时间，Unix时间戳格式，单位为毫秒 **取值范围**: 最小值0，最大值9223372036854775807
	StartTime *int64 `json:"start_time,omitempty"`

	// **参数解释**: 任务结束时间，Unix时间戳格式，单位为毫秒 **取值范围**: 最小值0，最大值9223372036854775807
	EndTime *int64 `json:"end_time,omitempty"`

	// **参数解释**: 扫描的漏洞类型列表 **取值范围**: 0-100个漏洞类型
	ScanVulTypes *[]string `json:"scan_vul_types,omitempty"`

	// **参数解释**: 扫描中的主机数，Serverless场景下为POD实例数 **取值范围**: 最小值0，最大值2147483647
	ScanningHostNum *int32 `json:"scanning_host_num,omitempty"`

	// **参数解释**: 扫描成功的主机数，Serverless场景下为POD实例数 **取值范围**: 最小值0，最大值2147483647
	SuccessHostNum *int32 `json:"success_host_num,omitempty"`

	RiskyHostNum *VulScanTaskStatisticsResponseInfoLastScanTaskRiskyHostNum `json:"risky_host_num,omitempty"`

	// **参数解释**: 扫描失败的主机数，Serverless场景下为POD实例数 **取值范围**: 最小值0，最大值2147483647
	FailedHostNum *int32 `json:"failed_host_num,omitempty"`

	// **参数解释**: 预计还需要多长时间，单位为秒 **取值范围**: 最小值0，最大值2147483647
	EstimatedTime *int32 `json:"estimated_time,omitempty"`
}

func (o VulScanTaskStatisticsResponseInfoLastScanTask) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "VulScanTaskStatisticsResponseInfoLastScanTask struct{}"
	}

	return strings.Join([]string{"VulScanTaskStatisticsResponseInfoLastScanTask", string(data)}, " ")
}
