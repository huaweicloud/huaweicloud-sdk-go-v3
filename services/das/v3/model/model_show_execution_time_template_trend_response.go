package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowExecutionTimeTemplateTrendResponse Response Object
type ShowExecutionTimeTemplateTrendResponse struct {

	// 趋势图的时间间隔
	IntervalMillis *int64 `json:"interval_millis,omitempty"`

	// SQL趋势列表
	ItemList       *[]ExTimeTrendItem `json:"item_list,omitempty"`
	HttpStatusCode int                `json:"-"`
}

func (o ShowExecutionTimeTemplateTrendResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowExecutionTimeTemplateTrendResponse struct{}"
	}

	return strings.Join([]string{"ShowExecutionTimeTemplateTrendResponse", string(data)}, " ")
}
