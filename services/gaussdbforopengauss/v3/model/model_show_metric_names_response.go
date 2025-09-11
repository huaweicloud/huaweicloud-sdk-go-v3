package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowMetricNamesResponse Response Object
type ShowMetricNamesResponse struct {

	// **参数解释**: 指标组名。 **取值范围**: 不涉及。
	GroupName *string `json:"group_name,omitempty"`

	// **参数解释**: 指标分组名称信息。
	MetricNames    *[]MetricNameResult `json:"metric_names,omitempty"`
	HttpStatusCode int                 `json:"-"`
}

func (o ShowMetricNamesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowMetricNamesResponse struct{}"
	}

	return strings.Join([]string{"ShowMetricNamesResponse", string(data)}, " ")
}
