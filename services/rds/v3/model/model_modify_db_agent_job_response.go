package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ModifyDbAgentJobResponse Response Object
type ModifyDbAgentJobResponse struct {

	// 作业id。
	JobId *string `json:"job_id,omitempty"`

	// 作业名。
	JobName *string `json:"job_name,omitempty"`

	// 是否可用。  true：可用。 false：不可用。
	IsEnabled *bool `json:"is_enabled,omitempty"`

	// 最新执行时间。格式为“yyyy-mm-ddThh:mm:ssZ”。 其中，T指某个时间的开始；Z指时区偏移量，例如北京时间偏移显示为+0800。
	RunTime *string `json:"run_time,omitempty"`

	// 作业执行状态。取值如下:  0:失败。 1:成功。 2:重试中。 3:已取消。 4:正在运行。
	RunStatus *string `json:"run_status,omitempty"`

	// 最近失败时间。格式为“yyyy-mm-ddThh:mm:ssZ”。 其中，T指某个时间的开始；Z指时区偏移量，例如北京时间偏移显示为+0800。
	LastFailureTime *string `json:"last_failure_time,omitempty"`

	// 历史失败次数。
	FailureCount *int32 `json:"failure_count,omitempty"`

	// 作业代理的类型。  snapshot：快照代理 log_reader：日志读取器代理 distribution：分发代理 merge:合并代理 queue_reader：队列读取器代理。
	AgentType *string `json:"agent_type,omitempty"`

	// 配置文件id。
	ProfileId *string `json:"profile_id,omitempty"`

	// 配置文件名。
	ProfileName    *string `json:"profile_name,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ModifyDbAgentJobResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ModifyDbAgentJobResponse struct{}"
	}

	return strings.Join([]string{"ModifyDbAgentJobResponse", string(data)}, " ")
}
