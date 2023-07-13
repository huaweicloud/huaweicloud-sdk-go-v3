package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type DiskResp struct {

	// 实例名称
	InstanceName *string `json:"instance_name,omitempty"`

	// 主机名称
	HostName *string `json:"host_name,omitempty"`

	// 磁盘名称
	DiskName *string `json:"disk_name,omitempty"`

	// 磁盘类型(系统盘、数据盘、日志盘)。
	DiskType *string `json:"disk_type,omitempty"`

	// 磁盘总容量(GB)。
	Total *float64 `json:"total,omitempty"`

	// 磁盘已使用容量(GB)。
	Used *float64 `json:"used,omitempty"`

	// 磁盘可用容量(GB)。
	Available *float64 `json:"available,omitempty"`

	// 磁盘使用率(%)。
	UsedPercentage *float64 `json:"used_percentage,omitempty"`

	// IO等待时间(ms)。
	Await *float64 `json:"await,omitempty"`

	// IO服务时间(ms)。
	Svctm *float64 `json:"svctm,omitempty"`

	// IO使用率(%)。
	Util *float64 `json:"util,omitempty"`

	// 磁盘读速率(KB/s)。
	ReadRate *float64 `json:"read_rate,omitempty"`

	// 磁盘写速率(KB/s)。
	WriteRate *float64 `json:"write_rate,omitempty"`
}

func (o DiskResp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DiskResp struct{}"
	}

	return strings.Join([]string{"DiskResp", string(data)}, " ")
}
