package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowAlertRequest Request Object
type ShowAlertRequest struct {

	// **参数解释：** 内容类型 - application/json;charset=UTF-8    普通API请求的类型  **约束限制：** 不涉及 **取值范围：** - application/json;charset=UTF-8  **默认取值：** 不涉及
	ContentType string `json:"content-type"`

	// **参数解释：** 工作空间id。 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	WorkspaceId string `json:"workspace_id"`

	// **参数解释：** 告警id **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	AlertId string `json:"alert_id"`
}

func (o ShowAlertRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowAlertRequest struct{}"
	}

	return strings.Join([]string{"ShowAlertRequest", string(data)}, " ")
}
