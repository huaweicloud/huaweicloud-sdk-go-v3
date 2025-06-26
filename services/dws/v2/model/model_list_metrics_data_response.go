package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListMetricsDataResponse Response Object
type ListMetricsDataResponse struct {

	// **参数解释**： 响应码。 **取值范围**： 不涉及。
	Code *int32 `json:"code,omitempty"`

	// **参数解释**： 响应信息。 **取值范围**： 不涉及。
	Msg *string `json:"msg,omitempty"`

	// **参数解释**： 指标采集数据列表。 **取值范围**： 不涉及。
	Data *[]map[string]interface{} `json:"data,omitempty"`

	// **参数解释**： 总列表大小。 **取值范围**： 不涉及。
	Count          *int64 `json:"count,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListMetricsDataResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListMetricsDataResponse struct{}"
	}

	return strings.Join([]string{"ListMetricsDataResponse", string(data)}, " ")
}
