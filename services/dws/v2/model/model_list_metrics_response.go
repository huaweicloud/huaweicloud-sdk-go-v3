package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListMetricsResponse Response Object
type ListMetricsResponse struct {

	// **参数解释**： 响应码。 **取值范围**： 不涉及。
	Code *int32 `json:"code,omitempty"`

	// **参数解释**： 响应信息。 **取值范围**： 不涉及。
	Msg *string `json:"msg,omitempty"`

	// **参数解释**： 指标列表。 **取值范围**： 不涉及。
	Data *[]ClusterMetric `json:"data,omitempty"`

	// **参数解释**： 总列表大小。 **取值范围**： 不涉及。
	Count          *int64 `json:"count,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListMetricsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListMetricsResponse struct{}"
	}

	return strings.Join([]string{"ListMetricsResponse", string(data)}, " ")
}
