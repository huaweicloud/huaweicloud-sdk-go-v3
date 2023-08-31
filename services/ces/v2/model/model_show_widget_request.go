package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowWidgetRequest Request Object
type ShowWidgetRequest struct {

	// 监控视图id
	WidgetId string `json:"widget_id"`
}

func (o ShowWidgetRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowWidgetRequest struct{}"
	}

	return strings.Join([]string{"ShowWidgetRequest", string(data)}, " ")
}
