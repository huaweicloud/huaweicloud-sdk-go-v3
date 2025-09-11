package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListMetricDatasResponse Response Object
type ListMetricDatasResponse struct {

	// **参数解释**: 实例信息。
	Instances *[]InstancesMetricResult `json:"instances,omitempty"`

	// **参数解释**: 总记录数。 **取值范围**: 不涉及。
	TotalCount     *int32 `json:"total_count,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListMetricDatasResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListMetricDatasResponse struct{}"
	}

	return strings.Join([]string{"ListMetricDatasResponse", string(data)}, " ")
}
