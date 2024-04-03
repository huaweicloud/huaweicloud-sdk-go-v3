package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type AdvancedQueueProperty struct {

	// 队列能启动的最大spark driver数量
	ComputeEngineMaxInstance *int32 `json:"computeEngine.maxInstance,omitempty"`

	// 单个spark driver能同时运行的最大任务数量
	JobMaxConcurrent *int32 `json:"job.maxConcurrent,omitempty"`

	// 队列预先启动的最大spark driver数量
	ComputeEngineMaxPrefetchInstance *int32 `json:"computeEngine.maxPrefetchInstance,omitempty"`

	// 设置队列网段
	NetworkCidrInVpc *string `json:"network.cidrInVpc,omitempty"`
}

func (o AdvancedQueueProperty) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AdvancedQueueProperty struct{}"
	}

	return strings.Join([]string{"AdvancedQueueProperty", string(data)}, " ")
}
