package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type NextflowTaskResourceUsage struct {

	// cpu占用率
	CpuPercent *float64 `json:"cpu_percent,omitempty"`

	// 内存占用率
	MemPercent *float64 `json:"mem_percent,omitempty"`

	// 读取字符数
	Rchar *int64 `json:"rchar,omitempty"`

	// 写入字符数
	Wchar *int64 `json:"wchar,omitempty"`

	// 读取字节数
	ReadBytes *int64 `json:"read_bytes,omitempty"`

	// 写入字符数
	WriteBytes *int64 `json:"write_bytes,omitempty"`

	// process虚拟内存大小
	Vmem *int64 `json:"vmem,omitempty"`

	// process实际内存大小
	Rss *int64 `json:"rss,omitempty"`

	// process虚拟内存峰值
	PeakVmem *int64 `json:"peak_vmem,omitempty"`

	// process实际内存峰值
	PeakRss *int64 `json:"peak_rss,omitempty"`

	// 系统调用次数
	Syscr *int64 `json:"syscr,omitempty"`

	// 系统调用次数
	Syscw *int64 `json:"syscw,omitempty"`

	// 自愿上下文切换数
	VolCtxt *int64 `json:"vol_ctxt,omitempty"`

	// 非自愿上下文切换数
	InvCtxt *int64 `json:"inv_ctxt,omitempty"`
}

func (o NextflowTaskResourceUsage) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NextflowTaskResourceUsage struct{}"
	}

	return strings.Join([]string{"NextflowTaskResourceUsage", string(data)}, " ")
}
