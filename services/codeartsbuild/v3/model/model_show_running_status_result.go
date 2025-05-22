package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowRunningStatusResult 构建成功率
type ShowRunningStatusResult struct {

	// 任务ID
	LastJobId *string `json:"last_job_id,omitempty"`

	// 最新构建编号
	LastBuildNo *string `json:"last_build_no,omitempty"`

	// 最新构建状态
	LastBuildStatus *string `json:"last_build_status,omitempty"`

	// 是否在运行中
	IsRunning *bool `json:"is_running,omitempty"`
}

func (o ShowRunningStatusResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowRunningStatusResult struct{}"
	}

	return strings.Join([]string{"ShowRunningStatusResult", string(data)}, " ")
}
