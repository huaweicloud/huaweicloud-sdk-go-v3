package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type AddEmailTemplateRequestBody struct {

	// 数据库类型
	DatastoreType string `json:"datastore_type"`

	// 邮件模板名称
	TemplateName string `json:"template_name"`

	// 实例组ID列表
	GroupId []int32 `json:"group_id"`

	// 健康等级列表
	HealthRank []string `json:"health_rank"`

	// 邮箱地址
	Email *string `json:"email,omitempty"`

	// 主题名称
	Topic *string `json:"topic,omitempty"`

	// 主题标识符
	TopicUrn *string `json:"topic_urn,omitempty"`

	// OBS桶名
	ObsBucketName *string `json:"obs_bucket_name,omitempty"`

	// 诊断时间，12:00-12:00（默认）或00:00-00:00
	InspectionTime string `json:"inspection_time"`

	// 发送时间
	SendTime string `json:"send_time"`

	// 时区
	TimeZone string `json:"time_zone"`
}

func (o AddEmailTemplateRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AddEmailTemplateRequestBody struct{}"
	}

	return strings.Join([]string{"AddEmailTemplateRequestBody", string(data)}, " ")
}
