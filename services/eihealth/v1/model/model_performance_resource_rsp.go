package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type PerformanceResourceRsp struct {

	// 实例ID
	Id string `json:"id"`

	// 资源ID
	ResourceId string `json:"resource_id"`

	// 实例名称
	Name string `json:"name"`

	Spec *SpecDto `json:"spec"`

	// 可用区
	AvailabilityZoneId string `json:"availability_zone_id"`

	// 最大容量, 单位GB
	Space int32 `json:"space"`

	// 可用容量，单位GB
	FreeSpace float64 `json:"free_space"`

	// 计费模式
	ChargeMode string `json:"charge_mode"`

	// 购买周期
	PeriodNum int32 `json:"period_num"`

	// 作业运行数
	RunningJobCount *int32 `json:"running_job_count,omitempty"`

	// 运行的最大作业数量
	JobQuota *int32 `json:"job_quota,omitempty"`

	// 购买时间
	CreateTime string `json:"create_time"`

	// 失败原因
	FailureReason string `json:"failure_reason"`

	// 状态
	Status string `json:"status"`

	// 资源是否可调度
	Schedulable bool `json:"schedulable"`
}

func (o PerformanceResourceRsp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PerformanceResourceRsp struct{}"
	}

	return strings.Join([]string{"PerformanceResourceRsp", string(data)}, " ")
}
