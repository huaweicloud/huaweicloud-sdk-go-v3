package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type RunJobResponse struct {

	// 临时任务名称
	OctopusJobName *string `json:"octopus_job_name,omitempty" xml:"octopus_job_name"`

	// 实际构建次数
	ActualBuildNumber *string `json:"actual_build_number,omitempty" xml:"actual_build_number"`

	// 构建每日编号
	DailyBuildNumber *string `json:"daily_build_number,omitempty" xml:"daily_build_number"`
	HttpStatusCode   int     `json:"-"`
}

func (o RunJobResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RunJobResponse struct{}"
	}

	return strings.Join([]string{"RunJobResponse", string(data)}, " ")
}
