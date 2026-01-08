package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CpuOptions struct {

	// 是否关闭实例超线程，1是关闭，2是开启
	HwCpuThreads *int32 `json:"hw_cpu_threads,omitempty"`
}

func (o CpuOptions) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CpuOptions struct{}"
	}

	return strings.Join([]string{"CpuOptions", string(data)}, " ")
}
