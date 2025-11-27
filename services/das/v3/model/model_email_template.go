package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type EmailTemplate struct {

	// 邮件模板ID
	TemplateId int32 `json:"template_id"`

	// 邮件模板名称
	TemplateName string `json:"template_name"`

	// 数据库类型
	DatastoreType string `json:"datastore_type"`

	// 实例组ID列表
	GroupId []int32 `json:"group_id"`

	// 健康等级列表
	HealthRank []string `json:"health_rank"`

	// 邮箱地址
	Email string `json:"email"`

	// 主题名称
	Topic string `json:"topic"`

	// 主题标识符
	TopicUrn string `json:"topic_urn"`

	// OBS桶名
	ObsBucketName string `json:"obs_bucket_name"`

	// 诊断时间，12:00-12:00（默认）或00:00-00:00
	InspectionTime string `json:"inspection_time"`

	// 发送时间
	SendTime string `json:"send_time"`

	// 时区
	TimeZone string `json:"time_zone"`

	// 更新时间
	UpdateAt int64 `json:"update_at"`

	// 最后修改人
	UserId string `json:"user_id"`

	// 邮件模板状态（0-未启用 1-启用 2-删除）
	Status int32 `json:"status"`
}

func (o EmailTemplate) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EmailTemplate struct{}"
	}

	return strings.Join([]string{"EmailTemplate", string(data)}, " ")
}
