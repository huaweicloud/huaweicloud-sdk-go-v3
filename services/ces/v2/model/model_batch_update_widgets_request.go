package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchUpdateWidgetsRequest Request Object
type BatchUpdateWidgetsRequest struct {

	// 待修改的监控视图列表
	Body *[]UpdateWidgetInfo `json:"body,omitempty"`
}

func (o BatchUpdateWidgetsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchUpdateWidgetsRequest struct{}"
	}

	return strings.Join([]string{"BatchUpdateWidgetsRequest", string(data)}, " ")
}
