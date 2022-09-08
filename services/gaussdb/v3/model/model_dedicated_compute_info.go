package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type DedicatedComputeInfo struct {

	// 专属资源池中cpu总数
	VcpusTotal int32 `json:"vcpus_total"`

	// 专属资源池已使用的cpu数
	VcpusUsed int32 `json:"vcpus_used"`

	// 专属资源池计算内存大小, 单位GB
	RamTotal int32 `json:"ram_total"`

	// 专属资源池已使用的计算内存大小，单位GB
	RamUsed int32 `json:"ram_used"`

	// 专属资源池计算资源规格码
	SpecCode string `json:"spec_code"`

	// 专属资源池计算主机数量
	HostNum int32 `json:"host_num"`
}

func (o DedicatedComputeInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DedicatedComputeInfo struct{}"
	}

	return strings.Join([]string{"DedicatedComputeInfo", string(data)}, " ")
}
