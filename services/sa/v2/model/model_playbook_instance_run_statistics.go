package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 剧本实例运行数据
type PlaybookInstanceRunStatistics struct {

	// 剧本实例ID
	PlaybookInstanceId *string `json:"playbook_instance_id,omitempty"`

	// 剧本实例名称
	PlaybookInstanceName *string `json:"playbook_instance_name,omitempty"`
}

func (o PlaybookInstanceRunStatistics) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PlaybookInstanceRunStatistics struct{}"
	}

	return strings.Join([]string{"PlaybookInstanceRunStatistics", string(data)}, " ")
}
