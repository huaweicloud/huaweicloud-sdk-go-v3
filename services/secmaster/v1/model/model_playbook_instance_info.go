package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// PlaybookInstanceInfo 剧本实例列表响应
type PlaybookInstanceInfo struct {

	// 剧本实例ID
	Id *string `json:"id,omitempty"`

	// 剧本实例名称
	Name *string `json:"name,omitempty"`

	// 项目ID
	ProjectId *string `json:"project_id,omitempty"`

	Playbook *PlaybookInfoRef `json:"playbook,omitempty"`

	Dataclass *DataclassInfoRef `json:"dataclass,omitempty"`

	Dataobject *DataobjectRef `json:"dataobject,omitempty"`

	// 剧本实例状态. (RUNNING--运行中、FINISHED--成功、FAILED--失败、RETRYING--重试中、TERMINATING--终止中、TERMINATED--已终止)
	Status *string `json:"status,omitempty"`

	// 触发类型. TIMER--定时触发, EVENT--事件触发
	TriggerType *string `json:"trigger_type,omitempty"`

	// 创建时间
	StartTime *string `json:"start_time,omitempty"`

	// 更新时间
	EndTime *string `json:"end_time,omitempty"`
}

func (o PlaybookInstanceInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PlaybookInstanceInfo struct{}"
	}

	return strings.Join([]string{"PlaybookInstanceInfo", string(data)}, " ")
}
