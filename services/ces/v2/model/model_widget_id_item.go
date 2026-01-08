package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type WidgetIdItem struct {

	// **参数解释** 视图id **取值范围** 字符串必须以wg开头，后跟22个字母和数字，总长度为24个字符
	WidgetId *string `json:"widget_id,omitempty"`
}

func (o WidgetIdItem) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "WidgetIdItem struct{}"
	}

	return strings.Join([]string{"WidgetIdItem", string(data)}, " ")
}
