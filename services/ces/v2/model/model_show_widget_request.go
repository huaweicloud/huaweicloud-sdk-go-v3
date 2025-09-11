package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowWidgetRequest Request Object
type ShowWidgetRequest struct {

	// **参数解释**: 监控视图id。 **约束限制**: 不涉及。 **取值范围**: 字符串必须以wg开头，包含22个字母和数字，长度为24个字符。 **默认取值**: 不涉及。
	WidgetId string `json:"widget_id"`
}

func (o ShowWidgetRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowWidgetRequest struct{}"
	}

	return strings.Join([]string{"ShowWidgetRequest", string(data)}, " ")
}
