package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListMonitorIndicatorDataRequest Request Object
type ListMonitorIndicatorDataRequest struct {

	// 开始时间。
	From string `json:"from"`

	// 结束时间。
	To string `json:"to"`

	// 取值方法。
	Function *string `json:"function,omitempty"`

	// 取值周期。
	Period *string `json:"period,omitempty"`

	// 指标名称。
	IndicatorName string `json:"indicator_name"`

	// 第一层级。
	Dim0 string `json:"dim0"`

	// 第二层级。
	Dim1 *string `json:"dim1,omitempty"`
}

func (o ListMonitorIndicatorDataRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListMonitorIndicatorDataRequest struct{}"
	}

	return strings.Join([]string{"ListMonitorIndicatorDataRequest", string(data)}, " ")
}
