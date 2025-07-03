package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type IncidentSimpleTicketInfo struct {

	// 事件单标题
	Title *string `json:"title,omitempty"`

	// 事件单单号
	TicketId *string `json:"ticket_id,omitempty"`

	// 事件单描述
	Description *string `json:"description,omitempty"`

	// 事件状态 参考：枚举 [事件状态](coc_api_04_03_001_006.xml) incident_status
	Status *string `json:"status,omitempty"`

	// 事件级别 参考：枚举 [事件级别](coc_api_04_03_001_006.xml) incident_level
	Level *string `json:"level,omitempty"`

	// 企业项目ID
	EnterpriseProject *string `json:"enterprise_project,omitempty"`

	// 单据来源 参考：枚举 [事件来源](coc_api_04_03_001_006.xml) incident_source
	Source *string `json:"source,omitempty"`

	// 创建人工号
	Creator *string `json:"creator,omitempty"`

	// 当前责任人，可能存在多个责任人使用列表展示，内容为责任人用户ID信息
	Assignee *[]string `json:"assignee,omitempty"`
}

func (o IncidentSimpleTicketInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "IncidentSimpleTicketInfo struct{}"
	}

	return strings.Join([]string{"IncidentSimpleTicketInfo", string(data)}, " ")
}
