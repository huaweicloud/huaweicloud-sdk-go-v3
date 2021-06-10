package model

import (
	"encoding/json"

	"strings"
)

// 容灾监控数据响应体
type QueryDataGuardMonitorResponse struct {
	// 带宽。

	Bandwidth *string `json:"bandwidth,omitempty"`
	// cpu百分比。

	CpuUsedPercent *string `json:"cpuUsed_percent,omitempty"`
	// 目标库时延。

	DstDelay *int64 `json:"dst_delay,omitempty"`
	// 目标io。

	DstIo *string `json:"dst_io,omitempty"`
	// 目标库连接状态。

	DstNormal *bool `json:"dst_normal,omitempty"`
	// 目标库offSet位点。

	DstOffset *string `json:"dst_offset,omitempty"`
	// 目标rps。

	DstRps *string `json:"dst_rps,omitempty"`
	// 内存使用。

	MemUsedInMB *string `json:"mem_used_inMB,omitempty"`
	// node内存总大小。

	NodeMemInMB *int64 `json:"node_mem_inMB,omitempty"`
	// 迁移实例offSet位点。

	NodeOffset *string `json:"node_offset,omitempty"`
	// node磁盘总大小。

	NodeVolumeInGB *int64 `json:"node_volume_inGB,omitempty"`
	// 源库时延。

	SrDelay *int64 `json:"sr_delay,omitempty"`
	// 源库offSet位点。

	SrOffset *string `json:"sr_offset,omitempty"`
	// 源io。

	SrcIo *string `json:"src_io,omitempty"`
	// 源库连接状态。

	SrcNormal *bool `json:"src_normal,omitempty"`
	// 源rps。

	SrcRps *string `json:"src_rps,omitempty"`
	// 迁移数据量。

	TransInMB *string `json:"trans_inMB,omitempty"`
	// 迁移数据行数。

	TransLines *string `json:"trans_lines,omitempty"`
	// 磁盘使用。

	VolumeUsedInGB *string `json:"volume_used_inGB,omitempty"`
}

func (o QueryDataGuardMonitorResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "QueryDataGuardMonitorResponse struct{}"
	}

	return strings.Join([]string{"QueryDataGuardMonitorResponse", string(data)}, " ")
}
