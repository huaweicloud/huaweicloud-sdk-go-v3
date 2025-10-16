package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteAlertNoticeConfigRequest Request Object
type DeleteAlertNoticeConfigRequest struct {

	// **参数解释：** zh-cn/en-us **约束限制：** 不涉及 **取值范围：** - zh-cn - en-us  **默认取值：** 不涉及
	XLanguage string `json:"X-Language"`

	// **参数解释：** 告警通知id **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	AlertId string `json:"alert_id"`
}

func (o DeleteAlertNoticeConfigRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteAlertNoticeConfigRequest struct{}"
	}

	return strings.Join([]string{"DeleteAlertNoticeConfigRequest", string(data)}, " ")
}
