package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Jobs 作业信息列表
type Jobs struct {

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

func (o Jobs) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Jobs struct{}"
	}

	return strings.Join([]string{"Jobs", string(data)}, " ")
}
