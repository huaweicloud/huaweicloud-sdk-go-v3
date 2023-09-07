package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowMonitorDataResponse Response Object
type ShowMonitorDataResponse struct {

	// EIP带宽，单位：MB/S
	Bandwidth *string `json:"bandwidth,omitempty"`

	// 源库连接状态是否正常。
	IsSrcNormal *bool `json:"is_src_normal,omitempty"`

	// 目标库连接状态是否正常。
	IsDstNormal *bool `json:"is_dst_normal,omitempty"`

	// 源库offSet位点。
	SrcOffset *string `json:"src_offset,omitempty"`

	// 迁移实例offSet位点。
	NodeOffset *string `json:"node_offset,omitempty"`

	// 目标库offSet位点。
	DstOffset *string `json:"dst_offset,omitempty"`

	// 源库时延。
	SrcDelay *int32 `json:"src_delay,omitempty"`

	// 目标库时延。
	DstDelay *int32 `json:"dst_delay,omitempty"`

	// 源库RPS。
	SrcRps *string `json:"src_rps,omitempty"`

	// 源库IO。
	SrcIo *string `json:"src_io,omitempty"`

	// 目标库RPS。
	DstRps *string `json:"dst_rps,omitempty"`

	// 目标库IO。
	DstIo *string `json:"dst_io,omitempty"`

	// 迁移数据量。单位：MB
	TransData *string `json:"trans_data,omitempty"`

	// 迁移数据行数。
	TransLines *string `json:"trans_lines,omitempty"`

	// 磁盘使用量。单位：GB
	UsedVolumes *string `json:"used_volumes,omitempty"`

	// 内存使用量。单位：MB
	UsedMemory *string `json:"used_memory,omitempty"`

	// CPU使用百分比。
	UsedCpuPercent *string `json:"used_cpu_percent,omitempty"`

	// node磁盘总大小。单位：GB
	NodeVolumeSize *int32 `json:"node_volume_size,omitempty"`

	// node内存总大小。单位：MB
	NodeMemorySize *int32 `json:"node_memory_size,omitempty"`

	// 更新时间。
	UpdateTime *string `json:"update_time,omitempty"`

	// 同步速度。单位：byte/s
	ApplyRate      *int32 `json:"apply_rate,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ShowMonitorDataResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowMonitorDataResponse struct{}"
	}

	return strings.Join([]string{"ShowMonitorDataResponse", string(data)}, " ")
}
