package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// NotificationTemplate 消息通知模板列表。
type NotificationTemplate struct {

	// 消息通知模板创建时间。
	CreateTime int64 `json:"create_time"`

	// 消息通知模板描述。
	Desc string `json:"desc"`

	// 消息通知模板所属企业项目id。
	EnterpriseProjectId string `json:"enterprise_project_id"`

	// 消息通知模板语言。
	Locale string `json:"locale"`

	// 消息通知模板修改时间。
	ModifyTime int64 `json:"modify_time"`

	// 消息通知模板名称。
	Name string `json:"name"`

	// 消息通知模板所属项目id。
	ProjectId string `json:"project_id"`

	// 消息通知模板来源。
	Source *string `json:"source,omitempty"`

	// 消息通知模板内容。 消息通知模板内容为json字符串，具体内容是由下列参数拼接成json数组后转义而来。  | 名称 |是否必选  |参数类型  |说明  | |--|--|--|--| |content  |是  |string  |消息模板内容。|  |subType  |是  |string  |消息模板发送类型，支持：email，sms，webhook。| | topic | 否 | string |邮件主题。| | sendType |否  |string  |当消息模板发送类型为“webhook”时需要指定消息模板格式，支持：HTML、JSON。  | | version |是  |string  |默认为v2。 |
	Templates string `json:"templates"`

	// 消息通知方式。
	Type []string `json:"type"`
}

func (o NotificationTemplate) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NotificationTemplate struct{}"
	}

	return strings.Join([]string{"NotificationTemplate", string(data)}, " ")
}
