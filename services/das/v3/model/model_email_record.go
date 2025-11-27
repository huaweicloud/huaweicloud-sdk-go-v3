package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type EmailRecord struct {

	// 发送时间（Unix timestamp），单位：毫秒。
	SendAt int64 `json:"send_at"`

	// 发送状态
	SendStatus int32 `json:"send_status"`

	// 邮箱地址
	Email string `json:"email"`

	// 主题名称
	Topic string `json:"topic"`

	// 主题标识符
	TopicUrn string `json:"topic_urn"`

	// 实例诊断报告列表
	InstanceHealthReportList []InstanceHealthReport `json:"instance_health_report_list"`
}

func (o EmailRecord) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EmailRecord struct{}"
	}

	return strings.Join([]string{"EmailRecord", string(data)}, " ")
}
