package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// PerformanceResourcesRsp 性能加速资源信息
type PerformanceResourcesRsp struct {

	// 性能加速资源id
	Id *string `json:"id,omitempty"`

	// 性能加速资源名称
	Name *string `json:"name,omitempty"`

	// 当前运行中的作业总数
	RunningJobCount *int32 `json:"running_job_count,omitempty"`

	// 资源是否可调度
	Schedulable *bool `json:"schedulable,omitempty"`
}

func (o PerformanceResourcesRsp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PerformanceResourcesRsp struct{}"
	}

	return strings.Join([]string{"PerformanceResourcesRsp", string(data)}, " ")
}
