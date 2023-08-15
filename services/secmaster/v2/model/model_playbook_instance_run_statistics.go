package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// PlaybookInstanceRunStatistics 剧本实例运行数据
type PlaybookInstanceRunStatistics struct {

	// 剧本实例ID
	PlaybookInstanceId *string `json:"playbook_instance_id,omitempty"`

	// 剧本实例名称
	PlaybookInstanceName *string `json:"playbook_instance_name,omitempty"`

	// 剧本实例运行时间
	PlaybookInstanceRunTime float32 `json:"playbook_instance_run_time,omitempty"`
}

func (o PlaybookInstanceRunStatistics) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PlaybookInstanceRunStatistics struct{}"
	}

	return strings.Join([]string{"PlaybookInstanceRunStatistics", string(data)}, " ")
}
