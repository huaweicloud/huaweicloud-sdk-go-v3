package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Job 作业信息列表
type Job struct {

	// 作业ID
	JobId *int32 `json:"job_id,omitempty"`

	// 作业状态
	Status string `json:"status"`

	// 时间戳
	CreateTime int64 `json:"create_time"`

	// 作业异常信息
	Exceptions *string `json:"exceptions,omitempty"`

	// 作业指标信息
	Metrics *string `json:"metrics,omitempty"`

	// 作业执行计划
	Plan *string `json:"plan,omitempty"`
}

func (o Job) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Job struct{}"
	}

	return strings.Join([]string{"Job", string(data)}, " ")
}
