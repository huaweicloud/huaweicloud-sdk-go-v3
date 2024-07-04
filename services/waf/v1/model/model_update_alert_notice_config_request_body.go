package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateAlertNoticeConfigRequestBody 更新告警通知请求信息
type UpdateAlertNoticeConfigRequestBody struct {

	// 告警通知名称
	Name string `json:"name"`

	// 是否开启   - false: 不开启   - true: 开启
	Enabled *bool `json:"enabled,omitempty"`

	// 主题URN，通过消息通知服务获取
	TopicUrn string `json:"topic_urn"`

	// 时间间隔，单位为分钟。当通知类型为防护事件时，该参数表示在该时间间隔内，攻击次数等于或者大于设定阈值时，将发送告警通知，支持的值：5、15、30、60、120、360、720、1440；当通知类型为证书到期时，该参数表示每隔多长时间发送一次告警通知，支持的值为1440、10080（单位为分钟）。
	Sendfreq *int32 `json:"sendfreq,omitempty"`

	// 语言   - zh-cn：中文   - en-us
	Locale *string `json:"locale,omitempty"`

	// 当通知类型为防护事件时，需要填写该参数。在该时间间隔内，当攻击次数大于或等于您设置的阈值时才会发送告警通知
	Times *int32 `json:"times,omitempty"`

	// 事件类型
	Threat *[]string `json:"threat,omitempty"`

	// 通知类型    - threat_alert_notice: 防护事件    - cert_alert_notice: 证书到期
	NoticeClass string `json:"notice_class"`

	// 提前通知天数，通知类型为证书到期通知需要填写该参数
	NearlyExpiredTime *string `json:"nearly_expired_time,omitempty"`

	// 是否是所有企业项目
	IsAllEnterpriseProject *bool `json:"is_all_enterprise_project,omitempty"`
}

func (o UpdateAlertNoticeConfigRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateAlertNoticeConfigRequestBody struct{}"
	}

	return strings.Join([]string{"UpdateAlertNoticeConfigRequestBody", string(data)}, " ")
}
