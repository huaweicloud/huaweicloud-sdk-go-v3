package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowStatusInstanceResponse Response Object
type ShowStatusInstanceResponse struct {

	// CPU使用率
	CpuUtils *[]ShowStatusInstanceItem `json:"cpu_utils,omitempty"`

	// 内存使用率
	MemUtils *[]ShowStatusInstanceItem `json:"mem_utils,omitempty"`

	// 磁盘使用率
	DiskUtils *[]ShowStatusInstanceItem `json:"disk_utils,omitempty"`

	// CPU使用率对应实例名称
	CpuUtilInstanceNames *[]string `json:"cpu_util_instance_names,omitempty"`

	// 内存使用率对应实例名称
	MemUtilInstanceNames *[]string `json:"mem_util_instance_names,omitempty"`

	// 磁盘使用率对应实例名称
	DiskUtilInstanceNames *[]string `json:"disk_util_instance_names,omitempty"`

	// CPU使用率平均值
	CpuUtilAverages *[]float64 `json:"cpu_util_averages,omitempty"`

	// 内存使用率平均值
	MemUtilAverages *[]float64 `json:"mem_util_averages,omitempty"`

	// 磁盘使用率平均值
	DiskUtilAverages *[]float64 `json:"disk_util_averages,omitempty"`
	HttpStatusCode   int        `json:"-"`
}

func (o ShowStatusInstanceResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowStatusInstanceResponse struct{}"
	}

	return strings.Join([]string{"ShowStatusInstanceResponse", string(data)}, " ")
}
