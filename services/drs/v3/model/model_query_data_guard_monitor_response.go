package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 容灾监控数据响应体
type QueryDataGuardMonitorResponse struct {

	// 带宽。
	Bandwidth *string `json:"bandwidth,omitempty" xml:"bandwidth"`

	// cpu百分比。
	CpuUsedPercent *string `json:"cpuUsed_percent,omitempty" xml:"cpuUsed_percent"`

	// 目标库时延。
	DstDelay *int64 `json:"dst_delay,omitempty" xml:"dst_delay"`

	// 目标io。
	DstIo *string `json:"dst_io,omitempty" xml:"dst_io"`

	// 目标库连接状态。
	DstNormal *bool `json:"dst_normal,omitempty" xml:"dst_normal"`

	// 目标库offSet位点。
	DstOffset *string `json:"dst_offset,omitempty" xml:"dst_offset"`

	// 目标rps。
	DstRps *string `json:"dst_rps,omitempty" xml:"dst_rps"`

	// 内存使用。
	MemUsedInMB *string `json:"mem_used_inMB,omitempty" xml:"mem_used_inMB"`

	// node内存总大小。
	NodeMemInMB *int64 `json:"node_mem_inMB,omitempty" xml:"node_mem_inMB"`

	// 迁移实例offSet位点。
	NodeOffset *string `json:"node_offset,omitempty" xml:"node_offset"`

	// node磁盘总大小。
	NodeVolumeInGB *int64 `json:"node_volume_inGB,omitempty" xml:"node_volume_inGB"`

	// 源库时延。
	SrDelay *int64 `json:"sr_delay,omitempty" xml:"sr_delay"`

	// 源库offSet位点。
	SrOffset *string `json:"sr_offset,omitempty" xml:"sr_offset"`

	// 源io。
	SrcIo *string `json:"src_io,omitempty" xml:"src_io"`

	// 源库连接状态。
	SrcNormal *bool `json:"src_normal,omitempty" xml:"src_normal"`

	// 源rps。
	SrcRps *string `json:"src_rps,omitempty" xml:"src_rps"`

	// 迁移数据量。
	TransInMB *string `json:"trans_inMB,omitempty" xml:"trans_inMB"`

	// 迁移数据行数。
	TransLines *string `json:"trans_lines,omitempty" xml:"trans_lines"`

	// 磁盘使用。
	VolumeUsedInGB *string `json:"volume_used_inGB,omitempty" xml:"volume_used_inGB"`
}

func (o QueryDataGuardMonitorResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "QueryDataGuardMonitorResponse struct{}"
	}

	return strings.Join([]string{"QueryDataGuardMonitorResponse", string(data)}, " ")
}
