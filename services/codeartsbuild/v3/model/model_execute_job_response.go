package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ExecuteJobResponse Response Object
type ExecuteJobResponse struct {

	// 临时任务名称
	OctopusJobName *string `json:"octopus_job_name,omitempty"`

	// 实际构建次数
	ActualBuildNumber *string `json:"actual_build_number,omitempty"`

	// 构建每日编号
	DailyBuildNumber *string `json:"daily_build_number,omitempty"`
	HttpStatusCode   int     `json:"-"`
}

func (o ExecuteJobResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExecuteJobResponse struct{}"
	}

	return strings.Join([]string{"ExecuteJobResponse", string(data)}, " ")
}
