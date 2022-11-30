package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 告警行动规则
type ActionRule struct {

	// 规则名称
	RuleName string `json:"rule_name"`

	// 项目ID
	ProjectId string `json:"project_id"`

	// 子账号名称
	UserName string `json:"user_name"`

	// 规则描述
	Desc *string `json:"desc,omitempty"`

	// 规则类型。1：通知，2：用户
	Type string `json:"type"`

	// 消息模板
	NotificationTemplate string `json:"notification_template"`

	// 创建时间
	CreateTime *int64 `json:"create_time,omitempty"`

	// 修改时间
	UpdateTime *int64 `json:"update_time,omitempty"`

	// 时区
	TimeZone *string `json:"time_zone,omitempty"`

	// SMN主题信息，不能大于5
	SmnTopics []SmnTopics `json:"smn_topics"`
}

func (o ActionRule) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ActionRule struct{}"
	}

	return strings.Join([]string{"ActionRule", string(data)}, " ")
}
