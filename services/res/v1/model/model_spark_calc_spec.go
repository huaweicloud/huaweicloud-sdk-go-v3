package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SparkCalcSpec spark自定义计算规格
type SparkCalcSpec struct {

	// driver内存
	DriverMemory string `json:"driver_memory"`

	// driver核数
	DriverCores int32 `json:"driver_cores"`

	// executor内存
	ExecutorMemory string `json:"executor_memory"`

	// executor核数
	ExecutorCores int32 `json:"executor_cores"`

	// executor个数
	NumExecutors int32 `json:"num_executors"`
}

func (o SparkCalcSpec) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SparkCalcSpec struct{}"
	}

	return strings.Join([]string{"SparkCalcSpec", string(data)}, " ")
}
