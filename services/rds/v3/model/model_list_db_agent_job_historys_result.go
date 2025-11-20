package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListDbAgentJobHistorysResult 数据库代理作业执行历史步骤。
type ListDbAgentJobHistorysResult struct {

	// 作业历史记录id。
	HistoryId *string `json:"history_id,omitempty"`

	// 作业执行状态。取值如下:  failed:失败。 succeeded:成功。 retrying:重试中。 canceled:已取消。 in_progress:正在运行。
	RunStatus *string `json:"run_status,omitempty"`

	// 作业执行开始时间。格式为“yyyy-mm-ddThh:mm:ssZ”。 其中，T指某个时间的开始；Z指时区偏移量，例如北京时间偏移显示为+0800。
	RunTime *string `json:"run_time,omitempty"`

	// 作业执行时长。格式为hh:mm:ss
	RunDuration *string `json:"run_duration,omitempty"`

	// 执行信息。
	Message *string `json:"message,omitempty"`
}

func (o ListDbAgentJobHistorysResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListDbAgentJobHistorysResult struct{}"
	}

	return strings.Join([]string{"ListDbAgentJobHistorysResult", string(data)}, " ")
}
