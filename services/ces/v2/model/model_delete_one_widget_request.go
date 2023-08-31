package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteOneWidgetRequest Request Object
type DeleteOneWidgetRequest struct {

	// 监控视图id
	WidgetId string `json:"widget_id"`
}

func (o DeleteOneWidgetRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteOneWidgetRequest struct{}"
	}

	return strings.Join([]string{"DeleteOneWidgetRequest", string(data)}, " ")
}
