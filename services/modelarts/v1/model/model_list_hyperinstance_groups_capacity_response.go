package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListHyperinstanceGroupsCapacityResponse Response Object
type ListHyperinstanceGroupsCapacityResponse struct {

	// **参数解释**：hyperinstance-group容量预测结果列表。 **默认取值**：不涉及。
	Capacities     *[]ServerHpsGroupCapacity `json:"capacities,omitempty"`
	HttpStatusCode int                       `json:"-"`
}

func (o ListHyperinstanceGroupsCapacityResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListHyperinstanceGroupsCapacityResponse struct{}"
	}

	return strings.Join([]string{"ListHyperinstanceGroupsCapacityResponse", string(data)}, " ")
}
