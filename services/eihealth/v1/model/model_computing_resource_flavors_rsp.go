package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ComputingResourceFlavorsRsp struct {

	// 规格编码
	Code string `json:"code"`

	// 规格名称
	Name string `json:"name"`

	// 内存
	Ram int64 `json:"ram"`

	// vcpus
	Vcpus string `json:"vcpus"`

	// 最大带宽
	MaxRate *string `json:"max_rate,omitempty"`

	// 基准带宽
	MinRate *string `json:"min_rate,omitempty"`

	// 最大收发包能力
	MaxPps *string `json:"max_pps,omitempty"`

	// 是否售罄
	SoldOut bool `json:"sold_out"`

	// 可用区
	Az *string `json:"az,omitempty"`

	// CPU物理规格描述信息
	CpuDetail *string `json:"cpu_detail,omitempty"`

	// 磁盘物理规格描述信息
	DiskDetail *string `json:"disk_detail,omitempty"`

	// 内存物理规格描述信息
	MemoryDetail *string `json:"memory_detail,omitempty"`

	// 网卡物理规格描述信息
	NetcardDetail *string `json:"netcard_detail,omitempty"`

	// 裸金属服务器的CPU架构类型
	CpuArch *string `json:"cpu_arch,omitempty"`
}

func (o ComputingResourceFlavorsRsp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ComputingResourceFlavorsRsp struct{}"
	}

	return strings.Join([]string{"ComputingResourceFlavorsRsp", string(data)}, " ")
}
