package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateNotificationTemplateResponse Response Object
type UpdateNotificationTemplateResponse struct {

	// **参数解释：**  消息模板名称。 **取值范围：**  不涉及。
	Name *string `json:"name,omitempty"`

	// **参数解释：**  消息通知方式。 **取值范围：**  - sms - dingding - wechat - webhook - email - voice - feishu - welink
	Type *[]string `json:"type,omitempty"`

	// **参数解释：**  消息模板描述。 **取值范围：**  不涉及。
	Desc *string `json:"desc,omitempty"`

	// **参数解释：**  消息模板来源。 **取值范围：**  不涉及。
	Source *string `json:"source,omitempty"`

	// **参数解释：**  不同通知渠道下消息模板的详细信息。
	Templates *[]SubTemplateResBody `json:"templates,omitempty"`

	// **参数解释：**  消息头语言，系统在发送消息时会默认添加消息头，中文如：“尊敬的用户...”；英文如：“Dear User...”。 **取值范围：**  - zh-cn - en-us
	Locale         *string `json:"locale,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o UpdateNotificationTemplateResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateNotificationTemplateResponse struct{}"
	}

	return strings.Join([]string{"UpdateNotificationTemplateResponse", string(data)}, " ")
}
