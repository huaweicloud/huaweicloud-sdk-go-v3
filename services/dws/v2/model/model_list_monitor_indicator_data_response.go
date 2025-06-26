package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListMonitorIndicatorDataResponse Response Object
type ListMonitorIndicatorDataResponse struct {

	// **参数解释**： 历史监控数据响应。 **取值范围**： 不涉及。
	Body           *[]TrendQueryDataResp `json:"body,omitempty"`
	HttpStatusCode int                   `json:"-"`
}

func (o ListMonitorIndicatorDataResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListMonitorIndicatorDataResponse struct{}"
	}

	return strings.Join([]string{"ListMonitorIndicatorDataResponse", string(data)}, " ")
}
