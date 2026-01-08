package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchUpdateWidgetsRequest Request Object
type BatchUpdateWidgetsRequest struct {

	// **参数解释** 待修改的监控视图列表 **约束限制** 包含的监控视图对象个数为[1,50]
	Body *[]UpdateWidgetInfo `json:"body,omitempty"`
}

func (o BatchUpdateWidgetsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchUpdateWidgetsRequest struct{}"
	}

	return strings.Join([]string{"BatchUpdateWidgetsRequest", string(data)}, " ")
}
