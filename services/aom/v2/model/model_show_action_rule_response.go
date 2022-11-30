package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ShowActionRuleResponse struct {

	// 规则名称
	RuleName *string `json:"rule_name,omitempty"`

	// 项目ID
	ProjectId *string `json:"project_id,omitempty"`

	// 子账号名称
	UserName *string `json:"user_name,omitempty"`

	// 规则描述
	Desc *string `json:"desc,omitempty"`

	// 规则类型。1：通知，2：用户
	Type *string `json:"type,omitempty"`

	// 消息模板
	NotificationTemplate *string `json:"notification_template,omitempty"`

	// 创建时间
	CreateTime *int64 `json:"create_time,omitempty"`

	// 修改时间
	UpdateTime *int64 `json:"update_time,omitempty"`

	// 时区
	TimeZone *string `json:"time_zone,omitempty"`

	// SMN主题信息，不能大于5
	SmnTopics      *[]SmnTopics `json:"smn_topics,omitempty"`
	HttpStatusCode int          `json:"-"`
}

func (o ShowActionRuleResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowActionRuleResponse struct{}"
	}

	return strings.Join([]string{"ShowActionRuleResponse", string(data)}, " ")
}
