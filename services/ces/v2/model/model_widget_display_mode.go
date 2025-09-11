package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type WidgetDisplayMode struct {

	// **参数解释** 每行展示视图数量 **取值范围** - 0:表示自定义坐标 - 1:代表每行1个视图 - 2:代表每行2个视图 - 3:代表每行3个视图
	RowWidgetNum *int32 `json:"row_widget_num,omitempty"`
}

func (o WidgetDisplayMode) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "WidgetDisplayMode struct{}"
	}

	return strings.Join([]string{"WidgetDisplayMode", string(data)}, " ")
}
