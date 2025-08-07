package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ConfirmThreatMapResponse Response Object
type ConfirmThreatMapResponse struct {

	// **参数解释：** 告警通知里可选的事件类型 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	Threats *[]string `json:"threats,omitempty"`

	Locale         *ThreatMapResponseBodyLocale `json:"locale,omitempty"`
	HttpStatusCode int                          `json:"-"`
}

func (o ConfirmThreatMapResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ConfirmThreatMapResponse struct{}"
	}

	return strings.Join([]string{"ConfirmThreatMapResponse", string(data)}, " ")
}
