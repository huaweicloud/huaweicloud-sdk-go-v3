package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowNotificationTemplateResponse Response Object
type ShowNotificationTemplateResponse struct {

	// **参数解释：**  消息模板名称。 **取值范围：**  不涉及。
	Name string `json:"name"`

	// **参数解释：**  消息通知方式。 **取值范围：**  - sms - dingding - wechat - webhook - email - voice - feishu - welink
	Type *[]string `json:"type,omitempty"`

	// **参数解释：**  消息模板描述。 **取值范围：**  不涉及。
	Desc *string `json:"desc,omitempty"`

	// **参数解释：**  消息模板来源。 **取值范围：**  不涉及。
	Source string `json:"source"`

	// **参数解释：**  不同通知渠道下消息模板的详细信息。
	Templates []SubTemplateResBody `json:"templates"`

	// **参数解释：**  消息头语言，系统在发送消息时会默认添加消息头，中文如：“尊敬的用户...”；英文如：“Dear User...”。 **取值范围：**  - zh-cn - en-us
	Locale *string `json:"locale,omitempty"`

	// 创建时间，为毫秒时间戳
	CreateTime int64 `json:"create_time"`

	// 更新时间，为毫秒时间戳
	ModifyTime int64 `json:"modify_time"`

	// 项目ID，获取方式请参见：获取账号ID、项目ID、日志组ID、日志流ID（https://support.huaweicloud.com/api-lts/lts_api_0006.html）。
	ProjectId      string `json:"project_id"`
	HttpStatusCode int    `json:"-"`
}

func (o ShowNotificationTemplateResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowNotificationTemplateResponse struct{}"
	}

	return strings.Join([]string{"ShowNotificationTemplateResponse", string(data)}, " ")
}
