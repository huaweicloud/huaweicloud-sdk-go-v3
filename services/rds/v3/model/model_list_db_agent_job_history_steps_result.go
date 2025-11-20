package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListDbAgentJobHistoryStepsResult 数据库代理作业执行历史步骤。
type ListDbAgentJobHistoryStepsResult struct {

	// 步骤id。
	StepId *string `json:"step_id,omitempty"`

	// 步骤名。
	StepName *string `json:"step_name,omitempty"`

	// 作业执行状态。取值如下:  failed:失败。 succeeded:成功。 retrying:重试中。 canceled:已取消。 in_progress:正在运行。
	RunStatus *string `json:"run_status,omitempty"`

	// 作业执行开始时间。格式为“yyyy-mm-ddThh:mm:ssZ”。 其中，T指某个时间的开始；Z指时区偏移量，例如北京时间偏移显示为+0800。
	RunTime *string `json:"run_time,omitempty"`

	// 作业执行时长。格式为hh:mm:ss
	RunDuration *string `json:"run_duration,omitempty"`

	// 执行信息。
	Message *string `json:"message,omitempty"`
}

func (o ListDbAgentJobHistoryStepsResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListDbAgentJobHistoryStepsResult struct{}"
	}

	return strings.Join([]string{"ListDbAgentJobHistoryStepsResult", string(data)}, " ")
}
