package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListMonitorIndicatorsResponse Response Object
type ListMonitorIndicatorsResponse struct {

	// **参数解释**： 性能监控指标信息。 **取值范围**： 不涉及。
	Body           *[]IndicatorInfo `json:"body,omitempty"`
	HttpStatusCode int              `json:"-"`
}

func (o ListMonitorIndicatorsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListMonitorIndicatorsResponse struct{}"
	}

	return strings.Join([]string{"ListMonitorIndicatorsResponse", string(data)}, " ")
}
