package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RingHost 集群主机信息
type RingHost struct {

	// 主机名称
	HostName string `json:"host_name"`

	// 后端IP地址
	BackIp string `json:"back_ip"`

	// 主机CPU核数
	CpuCores int32 `json:"cpu_cores"`

	// 主机内存
	Memory float64 `json:"memory"`

	// 主机磁盘大小
	DiskSize float64 `json:"disk_size"`
}

func (o RingHost) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RingHost struct{}"
	}

	return strings.Join([]string{"RingHost", string(data)}, " ")
}
