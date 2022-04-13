package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ListAlarmsRequest struct {
	// 告警规则ID

	AlarmId *string `json:"alarm_id,omitempty"`
	// 告警规则名称

	Name *string `json:"name,omitempty"`
	// 命名空间

	Namespace *string `json:"namespace,omitempty"`
	// 维度的拼接串Dimension1,Dimension2...

	ResourceId *string `json:"resource_id,omitempty"`
	// 企业项目ID

	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`
	// 偏移量

	Offset *int32 `json:"offset,omitempty"`
	// 希望的查询的数据量

	Limit *int32 `json:"limit,omitempty"`
}

func (o ListAlarmsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAlarmsRequest struct{}"
	}

	return strings.Join([]string{"ListAlarmsRequest", string(data)}, " ")
}
