package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateAlertNoticeConfigResponse Response Object
type UpdateAlertNoticeConfigResponse struct {

	// ID
	Id *string `json:"id,omitempty"`

	// 告警通知名称
	Name *string `json:"name,omitempty"`

	// 是否开启   - false: 不开启   - true: 开启
	Enabled *bool `json:"enabled,omitempty"`

	// 主题
	TopicUrn *string `json:"topic_urn,omitempty"`

	// 时间间隔，单位为分钟。当通知类型为防护事件时，该参数表示在该时间间隔内，攻击次数等于或者大于设定阈值时，将发送告警通知，支持的值：5、15、30、60、120、360、720、1440；当通知类型为证书到期时，该参数表示每隔多长时间发送一次告警通知，支持的值为1440、10080（单位为分钟）。
	Sendfreq *int32 `json:"sendfreq,omitempty"`

	// 语言
	Locale *string `json:"locale,omitempty"`

	// 当通知类型为防护事件时，需要填写该参数。在该时间间隔内，当攻击次数大于或等于您设置的阈值时才会发送告警通知
	Times *int32 `json:"times,omitempty"`

	// 事件类型
	Threat *[]string `json:"threat,omitempty"`

	// 预留参数，可忽略
	PreferHtml *bool `json:"prefer_html,omitempty"`

	// 通知类型
	NoticeClass *string `json:"notice_class,omitempty"`

	// 提前通知天数
	NearlyExpiredTime *string `json:"nearly_expired_time,omitempty"`

	// 是否是所有企业项目
	IsAllEnterpriseProject *bool `json:"is_all_enterprise_project,omitempty"`

	// 企业项目ID
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	// 更新时间
	UpdateTime     *int64 `json:"update_time,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o UpdateAlertNoticeConfigResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateAlertNoticeConfigResponse struct{}"
	}

	return strings.Join([]string{"UpdateAlertNoticeConfigResponse", string(data)}, " ")
}
