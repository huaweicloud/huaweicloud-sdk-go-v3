package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowVulScanTaskEstimatedTimeResponse Response Object
type ShowVulScanTaskEstimatedTimeResponse struct {

	// **参数解释**： 执行漏洞扫描需要排队等待的时间，单位为分钟 **约束限制**： 不涉及 **取值范围**： 字符长度1-9223372036854775807 **默认取值**： 不涉及
	QueueTime *int64 `json:"queue_time,omitempty"`

	// **参数解释**: 本次任务执行需要的时间，单位为分钟 **约束限制**： 不涉及 **取值范围**： 字符长度1-9223372036854775807 **默认取值**： 不涉及
	RunTime *int64 `json:"run_time,omitempty"`

	// **参数解释**: 从现在到任务结束需要的总时间，单位为分钟 **约束限制**： 不涉及 **取值范围**： 字符长度1-9223372036854775807 **默认取值**： 不涉及
	TotalTime      *int32 `json:"total_time,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ShowVulScanTaskEstimatedTimeResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowVulScanTaskEstimatedTimeResponse struct{}"
	}

	return strings.Join([]string{"ShowVulScanTaskEstimatedTimeResponse", string(data)}, " ")
}
